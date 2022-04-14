package main

import (
  "encoding/json"
  "fmt"
  "io"
  "net/http"
  "net/url"
  "strconv"

  "github.com/PuerkitoBio/goquery"
  )

type OverlanderPoint struct {
  Id          int    `json:"id"`
  Name        string `json:"name"`
  Description string `json:"description"`
  Category    string `json:"category"`
  Location    struct { 
    Latitude    float64 `json:"latitude"`
    Longitude   float64 `json:"longitude"`
  } `json:"location"`
}

func FetchPoints() []OverlanderPoint {
  fmt.Println("Fetching points...")

  resp, err := http.Get(points_url()) 
  if err != nil {
    fmt.Println("Error fetching points")
  }
  defer resp.Body.Close()

  byteValue, _ := io.ReadAll(resp.Body)

  var parsed_points []OverlanderPoint
  if err := json.Unmarshal(byteValue, &parsed_points); err != nil {
    fmt.Println("Error unmarshalling JSON")
  }

  fmt.Printf("points: %d\n", len(parsed_points))

  return parsed_points
}

func fetch_page(id int) {
  u := url.URL{
    Scheme: "https",
    Host: "ioverlander.com",
    Path: "places/" + strconv.Itoa(id),
  }

  res, err := http.Get(u.String())
  if err != nil {
    fmt.Println("No response from request")
  }
  defer res.Body.Close()

  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    fmt.Println(err)
  }

  doc.Find(".placeCheckin").Each(func(i int, s *goquery.Selection) {
    description := s.Find("p").Text()
    fmt.Println(description)
  })
}

func points_url() string {
  u := url.URL{
    Scheme: "https",
    Host: "ioverlander.com",
    Path: "places/search.json",
  }

  v := url.Values{}
  v.Add("filter[]", "informal_campsite")
  v.Add("filter[]", "wild_campsite")
  // v.Add("searchboxmin", "40,-130")
  // v.Add("searchboxmax", "55,-110")
  v.Add("searchboxmin", "40,-130")
  v.Add("searchboxmax", "50,-110")

  u.RawQuery = v.Encode()

  return u.String()
}
