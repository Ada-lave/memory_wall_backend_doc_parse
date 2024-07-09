package utils

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

type TextFormatter struct{}

// Форматирует текст удаляя из него символ ":"
func (TF *TextFormatter) FormatText(text string) string {
	text = strings.ReplaceAll(text, ":", "")
	text = strings.TrimSpace(text)

	return text
}

func (TF *TextFormatter) ExtractDataFromText(text string, sub string, sep string) string {
	if strings.Contains(text, sub) {
		militaryRank := strings.Split(strings.Split(text, sub)[1], sep)[0]
		formattedText := TF.FormatText(militaryRank)

		return formattedText
	}
	return ""
}

func (TF *TextFormatter) CapitalizeWords(text string) string {
	text = strings.ToLower(text)
	var buf strings.Builder
	splittedText := strings.Split(text, " ")
	for i, word := range splittedText {
		r, size := utf8.DecodeRuneInString(word)
		if r == utf8.RuneError {
			buf.WriteString("err")
		} else {
			buf.WriteString(string(unicode.ToUpper(r)) + word[size:])
		}

		if i != len(splittedText)-1 {
			buf.WriteString(" ")
		}
	}

	return buf.String()
}

func CheckStringIsDate(date string) bool {
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
