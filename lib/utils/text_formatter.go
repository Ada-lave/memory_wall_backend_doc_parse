package utils

import "strings"

type TextFormatter struct{}

// Форматирует текст удаляя из него символ :
func (TF *TextFormatter) formatText(text string) string {
	text = strings.ReplaceAll(text, ":", "")
	text = strings.TrimSpace(text)

	return text
}


func (TF *TextFormatter) extractDataFromText(text string, sub string, sep string) string {
	if strings.Contains(text, sub) {
		militaryRank := strings.Split(strings.Split(text, sub)[1], sep)[0]
		formattedText := TF.formatText(militaryRank)

		return formattedText
	}
	return ""
}
