package memorywall

import (
	"memory_wall/lib/utils"
	"mime/multipart"

)

type MemoryWallService struct {

}

func (MS *MemoryWallService) parseDocx(file multipart.FileHeader) (ParseDocxResponse, error) {

	openedFile, err := file.Open()
	if err != nil {
		return ParseDocxResponse{}, err
	}

	name := utils.GetFileNameWithOutExt(file.Filename)
	description := utils.GetTextFromFile(openedFile, file.Size)
	var humanInfo HumanInfo = HumanInfo{
		Name: name,
		Description: description,
		Image: "test",
	}

	var resp ParseDocxResponse = ParseDocxResponse{
		Filename: file.Filename,
		HumanInfo: humanInfo,
	}
	return resp, nil
}

// func (MS *MemoryWallService) getAllDocxFileInfoFromStorage(path string) ([]string, error) {
// 	names, err := utils.WalkInDirAndFindAllFileNames(path)
// 	if err != nil {
// 		return []string{}, err
// 	}
// 	return names, nil
// }
