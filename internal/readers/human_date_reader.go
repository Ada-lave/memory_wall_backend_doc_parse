package readers

import (
	"memory_wall/lib/utils"
	"strings"
)

type HumanDateReader struct {
	text string
	dateTool utils.DateParseTool
	textFormatter utils.TextFormatter
}

func (HDR *HumanDateReader) GetBirthAndDeathDate() []string {
	for _, text := range strings.Split(HDR.text, "<br>") {
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

	placeAndDateOfСonscription := HDR.textFormatter.ExtractDataFromText(HDR.text, "Место и дата призыва:", "<br>")

	return placeAndDateOfСonscription
}

func NewHumanDateReader(text string) (HumanDateReader, error) {


	return HumanDateReader{
		text: text,
	}, nil
}
