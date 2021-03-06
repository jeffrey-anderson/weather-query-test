package main

import (
	"flag"
	"fmt"
	"github.com/uber/h3-go"
	"log"
	"math/rand"
	"time"
	"web-client/astra"
)

var defaultResolution = flag.Int("resolution", 1, "the H3 resolution to search")
var sampleSize = flag.Int("sampleSize", 100, "the number of samples to test")

func main() {
	flag.Parse()
	log.Println("Starting...")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var ttlTime time.Duration

	noMatchCount := 0
	matchCount := 0

	for i := 0; i < *sampleSize; i++ {
		lat := (r.Float64() * 180) - 90
		long := (r.Float64() * 360) - 180
		point := h3.FromGeo(h3.GeoCoord{Latitude: lat, Longitude: long}, *defaultResolution)
		// fmt.Printf("%f, %f index: %#v is valid: %v\n, ", lat, long, point, h3.IsValid(point))
		start := time.Now()
		loc, _ := astra.GetForecastForLocation(h3.ToString(point), "20020506")
		end := time.Now()
		ttlTime = ttlTime + end.Sub(start)
		if loc == nil {
			fmt.Printf("No match found for %#v\n", h3.ToString(point))
			noMatchCount = noMatchCount + 1
		} else {
			// fmt.Printf("Match found for %#v\n", h3.ToString(point))
			matchCount = matchCount + 1
		}
	}

	log.Printf("Hits: %d, Misses %d in %s. Average time %s\n", matchCount, noMatchCount, ttlTime, ttlTime/time.Duration(*sampleSize))

}
