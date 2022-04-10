package model

type RespParams struct {
	Key   string  `json:"key"`
	Value float32 `json:"value"`
}

type RespData struct {
	LocationId string       `json:"location_id"`
	DataDate   string       `json:"data_date"`
	DataHour   int          `json:"data_hour"`
	Parameters []RespParams `json:"parameters"`
}

type ResponseRecord struct {
	Count int        `json:"count"`
	Data  []RespData `json:"data"`
}
