package main

import (
  "encoding/json"
  "fmt"
  "io"
  "net/http"
  "net/url"
  // "strconv"
  // "strings"

  // "github.com/PuerkitoBio/goquery"
  )

type RawOverlanderPoint struct {
  Id          int    `json:"id"`
  Name        string `json:"name"`
  Description string `json:"description"`
  Category    string `json:"category"`
  Location    struct {
    Latitude    float64 `json:"latitude"`
    Longitude   float64 `json:"longitude"`
  } `json:"location"`
  Comments  string
}

func (point RawOverlanderPoint) toDb() OverlanderPoint {
  return OverlanderPoint{
    Id: point.Id,
    Name: point.Name,
    Description: point.Description,
    Category: point.Category,
    Latitude: fmt.Sprintf("%f", point.Location.Latitude),
    Longitude: fmt.Sprintf("%f", point.Location.Longitude),
  }
}

// func (point *OverlanderPoint) UpdateComments(comments string) {
//   point.Comments = comments
// }

func FetchPoints() []RawOverlanderPoint {
  fmt.Println("Fetching points...")

  resp, err := http.Get(points_url()) 
  if err != nil {
    fmt.Println("Error fetching points")
  }
  defer resp.Body.Close()

  byteValue, _ := io.ReadAll(resp.Body)

  var raw_points []RawOverlanderPoint
  if err := json.Unmarshal(byteValue, &raw_points); err != nil {
    fmt.Println("Error unmarshalling JSON")
  }

  // fmt.Println("Fetching comments...")
  //
  // for i := range parsed_points {
  //   fmt.Printf("[%d/%d]\n", i + 1, len(parsed_points))
  //   
  //   comments := fetch_comments(parsed_points[i].Id)
  //   // parsed_points[i].UpdateComments(comments)
  // }
  //
  return raw_points
}

// func fetch_comments(id int) string {
//   u := url.URL{
//     Scheme: "https",
//     Host: "ioverlander.com",
//     Path: "places/" + strconv.Itoa(id),
//   }
//
//   res, err := http.Get(u.String())
//   if err != nil {
//     fmt.Println("No response from request")
//   }
//   defer res.Body.Close()
//
//   // Load the HTML document
//   doc, err := goquery.NewDocumentFromReader(res.Body)
//   if err != nil {
//     fmt.Println(err)
//   }
//
//   comments := []string{}
//   doc.Find(".small.mb-3").Each(func(i int, s *goquery.Selection) {
//     date := s.Find("a").First().Text()
//     description := s.Find("p").Text()
//
//     comment := fmt.Sprintf("%s\n%s\n", date, description)
//     comments = append(comments, comment)
//   })
//
//   return strings.Join(comments, "\n")
// }

func points_url() string {
  u := url.URL{
    Scheme: "https",
    Host: "ioverlander.com",
    Path: "places/search.json",
  }

  v := url.Values{}
  v.Add("countrycode[]", "CAN")
  v.Add("filter[]", "informal_campsite")
  v.Add("filter[]", "wild_campsite")
  v.Add("last_verified", "94670856")

  u.RawQuery = v.Encode()
  return u.String()
}

