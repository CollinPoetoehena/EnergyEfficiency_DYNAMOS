package lib

import (
	"bytes"
	"compress/gzip"
	"io"
	"encoding/base64"

	"google.golang.org/protobuf/types/known/structpb"
)

// Compress compresses a given byte slice using gzip.
func Compress(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	_, err := writer.Write(input)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decompress data
func Decompress(input []byte) ([]byte, error) {
    buf := bytes.NewReader(input)
    reader, err := gzip.NewReader(buf)
    if err != nil {
        return nil, err
    }
    defer reader.Close()

    return io.ReadAll(reader)
}

// Helper function to get and decompress the value from the data
func GetDecompressedValue(data *structpb.Struct) ([]byte, error) {
	// Decode the Base64-encoded compressed data
	compressedBase64 := data.Fields["compressed_data"].GetStringValue()
	compressedBytes, err := base64.StdEncoding.DecodeString(compressedBase64)
	if err != nil {
		logger.Sugar().Errorf("Failed to decode Base64 data: %s", err)
		return nil, err
	}

	// Decompress the data
	decompressedData, err := Decompress(compressedBytes)
	if err != nil {
		logger.Sugar().Errorf("Failed to decompress data: %s", err)
		return nil, err
	}

    return decompressedData, nil
}

// // Helper function to get the compressed value from the data
// func GetCompressedValue(data *structpb.Struct) string {
//     return data.Fields["compressed_data"].GetStringValue()
// }