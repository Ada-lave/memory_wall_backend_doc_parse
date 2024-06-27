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
			var hyperlink string = ""
			for _, pc := range it.(*docx.Paragraph).Children {
				fmt.Printf("%#v\n", pc)
				switch pc.(type) {
				case *docx.Hyperlink:
					hyperlink = fmt.Sprintf("%v",pc.(*docx.Hyperlink).Run.Children[0].(*docx.Text).Text)
					if hyperlink != "" {
						buf.WriteString(fmt.Sprintf(" %v ", hyperlink))
					}
				case *docx.Run:
					for _, text := range pc.(*docx.Run).Children {
						buf.WriteString(text.(*docx.Text).Text)
					}	
				}	
			}

			buf.WriteString("<br>")
		}
	}

	return buf.String()
}