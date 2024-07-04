package utils

import (
	"github.com/fumiama/go-docx"
	"mime/multipart"
)

type DocxReader struct {
	Document *docx.Docx
	File     multipart.File
	FullText string
}

func NewDocxReader(file multipart.File, size int64) (DocxReader, error) {
	var err error
	var dr DocxReader
	dr.Document, err = docx.Parse(file, size)
	if err != nil {
		return DocxReader{}, err
	}
	dr.File = file

	return dr, nil
}
