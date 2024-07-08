package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func WalkInDirAndFindAllFileNames(dirPath string) ([]string, error) {
	files, err := os.ReadDir(dirPath)
	var fileNames []string
	if err != nil {
		return []string{}, err
	}
	for _, file := range files {
		fileNames = append(fileNames, GetFileNameWithOutExt(file.Name()))
	}
	return fileNames, nil
}

func GetFileNameWithOutExt(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}


