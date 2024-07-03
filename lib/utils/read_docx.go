package utils

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/fumiama/go-docx"
)

var textFormatter TextFormatter = TextFormatter{}
type DocxReader struct {
	Document *docx.Docx
	File     multipart.File
	FullText string
}

func (DR DocxReader) NewDocxReader(file multipart.File, size int64) (DocxReader, error) {
	var err error
	DR.Document, err = docx.Parse(file, size)
	if err != nil {
		return DocxReader{}, err
	}
	DR.File = file

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
	placeOfBirth := textFormatter.extractDataFromText(DR.FullText, "Место рождения", "<br>")

	return placeOfBirth
}

func (DR *DocxReader) GetBirthDate() string {
	if DR.FullText == "" {
		DR.GetFullDescription("<br>")
	}

	

	return ""
}

func (DR *DocxReader) GetPlaceAndDateOfСonscription() string {
	if DR.FullText == "" {
		DR.GetFullDescription("<br>")
	}
	placeAndDateOfСonscription := textFormatter.extractDataFromText(DR.FullText, "Место и дата призыва", "<br>")

	return placeAndDateOfСonscription
}

func (DR *DocxReader) GetMilitaryRank() string {
	if DR.FullText == "" {
		DR.GetFullDescription("<br>")
	}

	rank := textFormatter.extractDataFromText(DR.FullText, "Воинское звание, должность", "<br>")

	if len(rank) == 0 {
		rank = textFormatter.extractDataFromText(DR.FullText, "Воинское звание", "<br>")
	}

	return rank
}

func (DR *DocxReader) GetMedals() []string {
	var awards []string
	if DR.FullText == "" {
		DR.GetFullDescription("<br>")
	}
	if strings.Contains(DR.FullText, "Награды:") {
		textOfMedal := strings.Split(DR.FullText, "Награды:")[1]
		for _, medal := range strings.Split(textOfMedal, "<br>") {
			if medal != "" && strings.Contains(strings.ToLower(medal), "медаль") {
				awards = append(awards, medal)
			}
		}
	}
	return awards
}

func (DR *DocxReader) GetImages() (map[string][]byte, error) {
	fileBytes, err := io.ReadAll(DR.File)
	if err != nil {
		return nil, err
	}

	zipReader, err := zip.NewReader(bytes.NewReader(fileBytes), int64(len(fileBytes)))
	if err != nil {
		return nil, err
	}

	images := make(map[string][]byte)

	for _, zipFile := range zipReader.File {
		fmt.Println(zipFile.Name)
		if strings.Contains(zipFile.Name, "word/media") {
			imageReader, err := zipFile.Open()
			if err != nil {
				return nil, err
			}
			defer imageReader.Close()

			imagesBytes, err := io.ReadAll(imageReader)

			if err != nil {
				return nil, err
			}

			imageName := strings.Split(zipFile.Name, "/")[2]
			images[imageName] = imagesBytes
		}
	}

	return images, nil
}


