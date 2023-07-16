package main

import (
	"context"
	"fmt"

	"github.com/Jorrit05/DYNAMOS/pkg/lib"
	pb "github.com/Jorrit05/DYNAMOS/pkg/proto"
)

func handleIncomingMessages(ctx context.Context, grpcMsg *pb.RabbitMQMessage) error {

	ctx, span, err := lib.StartRemoteParentSpan(ctx, serviceName+"/func: handleSidecarMessages", grpcMsg.Trace)
	if err != nil {
		logger.Sugar().Errorf("Error starting span: %v", err)
		return err
	}
	defer span.End()

	logger.Sugar().Debugw("Type:", "MessageType", grpcMsg.Type)

	switch grpcMsg.Type {
	case "validationResponse":
		validationResponse := &pb.ValidationResponse{}
		if err := grpcMsg.Body.UnmarshalTo(validationResponse); err != nil {
			logger.Sugar().Fatalf("Failed to unmarshal message: %v", err)
		}
		mutex.Lock()
		// Look up the corresponding channel in the request map
		validationChannel, ok := validationMap[validationResponse.User.Id]
		mutex.Unlock()

		if ok {
			logger.Sugar().Info("Sending validation to channel")
			// Send a signal on the channel to indicate that the response is ready
			validationChannel <- validation{response: validationResponse, localContext: ctx}
		} else {
			logger.Sugar().Errorw("unknown validation response", "GUID", validationResponse.User.Id)
		}

	default:
		logger.Sugar().Errorf("Unknown message type: %s", grpcMsg.Type)
		return fmt.Errorf("unknown message type: %s", grpcMsg.Type)
	}

	return nil
}
