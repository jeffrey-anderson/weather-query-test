package model

type DbRecord struct {
	LocationId string     `json:"location_id"`
	DataDate   string     `json:"data_date"`
	DataHour   int        `json:"data_hour"`
	Parameters Parameters `json:"parameters"`
}
