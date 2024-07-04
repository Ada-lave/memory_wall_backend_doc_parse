package utils

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"regexp"
	"strings"
	"github.com/fumiama/go-docx"
)

var textFormatter TextFormatter = TextFormatter{}

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

func (DR *DocxReader) GetFIO() []string {
	var FIO []string
	if DR.FullText == "" {
		DR.GetFullDescription("<br>")
	}

	var data []string 
	for _, text := range strings.Split(DR.FullText, "<br>") {
		if len(text) != 0 {
			data = append(data, text)
		}
 	}
	FIO = strings.Split(data[0], " ")

	if len(FIO) != 3 {
		splitedNames := strings.Split(data[1], " ")
		FIO = append(FIO, splitedNames[0])
		FIO = append(FIO, splitedNames[1])
	}
	return FIO 
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

func (DR *DocxReader) GetBirthDate() []string {
	if DR.FullText == "" {
		DR.GetFullDescription("<br>")
	}

	for _, text := range strings.Split(DR.FullText, "<br>") {
		if len(text) != 0 {
			formattedText := strings.ReplaceAll(text, "(", "")
			formattedText = strings.ReplaceAll(formattedText, ")", "")
			formattedText = strings.ReplaceAll(formattedText, "-", "-")
			formattedText = strings.ReplaceAll(formattedText, "–", "-")

			dates := strings.Split(formattedText, "-")
			if len(dates) == 2 {
				dates[0] = strings.Trim(dates[0], " ")
				dates[1] = strings.Trim(dates[1], " ")
				if checkStringIsDate(dates[0]) || checkStringIsDate(dates[1]) {

					return dates
				}
			}
		}
	}

	return []string{}
}

func checkStringIsDate(date string) bool {
	var datePattern1 = `^\d{1,2} (январ[ья]|феврал[ья]|март[а]|апрел[ья]|ма[йя]|июн[ья]|июл[ья]|август[а]|сентябр[ья]|октябр[ья]|ноябр[ья]|декабр[ья]) \d{4}$`
	var datePattern2 = `^(январ[ья]|феврал[ья]|март[а]|апрел[ья]|ма[йя]|июн[ья]|июл[ья]|август[а]|сентябр[ья]|октябр[ья]|ноябр[ья]|декабр[ья]) \d{4}$`
	var datePattern3 = `\d{4}`
	var datePattern4 = `^d{1,2}\d{1,2}.\d{4}$`

	matched1, _ := regexp.MatchString(datePattern1, date)
	matched2, _ := regexp.MatchString(datePattern2, date)
	matched3, _ := regexp.MatchString(datePattern3, date)
	matched4, _ := regexp.MatchString(datePattern4, date)

	return matched1 || matched2 || matched3 || matched4
}
