package memorywall

import (
	"memory_wall/lib/utils"
	"mime/multipart"
	"sync"
)

type MemoryWallService struct {

}

func (MS *MemoryWallService) parseDocx(files []multipart.FileHeader) ([]ParseDocxResponse, error) {
	var response []ParseDocxResponse
	var wg sync.WaitGroup

	wg.Add(len(files))
	for _, file := range files {
		go func() {
			defer wg.Done()
			openedFile, err := file.Open()
			if err != nil {
				panic(err)
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
			response = append(response, resp)
		}()
		
	}
	wg.Wait()
	return response, nil
}

func parseFile(file multipart.FileHeader, response *[]ParseDocxResponse) error {
	openedFile, err := file.Open()
		if err != nil {
			return err
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
		*response = append(*response, resp)

		return nil
}

// func (MS *MemoryWallService) getAllDocxFileInfoFromStorage(path string) ([]string, error) {
// 	names, err := utils.WalkInDirAndFindAllFileNames(path)
// 	if err != nil {
// 		return []string{}, err
// 	}
// 	return names, nil
// }
