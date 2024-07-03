package memorywall

import (
	"encoding/base64"
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

		var humanInfo HumanInfo = HumanInfo{
			Name:                       utils.GetFileNameWithOutExt(file.Filename),
			Description:                docReader.GetFullDescription("<br>"),
			PlaceOfBirth:               docReader.GetPlaceOfBirth(),
			DateAndPlaceOfСonscription: docReader.GetPlaceAndDateOfСonscription(),
			MilitaryRank:               docReader.GetMilitaryRank(),
			Awards:                     docReader.GetMedals(),
			Images:                     preparedImages,
		}

		var resp ParseDocxResponse = ParseDocxResponse{
			Filename:  file.Filename,
			HumanInfo: humanInfo,
		}
		MS.Response = append(MS.Response, resp)
	}

	return MS.Response, nil
}

func (MS *MemoryWallService) PrepareImagesToSend(images map[string][]byte) (map[string]string, error) {
	convertedImages := make(map[string]string)
	for imgName, imgData := range images {
		convertedImages[imgName] = base64.StdEncoding.EncodeToString(imgData)
	}

	return convertedImages, nil
}
