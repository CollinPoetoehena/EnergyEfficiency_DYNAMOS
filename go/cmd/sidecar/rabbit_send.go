// Package main, that implements 'sidecar' functionality
//
// File: rabbit_send.go
//
// Description:
// This file contains the server implementations of sending an AMQ message to a destination queue.
// A DYNAMOS service will package the message, and call this gRPC implementation
// on the sidecar to send the message.
//
// Notes:
//
//
// Author: Jorrit Stutterheim

package main

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/Jorrit05/DYNAMOS/pkg/etcd"
	pb "github.com/Jorrit05/DYNAMOS/pkg/proto"
	bo "github.com/cenkalti/backoff/v4"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"
)

// This function sends a message to the specified AMQ target queue there is a retry mechanism in place
// that will retry sending the message if the target queue is not available.
//
// Parameters:
// - ctx: Context
// - message: amqp.Publishing formated message to send.
// - target: destination queue.
// - s:  serverInstance pointer, used to access the channel.
// - opts: The options to apply to the retry mechanism.
//
// Returns:
// - An empty protobuf message.
// - An error if the message could not be sent, otherwise nil.
//
// Notes:
// If the message is undeliverable, the message will be returned to the sender with a NO_ROUTE reply text.
// The function waits 5 seconds for this return message, otherwise it will continue. Everytime a message
// is sent the global variable running_messages is incremented, and decremented when the message is sent.
// This is to ensure that the server does not shut down while messages are still being sent.
func send(ctx context.Context, message amqp.Publishing, target string, s *serverInstance, opts ...etcd.Option) (*emptypb.Empty, error) {
	// Start with default options
	retryOpts := etcd.DefaultRetryOptions
	running_messages += 1

	// Apply any specified options
	for _, opt := range opts {
		opt(&retryOpts)
	}

	// Create a returns channel
	returns := make(chan amqp.Return)
	s.channel.NotifyReturn(returns)

	if message.Headers == nil {
		message.Headers = amqp.Table{}
	}

	sc := trace.FromContext(ctx).SpanContext()
	binarySc := propagation.Binary(sc)

	if retryOpts.AddJsonTrace {
		// Create a map to hold the span context values
		scMap := map[string]string{
			"TraceID": sc.TraceID.String(),
			"SpanID":  sc.SpanID.String(),
			// "TraceOptions": fmt.Sprintf("%02x", sc.TraceOptions.IsSampled()),
		}
		// Serialize the map to a JSON string
		scJson, err := json.Marshal(scMap)
		if err != nil {
			logger.Debug("ERRROR scJson MAP")
		}
		message.Headers["jsonTrace"] = scJson
	}

	message.Headers["binaryTrace"] = binarySc

	operation := func() error {
		// Log before sending message
		logger.Sugar().Infow("Sending message: ", "My routingKey", s.routingKey, "exchangeName", exchangeName, "target", target)

		// Create a context with a timeout
		timeoutCtx, cancel := context.WithTimeout(ctx, 8*time.Second)
		defer cancel()

		err := s.channel.PublishWithContext(timeoutCtx, exchangeName, target, true, false, message)
		if err != nil {
			logger.Sugar().Debugf("In error chan: %v", err)
			return err
		}

		select {
		case r := <-returns:
			if r.ReplyText == "NO_ROUTE" {
				logger.Sugar().Infof("No route to job yet: %v", target)
				// This will trigger a retry
				return errors.New("no route to target")
			} else {
				logger.Sugar().Errorf("Unknown reason message returned: %v", r)
				return bo.Permanent(errors.New("unknown error"))
			}
		case <-time.After(5 * time.Second): // Timeout if no message is received in 5 seconds
			logger.Sugar().Debugf("5 seconds have passed for target: %v", target)
		}

		return nil
	}

	// Create a new exponential backoff
	backoff := bo.NewExponentialBackOff()
	backoff.InitialInterval = retryOpts.InitialInterval
	backoff.MaxInterval = retryOpts.MaxInterval
	backoff.MaxElapsedTime = retryOpts.MaxElapsedTime

	err := bo.Retry(operation, backoff)
	logger.Sugar().Debugf("Backoff!!")
	if err != nil {
		logger.Sugar().Errorf("Publish failed after %v seconds, err: %s", backoff.MaxElapsedTime, err)
		running_messages -= 1
		return &emptypb.Empty{}, err
	}
	running_messages -= 1

	return &emptypb.Empty{}, nil
}

func (s *serverInstance) SendRequestApproval(ctx context.Context, in *pb.RequestApproval) (*emptypb.Empty, error) {
	data, err := proto.Marshal(in)
	if err != nil {
		logger.Sugar().Errorf("Marshal requestApproval failed: %s", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	// Do other stuff
	message := amqp.Publishing{
		Body: data,
		Type: "requestApproval",
	}

	go send(ctx, message, in.DestinationQueue, s)
	return &emptypb.Empty{}, nil
}

func (s *serverInstance) SendValidationResponse(ctx context.Context, in *pb.ValidationResponse) (*emptypb.Empty, error) {
	data, err := proto.Marshal(in)
	if err != nil {
		logger.Sugar().Errorf("Marshal ValidationResponse failed: %s", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	// Do other stuff
	message := amqp.Publishing{
		Body: data,
		Type: "validationResponse",
	}

	go send(ctx, message, "orchestrator-in", s)
	return &emptypb.Empty{}, nil

}

func (s *serverInstance) SendCompositionRequest(ctx context.Context, in *pb.CompositionRequest) (*emptypb.Empty, error) {
	data, err := proto.Marshal(in)
	if err != nil {
		logger.Sugar().Errorf("Marshal CompositionRequest failed: %s", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	// Do other stuff
	message := amqp.Publishing{
		Body: data,
		Type: "compositionRequest",
	}

	go send(ctx, message, in.DestinationQueue, s)
	return &emptypb.Empty{}, nil

}

func (s *serverInstance) SendSqlDataRequest(ctx context.Context, in *pb.SqlDataRequest) (*emptypb.Empty, error) {
	data, err := proto.Marshal(in)
	if err != nil {
		logger.Sugar().Errorf("Marshal requestApproval failed: %s", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	// Do other stuff
	message := amqp.Publishing{
		CorrelationId: in.RequestMetadata.CorrelationId,
		Body:          data,
		Type:          "sqlDataRequest",
	}

	logger.Sugar().Debugf("SendSqlDataRequest destination queue: %v", in.RequestMetadata.DestinationQueue)
	go send(ctx, message, in.RequestMetadata.DestinationQueue, s, etcd.WithMaxElapsedTime(10*time.Second))
	return &emptypb.Empty{}, nil

}

func (s *serverInstance) SendPolicyUpdate(ctx context.Context, in *pb.PolicyUpdate) (*emptypb.Empty, error) {
	data, err := proto.Marshal(in)
	if err != nil {
		logger.Sugar().Errorf("Marshal PolicyUpdate failed: %s", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	// Do other stuff
	message := amqp.Publishing{
		CorrelationId: in.RequestMetadata.CorrelationId,
		Body:          data,
		Type:          "policyUpdate",
	}

	logger.Sugar().Debugf("PolicyUpdate destination queue: %s", in.RequestMetadata.DestinationQueue)
	go send(ctx, message, in.RequestMetadata.DestinationQueue, s, etcd.WithMaxElapsedTime(10*time.Second))
	return &emptypb.Empty{}, nil

}

func (s *serverInstance) SendMicroserviceComm(ctx context.Context, in *pb.MicroserviceCommunication) (*emptypb.Empty, error) {
	// Logging of data send for compression testing 
	// TODO compression: remove logging later when done
	logger.Sugar().Debugf("**********************Microservice communication type (in go/cmd/sidecar/rabbit_send.go): %s", in.Type)
	logger.Sugar().Debugf("**********************Microservice communication request type (in go/cmd/sidecar/rabbit_send.go): %s", in.RequestType)
	// Convert the google.protobuf.Struct to a JSON string
    dataJSON, err := protojson.Marshal(in.Data)
    if err != nil {
        logger.Sugar().Errorf("Failed to marshal data to JSON: %s", err)
        return nil, status.Error(codes.Internal, err.Error())
    }
    logger.Sugar().Debugf("**********************Microservice communication data (in go/cmd/sidecar/rabbit_send.go): %s", dataJSON)
    logger.Sugar().Debugf("**********************Microservice communication result (in go/cmd/sidecar/rabbit_send.go): %s", in.Result)
    // TODO compression: compress something in data here, such as .Result?

	data, err := proto.Marshal(in)
	if err != nil {
		logger.Sugar().Errorf("Marshal SendMicroserviceComm failed: %s", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	// TODO compression: remove this logging later as well
	// Convert the serialized data back to a protobuf message
    var unmarshaledMessage pb.MicroserviceCommunication
    err = proto.Unmarshal(data, &unmarshaledMessage)
    if err != nil {
        logger.Sugar().Errorf("Failed to unmarshal data: %s", err)
        return nil, status.Error(codes.Internal, err.Error())
    }
    // Convert the protobuf message to a JSON string
    messageJSON, err := protojson.Marshal(&unmarshaledMessage)
    if err != nil {
        logger.Sugar().Errorf("Failed to marshal message to JSON: %s", err)
        return nil, status.Error(codes.Internal, err.Error())
    }
    // Print the JSON string
    logger.Sugar().Debugf("**********************Serialized data (JSON) (in go/cmd/sidecar/rabbit_send.go): %s", messageJSON)
	// TODO compression: result field now compressed in data?

	// TODO compression: when testing this with compute to Data archetype this was the result:
	/*
	2025-01-17T10:45:05.829810935Z sidecar 1.737110705829561e+09    DEBUG    /app/cmd/sidecar/rabbit_send.go:252    
	**********************Microservice communication data (in go/cmd/sidecar/rabbit_send.go): {}  │
│ 2025-01-17T10:45:05.829817797Z sidecar 1.7371107058295848e+09    DEBUG    /app/cmd/sidecar/rabbit_send.go:253    
**********************Microservice communication result (in go/cmd/sidecar/rabbit_send.go):  │
│ 2025-01-17T10:45:05.829884542Z sidecar 1.737110705829745e+09    DEBUG    /app/cmd/sidecar/rabbit_send.go:277    
**********************Serialized data (JSON) (in go/cmd/sidecar/rabbit_send.go): {"type":"mic │
│ roserviceCommunication", "requestType":"sqlDataRequest", "originalRequest":{"@type":"type.googleapis.com/dynamos.SqlDataRequest", "type":"sqlDataRequest", "query":"SELECT DISTINCT p.Unieknr, p.Geslacht, p. │
│ Gebdat, s.Aanst_22, s.Functcat, s.Salschal as Salary FROM Personen p JOIN Aanstellingen s ON p.Unieknr = s.Unieknr LIMIT 5", "user":{"id":"12324", "userName":"jorrit.stutterheim@cloudnation.nl"}, "requestM │
│ etadata":{"jobId":"jorrit-stutterheim-fe4a5185"}, "options":{"aggregate":false, "graph":false}}, "requestMetadata":{"correlationId":"e98d7275-d169-4539-82f6-61c2ec0c76c5", "destinationQueue":"jorrit-stutte │
│ rheim-fe4a5185uva1", "returnAddress":"UVA-in"}} 
	*/
	// As you can see, there is no result and no data field set, and the body is the serialized JSON part, which is just the orignial request as 
	// MicroserviceCommunication .proto message, but the fields 'data' and 'result' are not present. 
	// TODO: test this for data through TTP archetype as well! If the same is found there, then no compression is necessary here and no changes in this part
	// are required, only in the other parts!
	

	message := amqp.Publishing{
		CorrelationId: in.RequestMetadata.CorrelationId,
		Body:          data,
		Type:          in.Type,
	}
	logger.Sugar().Debugf("SendMicroserviceComm destination queue: %s", in.RequestMetadata.DestinationQueue)
	go send(ctx, message, in.RequestMetadata.DestinationQueue, s, etcd.WithMaxElapsedTime(10*time.Second), etcd.WithJsonTrace())

	return &emptypb.Empty{}, nil
}

func (s *serverInstance) SendTest(ctx context.Context, in *pb.SqlDataRequest) (*emptypb.Empty, error) {
	data, err := proto.Marshal(in)
	if err != nil {
		logger.Sugar().Errorf("Marshal SendMicroserviceComm failed: %s", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	// Do other stuff
	message := amqp.Publishing{
		CorrelationId: "in.RequestMetadata.CorrelationId",
		Body:          data,
		Type:          "testSet",
	}
	go send(ctx, message, "no existss", s, etcd.WithMaxElapsedTime(10*time.Second))
	return &emptypb.Empty{}, nil
}

func (s *serverInstance) SendRequestApprovalResponse(ctx context.Context, in *pb.RequestApprovalResponse) (*emptypb.Empty, error) {
	data, err := proto.Marshal(in)
	if err != nil {
		logger.Sugar().Errorf("Marshal SendRequestApprovalResponse failed: %s", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	// Do other stuff
	message := amqp.Publishing{
		Body: data,
		Type: "requestApprovalResponse",
	}

	go send(ctx, message, in.RequestMetadata.DestinationQueue, s)
	return &emptypb.Empty{}, nil
}
