package main

import (
  "fmt"
  "log"
  "net/http"
  "strings"

  "github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
  // Request the HTML page.
  res, err := http.Get("http://scores.nbcsports.com/nba/scoreboard.asp?day=20190206")
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
  }

  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    log.Fatal(err)
  }

  doc.Find("table").Find("a[href]").Each(func(i int, s *goquery.Selection) {
    href, _ := s.Attr("href")
    if strings.Contains(href, "boxscore") {
      fmt.Printf("%d: %s\n", i, href)
    }
	})

  // Find the review items
  // doc.Find("events").Each(func(i int, s *goquery.Selection) {
  //   // For each item found, get the band and title
  //   fmt.Printf("%d", i )
  // })
}

func main() {
  ExampleScrape()
}
