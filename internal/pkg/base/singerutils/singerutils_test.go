package singerutils

import (
	"testing"

	"github.com/blackhorseya/tap-exchangeratesapi/internal/pkg/base/timex"
	"github.com/blackhorseya/tap-exchangeratesapi/internal/pkg/entity/singer"
)

var (
	msg = singer.NewSchemaMessage("exchange_rate", singer.NewSchema(), []string{"date"})

	time1, _ = timex.YYYYMMdd2Time("2021-01-01")

	timeStr1 = timex.Time2RFC(time1)

	record1 = &singer.Record{
		Date:  timeStr1,
		Rates: map[string]float64{"USD": 1.307716},
		Base:  map[string]float64{"EUR": 1.0},
	}

	msg1 = singer.NewRecordMessage("exchange_rate", record1)
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

func TestWriteRecord(t *testing.T) {
	type args struct {
		message *singer.RecordMessage
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "print message then success",
			args:    args{message: msg1},
			want:    "{\"type\":\"RECORD\",\"stream\":\"exchange_rate\",\"record\":{\"date\":\"2021-01-01T00:00:00.000000Z\",\"rates\":{\"USD\":1.307716},\"base\":{\"EUR\":1}},\"time_extracted\":\"\"}",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WriteRecord(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WriteRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}