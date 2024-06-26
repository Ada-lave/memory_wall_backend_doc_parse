package utils

import (
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/fumiama/go-docx"
)

func GetTextFromFile(file multipart.File, size int64) string {
	doc, err := docx.Parse(file, size)
	var buf strings.Builder
	if err != nil {
		panic(err)
	}



	for _, it := range doc.Document.Body.Items {
		switch it.(type) {
		case *docx.Paragraph:
			buf.WriteString(fmt.Sprintf("%v\n", it))
		}
	}

	return buf.String()
}