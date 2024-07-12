package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestExtractDataFromText(t *testing.T) {
	type args struct{
		text string
		sub string
		sep string
	}

	testCases := []struct {
		Name string
		textFormatter *TextFormatter
		Args args
		Want string
	}{
		{
			Name: "Test case where extract place of birth",
			Args: args{
				text: "<br><br>Место рождения: Костромская область, Сусанинский район, д.Мохнево<br>Место и дата призыва: 1942 год, Кинешменским РВК<br>",
				sub: "Место рождения:",
				sep: "<br>",
			},
			Want: "Костромская область, Сусанинский район, д.Мохнево",
		},
		{
			Name: "Test case where extract place of birth",
			Args: args{
				text: "ГРАЧЕВ<br>СЕРГЕЙ АЛЕКСАНДРОВИЧ<br>(8 сентября 1899 – 15 июля 1966  )<br><br>Место рождения: Костромская область, Сусанинский район, д.Мохнево<br>Место и дата призыва: 1942 год, Кинешменским РВК<br>Воинское звание, должность: красноармеец; стрелок, плотник-мостовик<br>Военная служба: Призван на фронт в январе 1942 года. До 25 ноября 1942 года воевал в составе 212-го стрелкового полка 49-й стрелковой дивизии на Донском (Сталинградском) фронте. Получил ранение, направлен на излечение в госпиталь. После возвращение в строй с мая 1943 года по 7 января 1946 года продолжил военную службу в составе  15-го мостового батальона в должности плотника-мостовика.<br>Ранение, контузия: легкое ранение 1942 год. ЭГ 602<br>Награды:<br>Медаль «За победу над Германией» (09.05.1945)<br>Юбилейные медали<br><br>",
				sub: "Место и дата призыва:",
				sep: "<br>",
			},
			Want: "1942 год, Кинешменским РВК",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			res := tc.textFormatter.ExtractDataFromText(tc.Args.text, tc.Args.sub, tc.Args.sep)

			assert.Equal(t, tc.Want, res)
		})
	}

}
func TestFormatText(t *testing.T) {
	testCases := []struct {
		Name string
		textFormatter *TextFormatter
		Text string
		Want string
	}{
		{
			Name: "Test case where all text is ':'",
			Text: "::::::::::::::::::::::::::::::::::::",
			Want: "",
		},
		{
			Name: "Test case where text is ':' and whitespaces",
			Text: "     ::::::::::::::::::::::::::::     ",
			Want: "",
		},
		{
			Name: "Normal test case",
			Text: "Al:e:x want smth....       ",
			Want: "Alex want smth....",
		},
		{
			Name: "Empty text test case",
			Text: "",
			Want: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			res := tc.textFormatter.FormatText(tc.Text)
			assert.Equal(t, tc.Want, res)
		})
	}
}

func TestCapitalizeWords(t *testing.T) {
	testCases := []struct {
		Name string
		textFormatter *TextFormatter
		Text string
		Want string
	}{
		{
			Name: "Good test case",
			Text: "alex want to eat bulka",
			Want: "Alex Want To Eat Bulka",
		},
		{
			Name: "Test case with full uppercase text",
			Text: "ALEX WANT TO EAT",
			Want: "Alex Want To Eat",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			res := tc.textFormatter.CapitalizeWords(tc.Text)
			assert.Equal(t, tc.Want, res)
		})
	}
}

func TestCheckStringIsDate(t *testing.T) {
	tests := []struct {
		name string
		date string
		want bool
	}{
		{
			name: "Basic test with normal date",
			date: "27.05.2005",
			want: true,
		},
		{
			name: "Test case with month and year",
			date: "июль 2005",
			want: true,
		},
		{
			name: "Test case with only year only",
			date: "1902",
			want: true,
		},
		{
			name: "Test case with typical date",
			date: "17 августа 1967",
			want: true,
		},
		{
			name: "Test case with no date",
			date: "HACKER",
			want: false,
		},
		{
			name: "Test case with empty string",
			date: "",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckStringIsDate(tt.date); got != tt.want {
				t.Errorf("CheckStringIsDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
