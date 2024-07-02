package utils

import (
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/fumiama/go-docx"
)

type DocxReader struct {
	Document *docx.Docx
	FullText string
}

func (DR DocxReader) NewDocxReader(file multipart.File, size int64) (DocxReader, error) {
	var err error
	DR.Document, err = docx.Parse(file, size)
	if err != nil {
		return DocxReader{}, err
	}

	return DR, nil
}

func (DR *DocxReader) GetFullDescription(sep string) string { 
	var buf strings.Builder

	for _, it := range DR.Document.Document.Body.Items { 
		switch it.(type) {
        case *docx.Paragraph:
            for _, pc := range it.(*docx.Paragraph).Children {
                switch pc.(type) {
                case *docx.Hyperlink:
                    buf.WriteString(fmt.Sprintf("%v", pc.(*docx.Hyperlink).Run.Children[0].(*docx.Text).Text))
                case *docx.Run:
                    for _, text := range pc.(*docx.Run).Children {
						switch t := text.(type) {
						case *docx.Text:
							buf.WriteString(t.Text)
						}
                        
                    }
                }
            }
            buf.WriteString(sep)
        }
	}
	DR.FullText = buf.String()
	return buf.String()
}

func (DR *DocxReader) GetPlaceOfBirth() string {
	if DR.FullText == "" {
		DR.GetFullDescription("<br>")
	}
	if strings.Contains(DR.FullText, "Место рождения") {
		placeOfBirth := strings.Split(strings.Split(DR.FullText, "Место рождения")[1], "<br>")[0]
		formattedText := formatText(placeOfBirth)
		return formattedText
	} 
	return ""
}

func (DR *DocxReader) GetPlaceAndDateOfСonscription() string {
	if DR.FullText == "" {
		DR.GetFullDescription("<br>")
	}

	if strings.Contains(DR.FullText, "Место и дата призыва") {
        placeAndDateOfСonscription := strings.Split(strings.Split(DR.FullText, "Место и дата призыва")[1], "<br>")[0]
        formattedText := formatText(placeAndDateOfСonscription)
        return formattedText
    }

	return ""
}

func formatText(text string) string {
	text = strings.ReplaceAll(text, ":", "")
	text = strings.TrimSpace(text)

	return text
}