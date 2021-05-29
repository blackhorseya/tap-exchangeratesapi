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
func WriteRecord() (string, error) {
	// todo: 2021-05-30|02:38|doggy|implement me
	panic("implement me")
}
