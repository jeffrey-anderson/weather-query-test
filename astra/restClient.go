package astra

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"web-client/model"
)

func GetForecastForLocation(location string, dataDate string) ([]model.DbRecord, error) {

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	astraUri := fmt.Sprintf("https://%s-%s.apps.astra.datastax.com/api/rest/v2/keyspaces/weather/api_data?where={\"location_id\":{\"$eq\":\"%s\"},\"data_date\":{\"$eq\":\"%s\"}}",
		os.Getenv("ASTRA_DB_ID"),
		os.Getenv("ASTRA_DB_REGION"),
		location, dataDate)
	bearerToken := os.Getenv("ASTRA_DB_APPLICATION_TOKEN")

	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodGet, astraUri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-cassandra-token", bearerToken)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: got %v", res.Status))
	}

	record := new(model.ResponseRecord)

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyBytes, &record)
	if err != nil {
		return nil, err
	}

	var returnList []model.DbRecord

	var returnRec model.DbRecord
	if record.Count > 0 {
		returnList = make([]model.DbRecord, record.Count)
		for i := 0; i < record.Count; i++ {
			returnRec = model.DbRecord{
				LocationId: record.Data[i].LocationId,
				DataDate:   record.Data[i].DataDate,
				DataHour:   record.Data[i].DataHour,
				Parameters: model.Parameters{},
			}
			for j := 0; j < len(record.Data[i].Parameters); j++ {
				switch record.Data[i].Parameters[j].Key {
				case "airTemp":
					returnRec.Parameters.AirTemp = record.Data[i].Parameters[j].Value
				case "cloudCover":
					returnRec.Parameters.CloudCover = record.Data[i].Parameters[j].Value
				case "conditionalProbFreezing":
					returnRec.Parameters.ConditionalProbFreezing = record.Data[i].Parameters[j].Value
				case "dewPoint":
					returnRec.Parameters.DewPoint = record.Data[i].Parameters[j].Value
				case "precipProb":
					returnRec.Parameters.PrecipProb = record.Data[i].Parameters[j].Value
				}
			}
			returnList[i] = returnRec
		}
	}
	return returnList, nil
}
