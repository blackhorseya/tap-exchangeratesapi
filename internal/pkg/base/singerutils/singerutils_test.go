package singerutils

import (
	"testing"

	"github.com/blackhorseya/tap-exchangeratesapi/internal/pkg/entity/singer"
)

var (
	msg = singer.NewSchemaMessage("exchange_rate", singer.NewSchema(), []string{"date"})
)

func TestWriteSchema(t *testing.T) {
	type args struct {
		message *singer.SchemaMessage
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "print message then success",
			args:    args{message: msg},
			want:    "{\"type\":\"SCHEMA\",\"stream\":\"exchange_rate\",\"schema\":{\"type\":\"object\",\"properties\":{\"date\":{\"type\":\"string\",\"format\":\"date-time\"}}},\"key_properties\":[\"date\"],\"selected\":false}",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WriteSchema(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WriteSchema() got = %v, want %v", got, tt.want)
			}
		})
	}
}