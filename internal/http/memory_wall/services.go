package memorywall

import (
	"memory_wall/internal/http/memory_wall/models"
	"memory_wall/internal/readers"
	"memory_wall/lib/utils"
	"mime/multipart"
	"strings"
)

type MemoryWallService struct {
}

func newMemoryWallService() *MemoryWallService {
	return &MemoryWallService{}
}

func (MS *MemoryWallService) ParseDocx(files []multipart.FileHeader) ([]models.ParseDocxResponse, error) {
	var response []models.ParseDocxResponse
	for _, file := range files {
		openedFile, err := file.Open()
		if err != nil {
			return []models.ParseDocxResponse{}, err
		}
		var humanInfo models.HumanInfo = models.HumanInfo{}
		fileResponse := models.ParseDocxResponse{
			Filename: utils.GetFileNameWithOutExt(file.Filename),
		}

		if file.Size < 1 {
			fileResponse.Errors = append(fileResponse.Errors, "file size 1")

		}
		humanReader, err := readers.NewHumanInfoReader(openedFile, file.Size)
		if err != nil {
			fileResponse.Errors = append(fileResponse.Errors, err.Error())
			var humanInfo models.HumanInfo

			MS.InjectionFIO(&file, &humanInfo)
			fileResponse.HumanInfo = humanInfo
			response = append(response, fileResponse)
			continue
		}

		description := humanReader.GetFullDescription("<br>")

		humanInfo.Description = description
		images, err := humanReader.GetImages()
		if err != nil {
			fileResponse.Errors = append(fileResponse.Errors, err.Error())

		}
		humanInfo.Images = images

		err = MS.InjectionPlaceAndDateOfСonscription(humanReader.FullText, &humanInfo)
		if err != nil {
			fileResponse.Errors = append(fileResponse.Errors, err.Error())
		}

		err = MS.InjectionBirthAndDeathDate(humanReader.FullText, &humanInfo)
		if err != nil {
			fileResponse.Errors = append(fileResponse.Errors, err.Error())
		}

		MS.InjectionFIO(&file, &humanInfo)
		MS.InjectionPlaceOfBirth(humanReader.FullText, &humanInfo)
		MS.InjectionMedals(humanReader.FullText, &humanInfo)
		MS.IjectionMilitaryRank(humanReader.FullText, &humanInfo)
		fileResponse.HumanInfo = humanInfo

		response = append(response, fileResponse)
	}
	return response, nil
}

func (MS *MemoryWallService) InjectionFIO(file *multipart.FileHeader, humanInfo *models.HumanInfo) {
	FIO := strings.Split(utils.GetFileNameWithOutExt(file.Filename), " ")
	switch len(FIO) {
	case 1:
		humanInfo.FirstName = FIO[0]
	case 2:
		humanInfo.LastName = FIO[0]
		humanInfo.FirstName = FIO[1]
	case 3:
		humanInfo.MiddleName = FIO[2]
		humanInfo.LastName = FIO[0]
		humanInfo.FirstName = FIO[1]
	case 4:
		humanInfo.MiddleName = FIO[2]
		humanInfo.LastName = FIO[0]
		humanInfo.FirstName = FIO[1]
	default:
		humanInfo.Name = utils.GetFileNameWithOutExt(file.Filename)
	}
}

func (MS *MemoryWallService) InjectionBirthAndDeathDate(text string, humanInfo *models.HumanInfo) error {
	humanDateReader, err := readers.NewHumanDateReader(text)
	if err != nil {
		return err
	}
	BirthAndDeadthDay := humanDateReader.GetBirthAndDeathDate()

	if len(BirthAndDeadthDay) == 2 {
		humanInfo.Birthday = BirthAndDeadthDay[0]
		humanInfo.Deathday = BirthAndDeadthDay[1]
	}
	return nil
}

func (MS *MemoryWallService) InjectionPlaceAndDateOfСonscription(text string, humanInfo *models.HumanInfo) error {
	humanDateReader, err := readers.NewHumanDateReader(text)
	if err != nil {
		return err
	}
	placeAndDate := humanDateReader.GetPlaceAndDateOfСonscription()
	humanInfo.DateAndPlaceOfСonscription = placeAndDate

	return nil
}

func (MS *MemoryWallService) InjectionMedals(text string, humanInfo *models.HumanInfo) {
	humanMilitaryReader := readers.NewHumanMilitaryReader(text)
	medals := humanMilitaryReader.GetMedals()
	humanInfo.Awards = medals
}

func (MS *MemoryWallService) InjectionPlaceOfBirth(text string, humanInfo *models.HumanInfo) {
	humanMilitaryReader := readers.NewHumanMilitaryReader(text)
	placeOfBirth := humanMilitaryReader.GetPlaceOfBirth()
	humanInfo.PlaceOfBirth = placeOfBirth
}

func (MS *MemoryWallService) IjectionMilitaryRank(text string, humanInfo *models.HumanInfo) {
	humanMilitaryReader := readers.NewHumanMilitaryReader(text)
	militaryRank := humanMilitaryReader.GetMilitaryRank()
	humanInfo.MilitaryRank = militaryRank
}
