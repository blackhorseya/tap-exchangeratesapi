package timex

import "time"

// YYYYMMdd2Time serve caller to given yyyy-MM-dd format string to parse time
func YYYYMMdd2Time(val string) (time.Time, error) {
	layout := "2006-01-02"
	loc, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return time.Time{}, err
	}

	t, err := time.ParseInLocation(layout, val, loc)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

// Time2YYYYMMdd serve caller to given time to string layout YYYY-MM-dd
func Time2YYYYMMdd(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}
