package model

type Parameters struct {
	AirTemp                 float32 `json:"airTemp,omitempty"`
	CloudCover              float32 `json:"cloudCover,omitempty"`
	ConditionalProbFreezing float32 `json:"conditionalProbFreezing,omitempty"`
	DewPoint                float32 `json:"dewPoint,omitempty"`
	PrecipProb              float32 `json:"precipProb,omitempty"`
}
