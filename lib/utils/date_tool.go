package utils

import (
	"fmt"
	"time"
)

var months1 = map[string]time.Month{
	"января":   time.January,
    "февраля":  time.February,
    "марта":    time.March,
    "апреля":   time.April,
    "мая":      time.May,
    "июня":     time.June,
    "июля":     time.July,
    "августа":  time.August,
    "сентября": time.September,
    "октября":  time.October,
    "ноября":   time.November,
    "декабря":  time.December,
}

var months2 = map[string]time.Month{
	"январь":   time.January,
    "февраль":  time.February,
    "март":    time.March,
    "апрель":   time.April,
    "май":      time.May,
    "июнь":     time.June,
    "июль":     time.July,
    "августа":  time.August,
    "сентябрь": time.September,
    "октября":  time.October,
    "ноября":   time.November,
    "декабря":  time.December,
}

type DateParseTool struct {

}

func (DPT *DateParseTool) ParseDateFromString(date string) time.Time {
	stringTime, err := time.Parse("2006-01-02", date)

	if err == nil {
		return stringTime
	}

	fmt.Println(stringTime.UTC())

	return stringTime
}

func (DPT *DateParseTool) ParseFullDate(date string) (time.Time, error) {


	return time.Time{}, nil
}