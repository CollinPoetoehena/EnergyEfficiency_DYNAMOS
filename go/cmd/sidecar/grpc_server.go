// Package main, that implements 'sidecar' functionality
//
// File: grpc_server.go
//
// Description:
// This file contains the gRPC server implementation for the sidecar.
// It contains the serverInstance struct and the method implemntations for gRPC calls that the sidecar uses.
//
// Notes:
// There are some generic gRPC methods implemented under the 'lib.sharedServer' struct. These methods
// when registering a server a choice can be made to register the 'serverInstance' server
// or the 'sharedServer' server or both. The sharedServer implements the Health and Generic gRPC services.
//
// Author: Jorrit Stutterheim

package main

import (
	"context"
	"fmt"

	"github.com/Jorrit05/DYNAMOS/pkg/lib"
	pb "github.com/Jorrit05/DYNAMOS/pkg/proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ConsumerManager manages the consumers for the gRPC server.
type ConsumerManager struct {
	stopChan chan struct{}
	cancel   context.CancelFunc
}

type serverInstance struct {
	pb.UnimplementedRabbitMQServer
	pb.UnimplementedEtcdServer
	pb.UnimplementedMicroserviceServer
	consumerManager *ConsumerManager
	channel         *amqp.Channel
	conn            *amqp.Connection
	routingKey      string
}

func (s *serverInstance) InitRabbitMq(ctx context.Context, in *pb.InitRequest) (*emptypb.Empty, error) {
	logger.Sugar().Infow("Received:", "Servicename", in.ServiceName, "RoutingKey", in.RoutingKey)

	// Call the SetupConnection function and handle the message consumption inside this function
	_, conn, channel, err := setupConnection(in.ServiceName, in.RoutingKey, in.QueueAutoDelete)
	s.channel = channel
	s.conn = conn

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *serverInstance) CreateQueue(ctx context.Context, in *pb.QueueInfo) (*emptypb.Empty, error) {
	queue, err := declareQueue(in.QueueName, s.channel, in.AutoDelete)
	if err != nil {
		logger.Sugar().Fatalw("Failed to declare queue: %v", err)
		return nil, err
	}
	if err := s.channel.QueueBind(
		queue.Name,       // name
		in.QueueName,     // key
		"topic_exchange", // exchange
		false,            // noWait
		nil,              // args
	); err != nil {
		logger.Sugar().Fatalw("Queue Bind: %s", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *serverInstance) DeleteQueue(ctx context.Context, in *pb.QueueInfo) (*emptypb.Empty, error) {
	logger.Sugar().Debugf("Delete Queue: %s", in.QueueName)

	purgedMessages, err := s.channel.QueueDelete(in.QueueName, false, false, false)
	if err != nil {
		logger.Sugar().Warnf("Error deleting queue: %v", err)
		return &emptypb.Empty{}, err
	}

	logger.Sugar().Infof("Deleted queue %s, purged %d, messages", in.QueueName, purgedMessages)
	return &emptypb.Empty{}, nil
}

// Consume consumes messages from a specified queue and handles them based on their type.
// It takes a ConsumeRequest and a stream (RabbitMQ_ConsumeServer) as input parameters.
// Returns an error if there was an issue consuming the messages or handling them.
func (s *serverInstance) Consume(in *pb.ConsumeRequest, stream pb.RabbitMQ_ConsumeServer) error {
	messages, err := s.channel.Consume(
		in.QueueName, // queue
		"",           // consumer
		in.AutoAck,   // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("Started consuming from %s", in.QueueName)

	for msg := range messages {
		logger.Sugar().Debugw("switchin: ", "msg,Type", msg.Type, "Port:", port)
		switch msg.Type {
		case "validationResponse":
			if err := s.handleValidationResponse(msg, stream); err != nil {
				logger.Sugar().Errorf("Error handling validation response: %v", err)
				return status.Error(codes.Internal, err.Error())
			}
		case "requestApproval":
			if err := s.handleRequestApprovalResponse(msg, stream); err != nil {
				logger.Sugar().Errorf("Error handling requestApproval response: %v", err)
				return status.Error(codes.Internal, err.Error())
			}
		case "compositionRequest":
			if err := s.handleCompositionRequestResponse(msg, stream); err != nil {
				logger.Sugar().Errorf("Error handling validation response: %v", err)
				return status.Error(codes.Internal, err.Error())
			}
		case "sqlDataRequest":
			if err := s.handleSqlDataRequest(msg, stream); err != nil {
				logger.Sugar().Errorf("Error handling sqlData request: %v", err)
				return status.Error(codes.Internal, err.Error())
			}
		case "microserviceCommunication":
			if err := s.handleMicroserviceCommunication(msg, stream); err != nil {
				logger.Sugar().Errorf("Error handling microserviceCommunication: %v", err)
				return status.Error(codes.Internal, err.Error())
			}
		case "policyUpdate":
			if err := s.handlePolicyUpdate(msg, stream); err != nil {
				logger.Sugar().Errorf("Error handling policyUpdate: %v", err)
				return status.Error(codes.Internal, err.Error())
			}
		case "requestApprovalResponse":
			if err := s.handleRequestApprovalToApiResponse(msg, stream); err != nil {
				logger.Sugar().Errorf("Error handling requestApprovalResponse: %v", err)
				return status.Error(codes.Internal, err.Error())
			}
		// Handle other message types...
		default:
			logger.Sugar().Errorf("Unknown message type: %s", msg.Type)
			return status.Error(codes.Unknown, fmt.Sprintf("Unknown message type: %s", msg.Type))
		}
	}
	s.channel.Close()
	return nil
}

// ServerInstance implementation for SendData, this is called by the sidecar of an agent to forward the
// message to RabbitMQ.
//
// The different SendData functions are picked by either registring the 'serverInstance' or the 'sharedServer' instances in your
// gRPC server.
//
// Parameters:
//   - ctx: The context of the request
//   - data: MicroserviceCommunication messages
//
// Returns:
//   - ContinueReceiving: A boolean indicating if the sidecar should continue receiving messages
//   - error: An error if the function fails
func (s *serverInstance) SendData(ctx context.Context, data *pb.MicroserviceCommunication) (*pb.ContinueReceiving, error) {
	logger.Sugar().Debugf("Starting (to AMQ) lib.SendData: %v", data.RequestMetadata.DestinationQueue)

	// Logging of data send for compression testing
	// TODO compression: remove logging later when done
	logger.Sugar().Debugf("**********************Microservice communication type (in go/cmd/sidecar/grpc_server.go): %s", data.Type)
	logger.Sugar().Debugf("**********************Microservice communication request type (in go/cmd/sidecar/grpc_server.go): %s", data.RequestType)
	// Convert the google.protobuf.Struct to a JSON string
    dataJSON, err := protojson.Marshal(data.Data)
    if err != nil {
        logger.Sugar().Errorf("Failed to marshal data to JSON: %s", err)
    }
	logger.Sugar().Debugf("**********************Microservice communication data (in go/cmd/sidecar/grpc_server.go): %s", dataJSON)
	logger.Sugar().Debugf("**********************Microservice communication result (in go/cmd/sidecar/grpc_server.go): %s", data.Result)
	// TODO compression: compress result again here
	// TODO compression: here the only fields used are data and result at the end of the job, so service1->service2->sidecar
	// then the final part (sidecar) is where the data and result are present. These fields both need to be compressed.
	// TODO: only compress result and data if they are higher than a certain threshold, such as 300 bytes. This ensures the other communications through this function
	// are not compressed when 

	// TODO: so here only data and result fields need to be compressed, but only if the 

	// TODO: use lib.compress and lib.decompress


	ctx, span, err := lib.StartRemoteParentSpan(ctx, "sidecar SendData/func:", data.Traces)
	if err != nil {
		logger.Sugar().Warnf("Error starting span: %v", err)
	}
	defer span.End()

	if _, err := SendDataThroughAMQ(ctx, data, s); err != nil {
		logger.Sugar().Errorf("Callback Error: %v", err)
		return &pb.ContinueReceiving{ContinueReceiving: false}, nil
	}

	logger.Debug("Returning from SendData (to Microservice)")
	return &pb.ContinueReceiving{ContinueReceiving: false}, nil
}
