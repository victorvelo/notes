package mapper

import (
	"encoding/json"
	"fmt"
)

func UnmarshalData[T any](data []byte) (T, error) {
	var result T

	if err := json.Unmarshal(data, &result); err != nil {
		return result, fmt.Errorf("deserialization of response body: %w", err)
	}

	return result, nil
}
