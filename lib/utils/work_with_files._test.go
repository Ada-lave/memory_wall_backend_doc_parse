package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileNameWithOutExt(t *testing.T) {
	testCases := []struct {
		Name string
		Filename string
		Want string
	}{
		{
			Name: "Base test",
			Filename: "Alex Batman.docx",
			Want: "Alex Batman",
		},
		{
			Name: "Test without extension",
			Filename: "Alex Batman Superman",
			Want: "Alex Batman Superman",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			res := GetFileNameWithOutExt(tc.Filename)
			assert.Equal(t, tc.Want, res)
		})
	}
}