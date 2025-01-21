package lib

import (
	"context"
	"fmt"
	"encoding/base64"

	pb "github.com/Jorrit05/DYNAMOS/pkg/proto"
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

	// Compress the data field if it is not nil (data field contains the data results between the different services)
	if data.Data != nil {
		// Serialize the Struct to a byte slice (required to extract bytes from the Struct type)
		serializedData, err := proto.Marshal(data.Data)
		if err != nil {
			logger.Sugar().Errorf("Failed to serialize data.Data field: %s", err)
		} else {
			// Compress the serialized data using the compress function from compression.go in the same package as this file
			compressedData, err := Compress(serializedData)
			if err != nil {
				logger.Sugar().Errorf("Failed to compress data.Data field: %s", err)
			} else {
				// Encode compressed data in Base64
				encodedData := base64.StdEncoding.EncodeToString(compressedData)

				// Store the compressed and encoded data as raw bytes string value in a new struct
				// Base64 avoids problems like: grpc: error while marshaling: string field contains invalid UTF-8
				compressedStruct := &structpb.Struct{
					Fields: map[string]*structpb.Value{
						"compressed_data": {Kind: &structpb.Value_StringValue{StringValue: encodedData}},
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
