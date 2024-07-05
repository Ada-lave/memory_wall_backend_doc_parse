package utils

import "testing"

func TestCheckStringIsDate(t *testing.T) {
	tests := []struct {
		name string
		date string
		want bool
	}{
		{
			name: "Basic test with normal date",
			date: "27.05.2005",
			want: true,
		},
		{
			name: "Test case with month and year",
			date: "июль 2005",
			want: true,
		},
		{
			name: "Test case with only year only",
			date: "1902",
			want: true,
		},
		{
			name: "Test case with typical date",
			date: "17 августа 1967",
			want: true,
		},
		{
			name: "Test case with no date",
			date: "HACKER",
			want: false,
		},
		{
			name: "Test case with empty string",
			date: "",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckStringIsDate(tt.date); got != tt.want {
				t.Errorf("CheckStringIsDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
