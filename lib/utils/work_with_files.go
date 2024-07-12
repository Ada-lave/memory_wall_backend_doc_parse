package utils

import (
	"path/filepath"
	"strings"
)

func GetFileNameWithOutExt(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}
