package models

import "mime/multipart"

type ParseDocxRequest struct {
	Files *[]multipart.FileHeader `form:"files" binding:"required"`
}

type ParseDocxResponse struct {
	Filename  string `json:"filename"`
	HumanInfo `json:"human_info"`
}

type HumanInfo struct {
	Name                       string            `json:"name"`
	FirstName                  string            `json:"first_name"`
	LastName                   string            `json:"last_name"`
	MiddleName                 string            `json:"middle_name"`
	Description                string            `json:"description"`
	Birthday                   string            `json:"birthday"`
	Deathday                   string            `json:"deathday"`
	PlaceOfBirth               string            `json:"place_of_birth"`
	DateAndPlaceOfСonscription string            `json:"date_and_place_of_conscription"`
	MilitaryRank               string            `json:"military_rank_and_position"`
	Awards                     []string          `json:"awards"`
	Images                     []HumanInfoImage `json:"images"`
}

type HumanInfoImage struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
}