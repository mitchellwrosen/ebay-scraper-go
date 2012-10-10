package main

import (
  "github.com/mitchellwrosen/ebay-scraper-go/scraper"
  "github.com/mitchellwrosen/ebay-scraper-go/urlInfo"

  "fmt"
)

type phone struct {
  model string
  brand string
  condition string
  carrier string
  storageCapacity string
  color string
}
var phones = [...]int{{'model','brand','cond','carrier','cap','color'}}

func main() {
  sem := make(chan int, 1)

  // Send to this channel
  auctionCh := make(chan string)

  scraper1 := scraper.New(sem,
                          auctionCh,
                          "scraper1",
                          urlInfo.UrlInfo{},
                          urlInfo.UrlInfo{},
                          urlInfo.UrlInfo{})
  scraper2 := scraper.New(sem,
                          auctionCh,
                          "scraper2",
                          urlInfo.UrlInfo{},
                          urlInfo.UrlInfo{},
                          urlInfo.UrlInfo{})
  go scraper1.Scrape()
  go scraper2.Scrape()
  for {
    fmt.Println(<-auctionCh)
  }
}
