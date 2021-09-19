package poker

import (
	"testing"
	"time"
)

func TestBlindAlerterFunc_ScheduleAlertAt(t *testing.T) {
	type args struct {
		duration time.Duration
		amount   int
	}
	tests := []struct {
		name string
		a    BlindAlerterFunc
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.ScheduleAlertAt(tt.args.duration, tt.args.amount)
		})
	}
}

func TestStdOutAlerter(t *testing.T) {
	type args struct {
		duration time.Duration
		amount   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StdOutAlerter(tt.args.duration, tt.args.amount)
		})
	}
}
