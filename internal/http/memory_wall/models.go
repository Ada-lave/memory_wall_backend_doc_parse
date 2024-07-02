package memorywall

import "mime/multipart"

type ParseDocxRequest struct {
	Files *[]multipart.FileHeader `form:"files" binding:"required"`
}

type ParseDocxResponse struct {
	Filename  string `json:"filename"`
	HumanInfo `json:"human_info"`
}

type HumanInfo struct {
	Name                       string   `json:"name"`
	Description                string   `json:"description"`
	PlaceOfBirth               string   `json:"place_of_birth"`
	DateAndPlaceOfСonscription string   `json:"date_and_place_of_conscription"`
	MilitaryRank               string   `json:"military_rank_and_position"`
	Awards                     []string `json:"awards"`
	Image                      string   `json:"image"`
}
