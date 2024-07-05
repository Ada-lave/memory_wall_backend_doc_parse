package utils

import (
	"reflect"
	"testing"
	"time"
)

func TestParseDateFromString(t *testing.T) {
	tests := []struct {
		name string
		date string
		want time.Time
	}{
		{
			name: "Base time test",
			date: "2023-10-12",
			want: time.Date(2023, 10, 12, 00, 00, 00, 00, time.UTC),
		},
		{
			name: "Test with moth like text",
			date: "25 июля 2024",
			want: time.Date(2024, 07, 00, 00, 00, 00, 00, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDateFromString(tt.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDateFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
