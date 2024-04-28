package just

import (
	"encoding/json"
	"fmt"
	"os"
)

// JsonParseType parse byte slice to specific type.
func JsonParseType[T any](bb []byte) (*T, error) {
	var target T
	if err := json.Unmarshal(bb, &target); err != nil {
		return nil, fmt.Errorf("unmarshal type: %w", err)
	}

	return &target, nil
}

// JsonParseTypeF parse json file into specific T.
func JsonParseTypeF[T any](filename string) (*T, error) {
	bb, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	return JsonParseType[T](bb)
}
