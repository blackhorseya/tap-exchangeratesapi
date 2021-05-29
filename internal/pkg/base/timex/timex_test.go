package timex

import (
	"reflect"
	"testing"
	"time"
)

var (
	time1, _ = YYYYMMdd2Time("2021-01-01")
)

func TestTime2YYYYMMdd(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "time then yyyy-mm-dd",
			args: args{t: time1},
			want: "2021-01-01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Time2YYYYMMdd(tt.args.t); got != tt.want {
				t.Errorf("Time2YYYYMMdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYYYYMMdd2Time(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "2021-01-01 then time",
			args:    args{val: "2021-01-01"},
			want:    time1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := YYYYMMdd2Time(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("YYYYMMdd2Time() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YYYYMMdd2Time() got = %v, want %v", got, tt.want)
			}
		})
	}
}