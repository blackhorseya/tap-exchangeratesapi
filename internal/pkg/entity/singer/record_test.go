package singer

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/tap-exchangeratesapi/internal/pkg/base/timex"
	"github.com/blackhorseya/tap-exchangeratesapi/internal/pkg/model"
)

var (
	time1, _ = timex.YYYYMMdd2Time("2021-01-01")

	timeStr1 = timex.Time2RFC(time1)

	record1 = &Record{
		Date:  timeStr1,
		Rates: map[string]float64{"USD": 1.307716},
		Base:  map[string]float64{"EUR": 1.0},
	}

	resp1 = &model.APIResponse{
		Success:    true,
		Timestamp:  1363478399,
		Historical: true,
		Base:       "EUR",
		Date:       "2021-01-01",
		Rates: map[string]float64{
			"USD": 1.307716,
		},
	}
)

func TestNewRecordFromResp(t *testing.T) {
	type args struct {
		resp *model.APIResponse
	}
	tests := []struct {
		name string
		args args
		want *Record
	}{
		{
			name: "response then record",
			args: args{resp: resp1},
			want: record1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRecordFromResp(tt.args.resp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRecordFromResp() = %v, want %v", got, tt.want)
			}
		})
	}
}
