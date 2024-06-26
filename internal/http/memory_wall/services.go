package memorywall

import (
	"io"
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
	_, err = io.ReadAll(openedFile)
	if err != nil {
		return ParseDocxResponse{}, err
	}
	name := utils.GetFileNameWithOutExt(file.Filename)
	var humanInfo HumanInfo = HumanInfo{
		Name: name,
		Description: "test",
		Image: "test",
	}

	var resp ParseDocxResponse = ParseDocxResponse{
		Filename: file.Filename,
		HumanInfo: humanInfo,
	}
	return resp, nil
}

func (MS *MemoryWallService) getAllDocxFileInfoFromStorage(path string) ([]string, error) {
	names, err := utils.WalkInDirAndFindAllFileNames(path)
	if err != nil {
		return []string{}, err
	}
	return names, nil
}
