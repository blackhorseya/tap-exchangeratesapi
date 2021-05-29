package singer

import (
	"github.com/blackhorseya/tap-exchangeratesapi/internal/pkg/base/timex"
	"github.com/blackhorseya/tap-exchangeratesapi/internal/pkg/model"
)

// Record declare record struct
type Record struct {
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
	Base  map[string]float64 `json:"base"`
}

// NewRecordFromResp serve caller to create a Record
func NewRecordFromResp(resp *model.APIResponse) *Record {
	time, _ := timex.YYYYMMdd2Time(resp.Date)

	return &Record{
		Date:  timex.Time2RFC(time),
		Rates: resp.Rates,
		Base:  map[string]float64{resp.Base: 1.0},
	}
}
