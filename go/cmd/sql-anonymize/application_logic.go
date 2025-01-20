package main

import (
	"context"

	pb "github.com/Jorrit05/DYNAMOS/pkg/proto"
	"github.com/Jorrit05/DYNAMOS/pkg/lib"

	"google.golang.org/protobuf/proto"
	"github.com/gogo/protobuf/jsonpb"
	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"
	"google.golang.org/protobuf/types/known/structpb"
)

// This is the function being called by the last microservice
func handleSqlDataRequest(ctx context.Context, msComm *pb.MicroserviceCommunication) (context.Context, error) {
	ctx, span := trace.StartSpan(ctx, "anonymize: handleSqlDataRequest")
	defer span.End()

	logger.Info("anonymize Start handleSqlDataRequest")

	sqlDataRequest := &pb.SqlDataRequest{}
	if err := msComm.OriginalRequest.UnmarshalTo(sqlDataRequest); err != nil {
		logger.Sugar().Errorf("Failed to unmarshal sqlDataRequest message: %v", err)
		return ctx, err
	}

	// TODO: currently in DYNAMOS anonymize is not used anymore I believe, however, still decompress logic is 
	// TODO: added, but when used again, test requesting data to see if decompression works just to be sure
	// Get and decompress the value from the data
	decompressedData, err := lib.GetDecompressedValue(msComm.Data)
    if err != nil {
        logger.Sugar().Errorf("Failed to decompress data: %s", err)
		return ctx, err
    } else {
        logger.Sugar().Debugf("*********Decompressed data size: %d", len(decompressedData))
    }
    // Unmarshal the decompressed data into a structpb.Struct
    decompressedStruct := &structpb.Struct{}
    if err := proto.Unmarshal(decompressedData, decompressedStruct); err != nil {
        logger.Sugar().Errorf("Failed to unmarshal decompressed data: %s", err)
		return ctx, err
    }

	anonymizeDatesInStruct(decompressedStruct)

	if sqlDataRequest.Options["graph"] {
		// jsonString, _ := json.Marshal(decompressedStruct)
		// msComm.Result = jsonString

		m := &jsonpb.Marshaler{}
		jsonString, _ := m.MarshalToString(decompressedStruct)
		msComm.Result = []byte(jsonString)

		return ctx, nil
	}

	msComm.Traces["binaryTrace"] = propagation.Binary(span.SpanContext())
	return ctx, nil
}

func anonymizeDatesInStruct(data *structpb.Struct) {
	fieldsToAnonymize := []string{"Ingdatdv", "Gebdat"}

	for _, field := range fieldsToAnonymize {
		fieldValue, ok := data.GetFields()[field]
		if !ok {
			continue
		}

		listValues := fieldValue.GetListValue().GetValues()
		for index, value := range listValues {
			stringValue := value.GetStringValue()
			if len(stringValue) >= 4 {
				// Create a new value with modified string and set it in the slice
				newValue := structpb.NewStringValue(stringValue[:4])
				listValues[index] = newValue
			}
		}
	}
}
