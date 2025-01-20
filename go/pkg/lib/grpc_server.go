package lib

import (
	"context"
	"fmt"

	pb "github.com/Jorrit05/DYNAMOS/pkg/proto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SharedServer struct {
	pb.UnimplementedMicroserviceServer
	pb.UnimplementedHealthServer
	pb.UnimplementedGenericServer
	ServiceName string
	Callback    func(ctx context.Context, data *pb.MicroserviceCommunication) error
}

func (s *SharedServer) Check(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{Status: pb.HealthCheckResponse_SERVING}, nil
}

func (s *SharedServer) InitTracer(ctx context.Context, in *pb.ServiceName) (*emptypb.Empty, error) {

	_, err := InitTracer(in.ServiceName + "/sidecar")
	if err != nil {
		logger.Sugar().Fatalf("Failed to create ocagent-exporter: %v", err)
	}

	return &emptypb.Empty{}, nil
}

// SharedServer implementation for SendData, this is called by the previous microservice
// or by the sidecar itself if a message is received from rabbitMQ
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
func (s *SharedServer) SendData(ctx context.Context, data *pb.MicroserviceCommunication) (*pb.ContinueReceiving, error) {
	logger.Sugar().Debugf("Starting (to next MS) lib.SendData: %v", data.RequestMetadata.DestinationQueue)

	// Logging of data send for compression testing
	logger.Sugar().Debugf("**********************Microservice communication type (in go/pkg/lib/grpc_server.go): %s", data.Type)
	logger.Sugar().Debugf("**********************Microservice communication request type (in go/pkg/lib/grpc_server.go): %s", data.RequestType)
	// Convert the google.protobuf.Struct to a JSON string
    dataJSON, err := protojson.Marshal(data.Data)
    if err != nil {
        logger.Sugar().Errorf("Failed to marshal data to JSON: %s", err)
    }
	logger.Sugar().Debugf("**********************Microservice communication data (in go/pkg/lib/grpc_server.go): %s", dataJSON)
	logger.Sugar().Debugf("**********************Microservice communication result (in go/pkg/lib/grpc_server.go): %s", data.Result)
	// TODO compression: compress result again here
	// TODO compression: found here that this SendData is used by the query services, such as sql-query and sql-algorithm.
	// these services ONLY change the 'data' field, the 'result field here is still empty all the time.

	// TODO: so here only the data field needs to be compressed
	// TODO: how to do decompressing, figure that out. But first add compression.
	// Always compress the data field if it is not nil (data field contains the data results between the different services)
	// Compress the `data.Data` field if it is not nil
	if data.Data != nil {
		// Serialize the Struct to a byte slice (required to extract bytes from the Struct type)
		serializedData, err := proto.Marshal(data.Data)
		logger.Sugar().Debugf("**********************Microservice communication serialized data size (in go/pkg/lib/grpc_server.go): %d", len(serializedData))
		if err != nil {
			logger.Sugar().Errorf("Failed to serialize data.Data field: %s", err)
		} else {
			// Compress the serialized data using the compress function from compression.go in the same package as this file
			compressedData, err := Compress(serializedData)
			if err != nil {
				logger.Sugar().Errorf("Failed to compress data.Data field: %s", err)
			} else {
				logger.Sugar().Debugf("**********************Microservice communication data compressed size (in go/pkg/lib/grpc_server.go): %d", len(compressedData))
				// Store the compressed data as raw bytes string value in a new struct
				compressedStruct := &structpb.Struct{
					Fields: map[string]*structpb.Value{
						"compressed_data": {Kind: &structpb.Value_StringValue{StringValue: string(compressedData)}},
					},
				}
				// Replace the original Struct with the compressed version
				data.Data = compressedStruct
			}
		}
	}

	ctx, span, err := StartRemoteParentSpan(ctx, fmt.Sprintf("%s SendData/func:", s.ServiceName), data.Traces)
	if err != nil {
		logger.Sugar().Warnf("Error starting span: %v", err)
	}
	defer span.End()

	if data.Type == "microserviceCommunication" {
		err = s.Callback(ctx, data)
		if err != nil {
			logger.Sugar().Errorf("Failed to process message: %v", err)
		}
	} else {
		logger.Sugar().Errorf("Unknown message type: %v", data.Type)
		return &pb.ContinueReceiving{ContinueReceiving: false}, fmt.Errorf("unknown message type: %s", data.Type)
	}
	return &pb.ContinueReceiving{ContinueReceiving: false}, err
}
