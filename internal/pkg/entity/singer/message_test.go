package singer

import (
	"reflect"
	"testing"
)

var (
	msg1 = &SchemaMessage{
		Type:          "SCHEMA",
		Stream:        "exchange_rate",
	}

	recordMsg = &RecordMessage{
		Type:          "RECORD",
		Stream:        "exchange_rate",
	}
)

func TestNewSchemaMessage(t *testing.T) {
	type args struct {
		stream        string
		schema        *Schema
		keyProperties []string
	}
	tests := []struct {
		name string
		args args
		want *SchemaMessage
	}{
		{
			name: "new then success",
			args: args{stream: "exchange_rate"},
			want: msg1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSchemaMessage(tt.args.stream, tt.args.schema, tt.args.keyProperties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSchemaMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRecordMessage(t *testing.T) {
	type args struct {
		stream string
		record *Record
	}
	tests := []struct {
		name string
		args args
		want *RecordMessage
	}{
		{
			name: "exchange rate then success",
			args: args{stream: "exchange_rate"},
			want: recordMsg,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRecordMessage(tt.args.stream, tt.args.record); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRecordMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}