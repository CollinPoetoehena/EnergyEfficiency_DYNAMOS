package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"crypto/sha256"
    "encoding/hex"

	"github.com/Jorrit05/DYNAMOS/pkg/api"
	"github.com/Jorrit05/DYNAMOS/pkg/etcd"
	"github.com/Jorrit05/DYNAMOS/pkg/lib"
	pb "github.com/Jorrit05/DYNAMOS/pkg/proto"
	"github.com/google/uuid"
	"go.opencensus.io/trace"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"github.com/redis/go-redis/v9"
)

// Getting the SQL request through HTTP. This means the request is coming from the user. So it can be either a computeToData or DataThroughTtp request.
// Based on the role we have, it will be handled as one or the other.
func sqlDataRequestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("Entering sqlDataRequestHandler")
		// Start a new span with the context that has a timeout

		ctxWithTimeout, cancel := context.WithTimeout(r.Context(), 30*time.Second)
		defer cancel()

		ctx, span := trace.StartSpan(ctxWithTimeout, serviceName+"/func: sqlDataRequestHandler")
		defer span.End()

		body, err := api.GetRequestBody(w, r, serviceName)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		sqlDataRequest := &pb.SqlDataRequest{}
		sqlDataRequest.RequestMetadata = &pb.RequestMetadata{}

		err = protojson.Unmarshal(body, sqlDataRequest)
		if err != nil {
			logger.Sugar().Warnf("Error unmarshalling sqlDataRequest: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if sqlDataRequest.RequestMetadata.JobId == "" {
			http.Error(w, "Job ID not passed", http.StatusInternalServerError)
			return
		}

		// Get the matching composition request and determine our role
		// /agents/jobs/UVA/jorrit-3141334
		logger.Sugar().Debug("Getting composition request")
		compositionRequest, err := getCompositionRequest(sqlDataRequest.User.UserName, sqlDataRequest.RequestMetadata.JobId)
		if err != nil {
			http.Error(w, "No job found for this user", http.StatusBadRequest)
			return
		}
		logger.Sugar().Debugf("Composition request: %+v", compositionRequest)

		// Generate correlationID for this request
		correlationId := uuid.New().String()

		// Get the role
		role := compositionRequest.Role
		logger.Sugar().Debugf("Role for data request: %+s", role)
		// Check role before caching functionality
		if !strings.EqualFold(role, "computeProvider") && !strings.EqualFold(role, "all") {
			logger.Sugar().Warnf("Unknown role or unexpected HTTP request: %s", compositionRequest.Role)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		// Create cache key based on request (use existing context (ctx) for redis)
        // Do this after all the policy checks (including role check)
		// Use hashing for efficient storage and retrieval
        queryHash := sha256.Sum256([]byte(sqlDataRequest.Query))
        queryHashStr := hex.EncodeToString(queryHash[:])
        optionsJSON, _ := json.Marshal(sqlDataRequest.Options)
        optionsHash := sha256.Sum256(optionsJSON)
        optionsHashStr := hex.EncodeToString(optionsHash[:])
        cacheKey := fmt.Sprintf("composition:%s:%s:%s:%s:%s", 
			// Use role for unique matching between archetypes
			role, sqlDataRequest.User.UserName, queryHashStr, 
			sqlDataRequest.Algorithm, optionsHashStr)
		logger.Sugar().Debugf("Cache key: %+s", cacheKey)
		// Check if the response is already cached
		cachedResponse, err := redisClient.Get(ctx, cacheKey).Result()
		if err == nil {
			// Cache hit: Return cached response
			logger.Sugar().Infof("Cache hit for key: %s", cacheKey)
			// Handle response information
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(cachedResponse))
			// Return to stop the function (cache hit found, no need to process further)
			return
		} else if err != redis.Nil {
			// Redis error
			logger.Sugar().Errorf("Redis error when checking cache in sqlDataRequest: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		// Cache miss: Process the request further
		logger.Sugar().Infof("Cache miss for key: %s", cacheKey)

		// Switch on the role we have in this data request
		logger.Sugar().Debug("Switching on role for this data request")
		if strings.EqualFold(role, "computeProvider") {
			ctx, err = handleSqlComputeProvider(ctx, compositionRequest.LocalJobName, compositionRequest, sqlDataRequest, correlationId)
			if err != nil {
				logger.Sugar().Errorf("Error in computeProvider role: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

		} else if strings.EqualFold(role, "all") {
			ctx, err = handleSqlAll(ctx, compositionRequest.LocalJobName, compositionRequest, sqlDataRequest, correlationId)
			if err != nil {
				logger.Sugar().Errorf("Error in all role: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
		} 

		// Create a channel to receive the response
		responseChan := make(chan dataResponse)

		// Store the request information in the map
		mutex.Lock()
		responseMap[correlationId] = responseChan
		mutex.Unlock()

		select {
		case dataResponseStruct := <-responseChan:
			msComm := dataResponseStruct.response

			logger.Sugar().Debugf("Received response, %s", msComm.RequestMetadata.CorrelationId)
			msgBytes, err := proto.Marshal(msComm)
			if err != nil {
				logger.Sugar().Warnf("error marshalling proto message, %v", err)
			}
			jsonBytes, err := json.Marshal(msComm)
			if err != nil {
				logger.Sugar().Warnf("error marshalling jsonBytes message, %v", err)
			}

			span.AddAttributes(trace.Int64Attribute("sqlDataRequestHandler.proto.messageSize", int64(len(msgBytes))))
			span.AddAttributes(trace.Int64Attribute("sqlDataRequestHandler.json.messageSize", int64(len(jsonBytes))))
			span.AddAttributes(trace.Int64Attribute("sqlDataRequestHandler.String.messageSize", int64(len(msComm.Result))))
			logger.Sugar().Debugf("Got result (size): %d", len(msComm.Result))
			// logger.Sugar().Debugf("Result: %s", msComm.Result)

			// Store the response in the cache with a TTL (Time-To-Live)
			// If code gets here no cache hit is found
			logger.Debug("Storing result in cache...")
			err = redisClient.Set(ctx, cacheKey, msComm.Result, 10*time.Minute).Err()
			if err != nil {
				logger.Sugar().Errorf("Failed to cache response: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			//Handle response information
			w.WriteHeader(http.StatusOK)
			w.Write(msComm.Result)
			return

		case <-ctx.Done():
			http.Error(w, "Request timed out", http.StatusRequestTimeout)
			return
		}
	}
}

// handleSqlAll means we do all work for this request, not third part involved (computeToData archeType)
func handleSqlAll(ctx context.Context, jobName string, compositionRequest *pb.CompositionRequest, sqlDataRequest *pb.SqlDataRequest, correlationId string) (context.Context, error) {
	logger.Sugar().Debug("Handling sql all")
	// Create msChain and deploy job.

	ctx, span := trace.StartSpan(ctx, serviceName+"/func: handleSqlAll")
	defer span.End()

	var err error
	ctx, _, err = generateChainAndDeploy(ctx, compositionRequest, jobName, sqlDataRequest.Options)
	if err != nil {
		logger.Sugar().Errorf("error deploying job: %v", err)
		return ctx, err
	}

	msComm := &pb.MicroserviceCommunication{}
	msComm.RequestMetadata = &pb.RequestMetadata{}
	msComm.Type = "microserviceCommunication"
	msComm.RequestMetadata.DestinationQueue = jobName
	msComm.RequestMetadata.ReturnAddress = agentConfig.RoutingKey
	msComm.RequestType = compositionRequest.RequestType

	any, err := anypb.New(sqlDataRequest)
	if err != nil {
		logger.Sugar().Error(err)
		return ctx, err
	}

	msComm.OriginalRequest = any
	msComm.RequestMetadata.CorrelationId = correlationId

	logger.Sugar().Debugf("Sending SendMicroserviceInput to: %s", jobName)

	key := fmt.Sprintf("/agents/jobs/%s/queueInfo/%s", serviceName, jobName)
	err = etcd.PutEtcdWithGrant(ctx, etcdClient, key, jobName, queueDeleteAfter)
	if err != nil {
		logger.Sugar().Errorf("Error PutEtcdWithGrant: %v", err)
	}

	c.SendMicroserviceComm(ctx, msComm)
	return ctx, nil
}

// handleSqlComputeProvider means we have a computeProvider role only (dataThroughTtp archeType)
// We are responsible for forwarding the request to all dataProviders.
func handleSqlComputeProvider(ctx context.Context, jobName string, compositionRequest *pb.CompositionRequest, sqlDataRequest *pb.SqlDataRequest, correlationId string) (context.Context, error) {
	logger.Sugar().Debug("Handling sql compute provider")

	ctx, span := trace.StartSpan(ctx, serviceName+"/func: handleSqlComputeProvider")
	defer span.End()

	// pack and send request to all data providers, add own routing key as return address
	// check request and spin up own job (generate mschain, deployjob)
	if len(compositionRequest.DataProviders) == 0 {
		return ctx, fmt.Errorf("expected to know dataproviders")
	}

	for _, dataProvider := range compositionRequest.DataProviders {
		dataProviderRoutingKey := fmt.Sprintf("/agents/online/%s", dataProvider)
		var agentData lib.AgentDetails
		_, err := etcd.GetAndUnmarshalJSON(etcdClient, dataProviderRoutingKey, &agentData)
		if err != nil {
			return ctx, fmt.Errorf("error getting dataProvider dns")
		}

		sqlDataRequest.RequestMetadata.DestinationQueue = agentData.RoutingKey

		// This is a bit confusing, but it tells the other agent to go back here.
		// The other agent, will reset the address to get the message from the job.
		sqlDataRequest.RequestMetadata.ReturnAddress = agentConfig.RoutingKey

		sqlDataRequest.RequestMetadata.CorrelationId = correlationId
		sqlDataRequest.RequestMetadata.JobName = compositionRequest.JobName
		logger.Sugar().Debugf("Sending sqlDataRequest to: %s", sqlDataRequest.RequestMetadata.DestinationQueue)

		key := fmt.Sprintf("/agents/jobs/%s/queueInfo/%s", serviceName, jobName)
		err = etcd.PutEtcdWithGrant(ctx, etcdClient, key, jobName, queueDeleteAfter)
		if err != nil {
			logger.Sugar().Errorf("Error PutEtcdWithGrant: %v", err)
		}

		_, err = c.SendSqlDataRequest(ctx, sqlDataRequest)
		if err != nil {
			logger.Sugar().Errorf("Error c.SendSqlDataRequest: %v", err)
		}
	}

	// TODO: Parse SQL request for extra compute services
	var err error
	ctx, createdJob, err := generateChainAndDeploy(ctx, compositionRequest, jobName, sqlDataRequest.Options)
	if err != nil {
		logger.Sugar().Errorf("error deploying job: %v", err)
	}
	logger.Sugar().Debugf("Created job: %s", createdJob.Name)
	waitingJobMutex.Lock()
	waitingJobMap[sqlDataRequest.RequestMetadata.CorrelationId] = &waitingJob{job: createdJob, nrOfDataStewards: len(compositionRequest.DataProviders)}
	waitingJobMutex.Unlock()
	logger.Sugar().Debugf("Created job nr of stewards: %d", waitingJobMap[sqlDataRequest.RequestMetadata.CorrelationId].nrOfDataStewards)

	return ctx, nil
}
