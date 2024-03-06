package dto

type PackSizesDTO struct {
	Sizes []int `json:"sizes"`
}

type PackQuantitiesDTO struct {
	Size     int `json:"size"`
	Quantity int `json:"quantity"`
}
