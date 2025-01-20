package lib

import (
	"bytes"
	"compress/gzip"
	"io"

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
    compressedValue := data.Fields["compressed_data"].GetStringValue()
    return Decompress([]byte(compressedValue))
}

// // Helper function to get the compressed value from the data
// func GetCompressedValue(data *structpb.Struct) string {
//     return data.Fields["compressed_data"].GetStringValue()
// }