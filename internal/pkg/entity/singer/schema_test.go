package singer

import (
	"reflect"
	"testing"
)

var (
	schema1 = &Schema{
		Type: "object",
		Properties: map[string]*Property{
			"date": {Type: "string", Format: "date-time"},
		},
	}
)

func TestNewSchema(t *testing.T) {
	tests := []struct {
		name string
		want *Schema
	}{
		{
			name: "new schema then success",
			want: schema1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSchema(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}
