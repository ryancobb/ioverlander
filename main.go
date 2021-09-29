package main

import (
  "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
  "os"
)

type OverlanderPoint struct {
  Name        string `json:"name"`
  Description string `json:"description"`
  Category    string `json:"category"`
  Location    struct { 
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
  } `json:"location"`
}
type Properties struct { Name        string `json:"name"`
  Category    string `json:"category"`
  Description string `json:"description"`
}

type Geometry struct {
  Type        string `json:"type"`
  Coordinates []float64 `json:"coordinates"`
}

type Feature struct {
  Type        string `json:"type"`
  Properties  Properties `json:"properties"`
  Geometry    Geometry `json:"geometry"`
}

func NewFeature(name string, description string, category string, latitude float64, longitude float64) Feature {
  feature := Feature{
    Type: "Feature",
    Properties: Properties{
      Name: name,
      Description: description,
      Category: category,
    },
    Geometry: Geometry{
      Type: "Point",
      Coordinates: []float64{longitude, latitude},
    },
  }

  return feature
}

type GeoJson struct {
  Type     string `json:"type"`
  Features []Feature `json:"features"`
}

func write_file(features []Feature, i int) {
  geojson := GeoJson{Type: "FeatureCollection", Features: features}
  json_data, _ := json.MarshalIndent(geojson, "", " ")
  file_name := fmt.Sprintf("export-%d.geojson", i)

  err := ioutil.WriteFile(file_name, json_data, 0644)
  
  if err != nil {
    fmt.Println(fmt.Sprintf("Error writing file: %s", file_name))
  }

  fmt.Println(fmt.Sprintf("Wrote file: %s", file_name))
}

func main() {
  resp, err := http.Get("https://ioverlander.com/places/search.json?searchboxmin=44,-125&searchboxmax=45,-124")
  if err != nil {
    fmt.Println("No response from request")
  }
  defer resp.Body.Close()

  jsonFile, err := os.Open("ioverlander-source.json")
  if err != nil {
    fmt.Println(err)
  }

//  byteValue, _ := ioutil.ReadAll(resp.Body)
  byteValue, _ := ioutil.ReadAll(jsonFile)

  var parsed_overlander_points []OverlanderPoint
  if err := json.Unmarshal(byteValue, &parsed_overlander_points); err != nil {
    fmt.Println("Can not unmarshal JSON")
  }

  var chunks [][]OverlanderPoint
  chunkSize := 999
  for i := 0; i < len(parsed_overlander_points); i += chunkSize {
    end := i + chunkSize

    if end > len(parsed_overlander_points) {
      end = len(parsed_overlander_points)
    }

    chunks = append(chunks, parsed_overlander_points[i:end])
  }

  for i, chunk := range chunks {
    var features []Feature

    for _, point := range chunk {
      feature := NewFeature(point.Name, point.Description, point.Category, point.Location.Latitude, point.Location.Longitude)
      features = append(features, feature)
    }

    write_file(features, i)
  }
}
