package readers

import (
	"memory_wall/lib/utils"
	"strings"
)

type HumanMilitaryReader struct {
	FullText      string
	textFormatter *utils.TextFormatter
}

func (HMR *HumanMilitaryReader) GetPlaceOfBirth() string {
	placeOfBirth := HMR.textFormatter.ExtractDataFromText(HMR.FullText, "Место рождения", "<br>")

	return placeOfBirth
}

func (HMR *HumanMilitaryReader) GetMilitaryRank() string {
	rank := HMR.textFormatter.ExtractDataFromText(HMR.FullText, "Воинское звание, должность", "<br>")

	if len(rank) == 0 {
		rank = HMR.textFormatter.ExtractDataFromText(HMR.FullText, "Воинское звание", "<br>")
	}

	return rank
}

func (HMR *HumanMilitaryReader) GetMedals() []string {
	var awards []string
	if strings.Contains(HMR.FullText, "Награды:") {
		textOfMedal := strings.Split(HMR.FullText, "Награды:")[1]
		for _, medal := range strings.Split(textOfMedal, "<br>") {
			if medal != "" && strings.Contains(strings.ToLower(medal), "медаль") {
				awards = append(awards, medal)
			}
		}
	}
	return awards
}

func NewHumanMilitaryReader(text string) HumanMilitaryReader {
	return HumanMilitaryReader{
		FullText: text,
	}
}
