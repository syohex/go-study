package dmmapi

import (
	"testing"
	"time"
)

func Test_ageToDateString(t *testing.T) {
	tests := []struct {
		name string
		arg  int64
		want string
	}{
		{"1 year", 1, "1999-01-01"},
		{"1999 year", 1, "1-01-01"},
	}

	base := time.Date(2000, 1, 1, 0, 0, 0, 0, nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ageToDateString(base, tt.arg); got != tt.want {
				t.Errorf("ageToDateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
