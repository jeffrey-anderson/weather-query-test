package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"web-client/model"
)

func main() {

	log.Println("Hello!")
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	astraUri := fmt.Sprintf("https://%s-%s.apps.astra.datastax.com/api/rest/v2/keyspaces/weather/api_data?where={\"location_id\":{\"$eq\":\"834c1dfffffffff\"},\"data_date\":{\"$eq\":\"20020506\"}},raw=true",
		os.Getenv("ASTRA_DB_ID"),
		os.Getenv("ASTRA_DB_REGION"))
	bearerToken := os.Getenv("ASTRA_DB_APPLICATION_TOKEN")

	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodGet, astraUri, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("x-cassandra-token", bearerToken)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: got %v", res.Status))
	}
	fmt.Println(res.Header.Get("Content-Type"))

	record := new(model.ResponseRecord)

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	log.Printf("Body: %s\n", string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &record)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", record)

}
