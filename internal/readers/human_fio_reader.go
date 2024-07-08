package readers

import (
	"mime/multipart"
	"strings"
)

type HumanFIOReader struct {
	HumanInfoReader
}

func (HFR *HumanFIOReader) GetFIO() []string {
	var FIO []string
	if HFR.FullText == "" {
		HFR.GetFullDescription("<br>")
	}
	if HFR.FullText == "<br>" {
		return []string{}
	}

	var data []string
	for _, text := range strings.Split(HFR.FullText, "<br>") {
		if len(text) != 0 {
			data = append(data, text)
		}
	}
	for _, text := range strings.Split(data[0], " ") {
		if text != "" {
			FIO = append(FIO, text)
		}
	}

	if len(FIO) != 3 && len(data) > 2 {
		splitedNames := strings.Split(data[1], " ")
		FIO = append(FIO, splitedNames[0])
		FIO = append(FIO, splitedNames[1])
	}
	return FIO
}

func NewHumanFIOReader(file multipart.File, size int64) (HumanFIOReader, error) {
	var humanInfoReader HumanInfoReader

	humanInfoReader, err := NewHumanInfoReader(file, size)
	if err != nil {
		return HumanFIOReader{}, err
	}

	return HumanFIOReader{
		HumanInfoReader: humanInfoReader,
	}, nil
}
