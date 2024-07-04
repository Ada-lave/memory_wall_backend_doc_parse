package memorywall

import (
	"encoding/base64"
	"memory_wall/lib/utils"
	"mime/multipart"
)

type MemoryWallService struct {
}

func newMemoryWallService() *MemoryWallService {
	return &MemoryWallService{}
}

func (MS *MemoryWallService) ParseDocx(files []multipart.FileHeader) ([]ParseDocxResponse, error) {
	var response []ParseDocxResponse
	for _, file := range files {
		openedFile, err := file.Open()
		if err != nil {
			return []ParseDocxResponse{}, err
		}
		defer openedFile.Close()

		var docReader utils.DocxReader
		docReader, err = docReader.NewDocxReader(openedFile, file.Size)
		if err != nil {
			return []ParseDocxResponse{}, err
		}

		images, err := docReader.GetImages()
		if err != nil {
			return []ParseDocxResponse{}, err
		}

		preparedImages, err := MS.PrepareImagesToSend(images)
		if err != nil {
			return []ParseDocxResponse{}, err
		}

		FIO := docReader.GetFIO()

		var humanInfo HumanInfo = HumanInfo{
			Description:                docReader.GetFullDescription("<br>"),
			PlaceOfBirth:               docReader.GetPlaceOfBirth(),
			DateAndPlaceOfСonscription: docReader.GetPlaceAndDateOfСonscription(),
			MilitaryRank:               docReader.GetMilitaryRank(),
			Awards:                     docReader.GetMedals(),
			Images:                     preparedImages,
		}

		if len(FIO) == 3 {
			humanInfo.FirstName = FIO[0]
			humanInfo.LastName = FIO[1]
			humanInfo.MiddleName = FIO[2]
		} else {
			humanInfo.Name = utils.GetFileNameWithOutExt(file.Filename)
		}

		birthDates := docReader.GetBirthDate()
		if len(birthDates) == 2 {
			humanInfo.Birthday = birthDates[0]
			humanInfo.Deathday = birthDates[1]
		}

		var resp ParseDocxResponse = ParseDocxResponse{
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
