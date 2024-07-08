package readers

import (
	"memory_wall/lib/utils"
	"mime/multipart"
	"strings"
)

type HumanDateReader struct {
	HumanInfoReader
}

func (HDR *HumanDateReader) GetBirthAndDeathDate() []string {
	if HDR.FullText == "" {
		HDR.GetFullDescription("<br>")
	}

	for _, text := range strings.Split(HDR.FullText, "<br>") {
		if len(text) != 0 {
			formattedText := strings.ReplaceAll(text, "(", "")
			formattedText = strings.ReplaceAll(formattedText, ")", "")
			formattedText = strings.ReplaceAll(formattedText, "-", "-")
			formattedText = strings.ReplaceAll(formattedText, "–", "-")
			formattedText = strings.ReplaceAll(formattedText, "—", "-")

			dates := strings.Split(formattedText, "-")
			if len(dates) == 2 {
				dates[0] = strings.Trim(dates[0], " ")
				dates[1] = strings.Trim(dates[1], " ")
				if utils.CheckStringIsDate(dates[0]) || utils.CheckStringIsDate(dates[1]) {

					if date1, err := HDR.dateTool.ParseDateFromString(dates[0]); err == nil {
						dates[0] = date1
					}
					if date2, err := HDR.dateTool.ParseDateFromString(dates[1]); err == nil {
						dates[1] = date2
					}

					return dates
				}
			}
		}
	}

	return []string{}
}


func (HDR *HumanDateReader) GetPlaceAndDateOfСonscription() string {
	if HDR.FullText == "" {
		HDR.GetFullDescription("<br>")
	}

	placeAndDateOfСonscription := HDR.textFormatter.ExtractDataFromText(HDR.FullText, "Место и дата призыва", "<br>")

	return placeAndDateOfСonscription
}

func NewHumanDateReader(file multipart.File, size int64) (HumanDateReader, error) {
	var humanInfoReader HumanInfoReader

	humanInfoReader, err := NewHumanInfoReader(file, size)
	if err != nil {
		return HumanDateReader{}, err
	}

	return HumanDateReader{
		HumanInfoReader: humanInfoReader,
	}, nil
}