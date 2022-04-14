package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  )

func WriteFiles(feature_collections []FeatureCollection) {
  fmt.Println("\nFiles:")

  for i, feature_collection := range feature_collections {
    write_file(feature_collection, i)
  }
}

func write_file(feature_collection FeatureCollection, i int) {
  json_data, _ := json.MarshalIndent(feature_collection, "", " ")
  file_name := fmt.Sprintf("exports/export-%d.geojson", i)

  fmt.Println(file_name)

  err := ioutil.WriteFile(file_name, json_data, 0644)

  if err != nil {
    fmt.Println(fmt.Sprintf("Error writing file: %s", file_name))
  }
}
