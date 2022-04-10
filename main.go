package main

import (
	"log"
	"web-client/astra"
)

func main() {
	records, err := astra.GetForecastForLocation("834c1dfffffffff", "20020506")
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(records); i++ {
		log.Printf("Record: %#v\n", records[i])
	}

}
