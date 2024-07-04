package utils

import "strings"

type TextFormatter struct{}

// TODO: Перенести функционал работы с текстом в отдельный класс
func (TF *TextFormatter) formatText(text string) string {
	text = strings.ReplaceAll(text, ":", "")
	text = strings.TrimSpace(text)

	return text
}

// TODO: Перенести функционал работы с текстом в отдельный класс
func (TF *TextFormatter) extractDataFromText(text string, sub string, sep string) string {
	if strings.Contains(text, sub) {
		militaryRank := strings.Split(strings.Split(text, sub)[1], sep)[0]
		formattedText := TF.formatText(militaryRank)

		return formattedText
	}
	return ""
}
