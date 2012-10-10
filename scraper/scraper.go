package scraper

import (
  "github.com/mitchellwrosen/ebay-scraper-go/urlInfo"

  "fmt"
  "time"
)

type info struct {
  latest int64
  averageSale float32
  urlInfo urlInfo.UrlInfo
}

type Scraper struct {
  sem chan int
  auctionCh chan string
  name string
  binInfo info
  endingInfo info
  endedInfo info
}

func New(sem chan int, auctionCh chan string, name string,
         binUrlInfo urlInfo.UrlInfo, endingUrlInfo urlInfo.UrlInfo,
         endedUrlInfo urlInfo.UrlInfo) *Scraper {
  return &Scraper{sem,
                  auctionCh,
                  name,
                  info{-1, -1, binUrlInfo},
                  info{-1, -1, endingUrlInfo},
                  info{-1, -1, endedUrlInfo}}
}

func (scraper *Scraper) Scrape() {
  // Ended goroutine.
  go func() {
    for {
      scraper.sem <- 1
      fmt.Printf("%s ended scraper\n", scraper.name)
      <-scraper.sem
      time.Sleep(10 * time.Second)
    }
  }()

  // Bin goroutine.
  go func() {
    for {
      scraper.sem <- 1
      scraper.auctionCh <- fmt.Sprintf("%s Bin sold for $500", scraper.name)
      <-scraper.sem
      time.Sleep(1 * time.Second)
    }
  }()

  // Ending goroutine.
  go func() {
    for {
      scraper.sem <- 1
      scraper.auctionCh <- fmt.Sprintf("%s ending in 5 mins", scraper.name)
      <-scraper.sem
      time.Sleep(5 * time.Second)
    }
  }()
}
