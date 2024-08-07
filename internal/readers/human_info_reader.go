package readers

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/fumiama/go-docx"
	"io"
	"memory_wall/internal/http/memory_wall/models"
	"memory_wall/lib/utils"
	"mime/multipart"
	"strings"
)

type HumanInfoReader struct {
	utils.DocxReader
	textFormatter *utils.TextFormatter
	FullText string
}

func (HIR *HumanInfoReader) GetFullDescription(sep string) string {
	var buf strings.Builder

	for _, it := range HIR.Document.Document.Body.Items {
		switch it.(type) {
		case *docx.Paragraph:
			for _, pc := range it.(*docx.Paragraph).Children {
				switch pc.(type) {
				case *docx.Hyperlink:
					if len(pc.(*docx.Hyperlink).Run.Children) > 1 {
						buf.WriteString(fmt.Sprintf("%v", pc.(*docx.Hyperlink).Run.Children[0].(*docx.Text).Text))
					}
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
	HIR.FullText = buf.String()
	return buf.String()
}

func (HIR *HumanInfoReader) GetPlaceOfBirth() string {
	if HIR.FullText == "" {
		HIR.GetFullDescription("<br>")
	}
	placeOfBirth := HIR.textFormatter.ExtractDataFromText(HIR.FullText, "Место рождения", "<br>")

	return placeOfBirth
}

func (HIR *HumanInfoReader) GetImages() ([]models.HumanInfoImage, error) {

	var images []models.HumanInfoImage

	fileBytes, err := io.ReadAll(HIR.File)
	if err != nil {
		return nil, err
	}

	zipReader, err := zip.NewReader(bytes.NewReader(fileBytes), int64(len(fileBytes)))
	if err != nil {
		return nil, err
	}

	image := models.HumanInfoImage{}

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
			image.Name = imageName
			image.Data = imagesBytes
			images = append(images, image)

		}
	}

	return images, nil
}

func NewHumanInfoReader(file multipart.File, size int64) (HumanInfoReader, error) {
	var err error
	var humanInfoReader HumanInfoReader
	humanInfoReader.Document, err = docx.Parse(file, size)
	if err != nil {
		return HumanInfoReader{}, err
	}
	humanInfoReader.File = file

	return humanInfoReader, nil
}
