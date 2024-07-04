package memorywall

import (
	"encoding/base64"
	"memory_wall/internal/http/memory_wall/models"
	"memory_wall/internal/readers"
	"memory_wall/lib/utils"
	"mime/multipart"
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
		defer openedFile.Close()

		humanReader, err := readers.ReadFromDocx(openedFile, file.Size)
		if err != nil {
			return []models.ParseDocxResponse{}, err
		}

		images, err := humanReader.GetImages()
		if err != nil {
			return []models.ParseDocxResponse{}, err
		}

		// preparedImages, err := MS.PrepareImagesToSend(images)
		if err != nil {
			return []models.ParseDocxResponse{}, err
		}

		FIO := humanReader.GetFIO()
		var humanInfo models.HumanInfo = models.HumanInfo{
			Description:                humanReader.GetFullDescription("<br>"),
			PlaceOfBirth:               humanReader.GetPlaceOfBirth(),
			DateAndPlaceOfСonscription: humanReader.GetPlaceAndDateOfСonscription(),
			MilitaryRank:               humanReader.GetMilitaryRank(),
			Awards:                     humanReader.GetMedals(),
			Images:                     images,
		}

		if len(FIO) == 3 {
			humanInfo.FirstName = FIO[0]
			humanInfo.LastName = FIO[1]
			humanInfo.MiddleName = FIO[2]
		} else {
			humanInfo.Name = utils.GetFileNameWithOutExt(file.Filename)
		}

		birthDates := humanReader.GetBirthDate()
		if len(birthDates) == 2 {
			humanInfo.Birthday = birthDates[0]
			humanInfo.Deathday = birthDates[1]
		}

		var resp models.ParseDocxResponse = models.ParseDocxResponse{
			Filename:  file.Filename,
			HumanInfo: humanInfo,
		}
		response = append(response, resp)
	}

	return response, nil
}

func (MS *MemoryWallService) PrepareImagesToSend(images map[string][]byte) (map[string]string, error) {
	convertedImages := make(map[string]string)
	for imgName, imgData := range images {
		convertedImages[imgName] = base64.StdEncoding.EncodeToString(imgData)
	}

	return convertedImages, nil
}
