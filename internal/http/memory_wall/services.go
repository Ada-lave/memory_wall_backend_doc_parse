package memorywall

import (
	"memory_wall/lib/utils"
	"mime/multipart"
)

type MemoryWallService struct {
	Response []ParseDocxResponse
}

func (MS *MemoryWallService) ParseDocx(files []multipart.FileHeader) ([]ParseDocxResponse, error) {
	for _, file := range files {	
		openedFile, err := file.Open()
		if err != nil {
			return []ParseDocxResponse{}, err
		}
		var docReader utils.DocxReader
		docReader, err = docReader.NewDocxReader(openedFile, file.Size)
		if err != nil {
			return []ParseDocxResponse{}, err
		}
		name := utils.GetFileNameWithOutExt(file.Filename)
		description :=	docReader.GetFullDescription("<br>")
		placeOfBirth := docReader.GetPlaceOfBirth()
		dateAndPlaceOfСonscription := docReader.GetPlaceAndDateOfСonscription()
		var humanInfo HumanInfo = HumanInfo{
			Name: name,
			Description: description,
			PlaceOfBirth: placeOfBirth,
			DateAndPlaceOfСonscription: dateAndPlaceOfСonscription,
			Image: "test",
		}

		var resp ParseDocxResponse = ParseDocxResponse{
			Filename: file.Filename,
			HumanInfo: humanInfo,
		}
		MS.Response = append(MS.Response, resp)
	}

	return MS.Response, nil
}
