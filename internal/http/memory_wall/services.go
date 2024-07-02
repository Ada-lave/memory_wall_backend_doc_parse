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

		var humanInfo HumanInfo = HumanInfo{
			Name: utils.GetFileNameWithOutExt(file.Filename),
			Description: docReader.GetFullDescription("<br>"),
			PlaceOfBirth: docReader.GetPlaceOfBirth(),
			DateAndPlaceOfСonscription: docReader.GetPlaceAndDateOfСonscription(),
			MilitaryRank: docReader.GetMilitaryRank(),
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
