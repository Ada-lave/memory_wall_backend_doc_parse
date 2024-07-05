package utils

import (
	"reflect"
	"testing"
	"time"
)

func TestDateParseTool_ParseDateFromString(t *testing.T) {
	tests := []struct {
		name    string
		DPT     *DateParseTool
		date    string
		want    string
		wantErr bool
	}{
		{
			name: "Base time test",
			date: "2023-10-12",
			want: time.Date(2023, 10, 12, 00, 00, 00, 00, time.UTC).String(),
		},
		{
			name: "Test with moth like text case 1",
			date: "25 июля 2024",
			want: time.Date(2024, 07, 25, 00, 00, 00, 00, time.UTC).String(),
		},
		{
			name: "Test with moth like text case 2",
			date: "25 июль 2024",
			want: time.Date(2024, 07, 25, 00, 00, 00, 00, time.UTC).String(),
		},
		{
			name: "Test month and year only",
			date: "июль 2024",
			want: time.Date(2024, 07, 0, 00, 00, 00, 00, time.UTC).String(),
		},
		{
			name: "Test year only",
			date: "2024",
			want: time.Date(2024, 0, 0, 00, 00, 00, 00, time.UTC).String(),
		},
		{
			name: "Test year with letter only",
			date: "2024г",
			want: time.Date(2024, 0, 0, 00, 00, 00, 00, time.UTC).String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DPT := &DateParseTool{}
			got, err := DPT.ParseDateFromString(tt.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("DateParseTool.ParseDateFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateParseTool.ParseDateFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
