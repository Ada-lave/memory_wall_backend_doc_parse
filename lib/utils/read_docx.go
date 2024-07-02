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
    placeOfBirth := extractDataFromText(DR.FullText, "Место рождения", "<br>")

    return placeOfBirth
}

func (DR *DocxReader) GetPlaceAndDateOfСonscription() string {
	if DR.FullText == "" {
		DR.GetFullDescription("<br>")
	}
	placeAndDateOfСonscription := extractDataFromText(DR.FullText, "Место и дата призыва", "<br>")

	return placeAndDateOfСonscription
}

func (DR *DocxReader) GetMilitaryRank() string {
	if DR.FullText == "" {
		DR.GetFullDescription("<br>")
	}

	rank := extractDataFromText(DR.FullText, "Воинское звание, должность", "<br>")

	return rank
}

func (DR *DocxReader) GetMedals() []string {
	var awards []string
	if DR.FullText == "" {
		DR.GetFullDescription("<br>")
	}
	if strings.Contains(DR.FullText, "Награды:") {
		textOfMedal := strings.Split(DR.FullText, "Награды:")[1]
		for _ ,medal := range strings.Split(textOfMedal, "<br>") {
			if medal != "" && strings.Contains(strings.ToLower(medal), "медаль") {
				awards = append(awards, medal)
			}
		}
	}
	return awards
}

func (DR *DocxReader) GetImages() []byte {
	for _, item := range DR.Document.Document.Body.Items {
		switch paragraph := item.(type) {
		case *docx.Paragraph:
			for _, ch := range paragraph.Children {
				switch run := ch.(type) {
				case *docx.Run:
					for _, ch := range run.Children {
						switch draw := ch.(type) {
						case *docx.Drawing:
							fmt.Printf("Image: %#v\n",draw.Inline.Graphic.GraphicData.Pic.BlipFill.Blip)
						}
					}
				}
			}
		}
		
	}

	return []byte{}
}

func formatText(text string) string {
	text = strings.ReplaceAll(text, ":", "")
	text = strings.TrimSpace(text)

	return text
}

func extractDataFromText(text string, sub string, sep string) string {
	if strings.Contains(text, sub) {
		militaryRank := strings.Split(strings.Split(text, sub)[1], sep)[0]
		formattedText := formatText(militaryRank)

		return formattedText
	}
	return ""
}
