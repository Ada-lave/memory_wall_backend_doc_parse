package readers

import (
	"fmt"
	"memory_wall/lib/utils"
	"strings"
)

type HumanFIOReader struct {
	text string
	textFormatter *utils.TextFormatter
}

func NewHumanFIOReader(text string) (HumanFIOReader, error) {
	return HumanFIOReader{
		text: text,
	}, nil
}

func (HFR *HumanFIOReader) GetFIO() []string {
	var splittedText []string

	text  := strings.ReplaceAll(HFR.text, " <br>", "<br>")

	// Избавляемся от пустых слов
	for _, word := range strings.Split(text, "<br>") {
		if word != "" {
			splittedText = append(splittedText, word)
		}
	}

	if len(splittedText) < 1 {
		return []string{}
	}
	fio := strings.Split(splittedText[0], " ")

	switch len(fio) {
	case 3:
		// full text on one line
		fio = TrimAllWords(fio)

		capitalizedText := HFR.textFormatter.CapitalizeWords(strings.Join(fio, " "))
		return strings.Split(capitalizedText, " ")
	case 2:
		// full text on different line
		if len(splittedText) > 1 && !utils.CheckStringIsDate(splittedText[1]) {
			fio = append(fio, splittedText[1])
			fio = TrimAllWords(fio)

			capitalizedText := HFR.textFormatter.CapitalizeWords(strings.Join(fio, " "))
			return strings.Split(capitalizedText, " ")
		} else {
			fio = TrimAllWords(fio)

			capitalizedText := HFR.textFormatter.CapitalizeWords(strings.Join(fio, " "))
			return strings.Split(capitalizedText, " ")
		}
	case 1:
		// All on different line
		if len(splittedText) > 2 && 
			!utils.CheckStringIsDate(splittedText[0]) && 
			!utils.CheckStringIsDate(splittedText[1]) && 
			(!utils.CheckStringIsDate(splittedText[2])) {

			fio = append(fio, splittedText[1])
			fio = append(fio, splittedText[2])
			fio = TrimAllWords(fio)

			capitalizedText := HFR.textFormatter.CapitalizeWords(strings.Join(fio, " "))
			fmt.Printf("%#v\n", fio)
			return strings.Split(capitalizedText, " ")
		} else if len(splittedText) > 1 && !utils.CheckStringIsDate(splittedText[0]) && !utils.CheckStringIsDate(splittedText[1]) {
			splittedText = strings.Split(splittedText[1], " ")

			if len(splittedText) > 1 {
				fio = append(fio, splittedText[0])
				fio = append(fio, splittedText[1])
				fio = TrimAllWords(fio)

				capitalizedText := HFR.textFormatter.CapitalizeWords(strings.Join(fio, " "))
				return strings.Split(capitalizedText, " ")
			}
		}
	}

	return []string{}
}

func TrimWhitespaces(text string) string {
	text = strings.Trim(text, " ")

	return text
}

func TrimAllWords(text []string) []string {
	for i, word := range text {
		text[i] = TrimWhitespaces(word)
	}

	return text
}
