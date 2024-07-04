package utils

import "testing"

func Test_checkStringIsDate(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "Base test",
			args: "июня 1904",
			want: true,
		},
		{
			name: "Test with only years",
			args: "1904",
			want: true,
		},
		{
			name: "Full date test",
			args: "4 сентября 1983",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkStringIsDate(tt.args); got != tt.want {
				t.Errorf("checkStringIsDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
