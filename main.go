package main

func main() {
  points := FetchPoints()
  features := ConvertToFeatures(points)
  feature_collections := ConvertToFeatureCollections(features)

  WriteFiles(feature_collections)
}
