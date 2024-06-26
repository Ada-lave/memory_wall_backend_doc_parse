package memorywall

import "mime/multipart"

type ParseDocxRequest struct {
	Files *[]multipart.FileHeader `form:"files" binding:"required"`
}

type ParseDocxResponse struct {
	Filename string `json:"filename"`
	HumanInfo `json:"human_info"`
}

type HumanInfo struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`

}