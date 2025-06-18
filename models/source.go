package models

type Source struct {
	Category string `json:"category"`
	Language string `json:"language"`
	Country  string `json:"country"`
}