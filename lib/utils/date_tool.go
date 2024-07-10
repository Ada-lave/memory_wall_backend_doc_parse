package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
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
	"март":     time.March,
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

func (DPT *DateParseTool) ParseDateFromString(date string) (string, error) {
	stringTime, err := time.Parse("2006-01-02", date)

	if err == nil {
		return stringTime.UTC().String(), nil
	}

	stringTime, err = time.Parse("02.01.2006", date)
	if err == nil {
		return stringTime.UTC().String(), nil
	}

	// fmt.Println(len(strings.Split(date, " ")))
	switch len(strings.Split(date, " ")) {
	case 1:
		year, err := DPT.CleanYear(date)
		if err != nil {

			return "", err
		}

		return time.Date(year, 0, 0, 0, 0, 0, 0, time.UTC).UTC().String(), nil
	case 2:
		stringTime, err = DPT.ParseMonthYearDate(date)

		if err != nil {
			return "", err
		}

		return stringTime.UTC().String(), nil
	case 3:
		stringTime, err = DPT.ParseDayMonthYearDate(date)

		if err != nil {
			return "", err
		}

		return stringTime.UTC().String(), nil
	
	case 4:
		
		fmt.Println(string(date[len(date)-1]))
		if string(date[len(date)-1]) == "³" {
			date = strings.TrimSuffix(date, "³")
			stringTime, err = DPT.ParseDayMonthYearDate(date)

			if err != nil {
				return "", err
			}

			return stringTime.UTC().String(), nil
		}

		if string(date[len(date)-1]) == "." {
			date = strings.TrimSuffix(date, "³.")
			stringTime, err = DPT.ParseDayMonthYearDate(date)

			if err != nil {
				return "", err
			}

			return stringTime.UTC().String(), nil
		}
		
	}

	return "нет данных", nil
}

func (DPT *DateParseTool) ParseDayMonthYearDate(date string) (time.Time, error) {

	dates := strings.Split(date, " ")

	day, err := strconv.Atoi(dates[0])

	if err != nil {
		return time.Time{}, err
	}

	var month time.Month
	if m, exists := months1[dates[1]]; exists {
		month = m
	} else {
		month = months2[dates[1]]
	}

	year, err := strconv.Atoi(dates[2])

	if err != nil {
		year, err = DPT.CleanYear(dates[2])

		if err != nil {
			return time.Time{}, err
		}
	}

	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC), nil
}

func (DPT *DateParseTool) ParseMonthYearDate(date string) (time.Time, error) {
	dates := strings.Split(date, " ")

	var month time.Month
	if m, exists := months1[dates[0]]; exists {
		month = m
	} else {
		month = months2[dates[0]]
	}

	year, err := strconv.Atoi(dates[1])

	if err != nil {
		year, err = DPT.CleanYear(dates[1])

		if err != nil {
			return time.Time{}, err
		}
	}

	return time.Date(year, month, 0, 0, 0, 0, 0, time.UTC), nil
}

func (DPT *DateParseTool) ParseYearDate(date string) (time.Time, error) {
	dates := strings.Split(date, " ")

	year, err := strconv.Atoi(dates[0])

	if err != nil {
		year, err = DPT.CleanYear(dates[0])

		if err != nil {
			return time.Time{}, err
		}
	}

	return time.Date(year, 0, 0, 0, 0, 0, 0, time.UTC), nil
}

func (DPT *DateParseTool) CleanYear(date string) (int, error) {
	var buf strings.Builder

	for _, ch := range date {
		if unicode.IsDigit(ch) {
			buf.WriteString(string(ch))
		}
	}

	year, err := strconv.Atoi(buf.String())

	if err != nil {
		return 0, err
	}

	return year, nil
}