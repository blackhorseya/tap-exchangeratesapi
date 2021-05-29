package singerutils

import (
	"encoding/json"

	"github.com/blackhorseya/tap-exchangeratesapi/internal/pkg/entity/singer"
)

// WriteSchema serve caller to print schema message
func WriteSchema(message *singer.SchemaMessage) (string, error) {
	data, err := json.Marshal(message)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// WriteRecord serve caller to return record message
func WriteRecord(message *singer.RecordMessage) (string, error) {
	data, err := json.Marshal(message)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
