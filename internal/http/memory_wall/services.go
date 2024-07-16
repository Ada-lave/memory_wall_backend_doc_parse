package memorywall

import (
	"io"
	"memory_wall/internal/http/memory_wall/models"
	"memory_wall/internal/readers"
	"memory_wall/lib/utils"
	"mime/multipart"
	"os"
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
			FIO :=  strings.Split(utils.GetFileNameWithOutExt(file.Filename), " ") 
			
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

			fileResponse.HumanInfo = humanInfo
			response = append(response, fileResponse)
			continue
		}

		FIO :=  strings.Split(utils.GetFileNameWithOutExt(file.Filename), " ") 

		dateAndPlaceOfСonscription, err := MS.ExtractPlaceAndDateOfСonscription(humanReader.FullText)
		if err != nil {
			fileResponse.Errors = append(fileResponse.Errors, err.Error())
		
		}

		images, err := humanReader.GetImages()
		if err != nil {
			fileResponse.Errors = append(fileResponse.Errors, err.Error())
		
		}

		var humanInfo models.HumanInfo = models.HumanInfo{
			Description:                humanReader.GetFullDescription("<br>"),
			PlaceOfBirth:               humanReader.GetPlaceOfBirth(),
			DateAndPlaceOfСonscription: dateAndPlaceOfСonscription,
			MilitaryRank:               humanReader.GetMilitaryRank(),
			Awards:                     humanReader.GetMedals(),
			Images:                     images,
		}

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
		birthDates, err := MS.ExtractBirthAndDeathDate(humanReader.FullText)
		if err != nil {
			response = append(response, models.ParseDocxResponse{
				Filename: "err",
			})
			continue
		}
		if len(birthDates) == 2 {
			humanInfo.Birthday = birthDates[0]
			humanInfo.Deathday = birthDates[1]
		}
		
		fileResponse.HumanInfo = humanInfo
		
		response = append(response, fileResponse)
	}
	return response, nil
}

func (MS *MemoryWallService) ExtractFIO(text string) ([]string, error) {
	humanFIOReader, err := readers.NewHumanFIOReader(text)
	if err != nil {
		return []string{}, err
	}
	fio := humanFIOReader.GetFIO()

	return fio, nil
}

func (MS *MemoryWallService) ExtractBirthAndDeathDate(text string) ([]string, error) {
	humanDateReader, err := readers.NewHumanDateReader(text)
	if err != nil {
		return []string{}, err
	}
	placeAndDateOfСonscription := humanDateReader.GetBirthAndDeathDate()

	return placeAndDateOfСonscription, nil
}

func (MS *MemoryWallService) ExtractPlaceAndDateOfСonscription(text string) (string, error) {
	humanDateReader, err := readers.NewHumanDateReader(text)
	if err != nil {
		return "", err
	}
	placeAndDate := humanDateReader.GetPlaceAndDateOfСonscription()

	return placeAndDate, nil
}

func (MS *MemoryWallService) SaveBadFile(file multipart.File, filename string) {
	newFile, err := os.Create("storage/bad_files/" + filename)
	if err != nil {
		panic(err)

	}

	defer newFile.Close()

	_, err = io.Copy(newFile, file)
	if err != nil {
		panic(err)
	}
}

// TODO: Вынести это функционал
// func (MS *MemoryWallService) PrepareImagesToSend(images map[string][]byte) (map[string]string, error) {
// 	convertedImages := make(map[string]string)
// 	for imgName, imgData := range images {
// 		convertedImages[imgName] = base64.StdEncoding.EncodeToString(imgData)
// 	}

// 	return convertedImages, nil
// }
