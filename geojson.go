package main
//
// import (
//   "fmt"
//   )
//
// type Properties struct { 
//   Name        string `json:"name"`
//   Category    string `json:"category"`
//   Description string `json:"description"`
// }
//
// type Geometry struct {
//   Type        string `json:"type"`
//   Coordinates []float64 `json:"coordinates"`
// }
//
// type FeatureCollection struct {
//   Type     string `json:"type"`
//   Features []Feature `json:"features"`
// }
//
// func NewFeatureCollection(features []Feature) FeatureCollection {
//   feature_collection := FeatureCollection{Type: "FeatureCollection", Features: features}
//
//   return feature_collection
// }
//
// type Feature struct {
//   Type        string `json:"type"`
//   Properties  Properties `json:"properties"`
//   Geometry    Geometry `json:"geometry"`
// }
//
// func NewFeature(point OverlanderPoint) Feature {
// // func NewFeature(name string, description string, category string, latitude float64, longitude float64) Feature {
//   feature := Feature{
//     Type: "Feature",
//     Properties: Properties{
//       Name: point.Name,
//       Description: build_description(point),
//       Category: point.Category,
//     },
//     Geometry: Geometry{
//       Type: "Point",
//       Coordinates: []float64{point.Location.Longitude, point.Location.Latitude},
//     },
//   }
//
//   return feature
// }
//
// func ConvertToFeatures(points []OverlanderPoint) []Feature {
//   var features []Feature
//
//   for _, point := range points {
//     feature := NewFeature(point)
//     features = append(features, feature)
//   }
//
//   return features
// }
//
// func ConvertToFeatureCollections(features []Feature) []FeatureCollection {
//   var collections []FeatureCollection
//   collection_size := 999
//
//   for i := 0; i < len(features); i += collection_size {
//     end := i + collection_size
//
//     if end > len(features) {
//       end = len(features)
//     }
//
//     collections = append(collections, NewFeatureCollection(features[i:end]))
//   }
//
//   return collections
// }
//
// func build_description(point OverlanderPoint) string {
//   description := fmt.Sprintf("%s\n\n%s\n\n%s", point.Category, point.Description, point.Comments)
//  
//   return description
// }
//
