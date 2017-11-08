package hatena

import (
  "net/url"
  "strconv"
  "github.com/mmcdole/gofeed"
)

var SEARCH_TEXT_URL string = "http://b.hatena.ne.jp/search/text"

type SearchParameter struct {
  Query string
  Sort string
  Threshold int
  DateBegin string
  DateEnd string
  Safe bool
}

type RssFeed struct {
  Items []*FeedItem
}

type FeedItem struct {
  Title string
  Link string
  BookmarkCount int
}

func SearchTextRss(param *SearchParameter) (*RssFeed){
  values := url.Values{}
  values.Set("mode", "rss")
  values.Set("q", param.Query)

  if param.Sort != "" {
    values.Set("sort", param.Sort)
  }
  
  if param.Threshold > 0 {
    values.Set("threshold", strconv.Itoa(param.Threshold))
  }

  if param.DateBegin != "" {
    values.Set("date_begin", param.DateBegin)
  }

  if param.DateEnd != "" {
    values.Set("date_end", param.DateEnd)
  }

  if param.Safe ==  false {
    values.Set("safe", "off")
  }

  url := SEARCH_TEXT_URL + "?" + values.Encode()
  return createFeed(url)
}

func createFeed(url string) (*RssFeed){
  fp := gofeed.NewParser()
  feed, _ := fp.ParseURL(url)
  items := feed.Items
  feedItems := []*FeedItem{}
  for _, item := range items {
    feedItem := parseItem(item)
    feedItems = append(feedItems, feedItem)
  }
  return &RssFeed{Items: feedItems}
}

func parseItem(item *gofeed.Item) (*FeedItem) {
  extension := item.Extensions
  bookmarkcount := extension["hatena"]["bookmarkcount"]
  count, _ := strconv.Atoi(bookmarkcount[0].Value)
  return &FeedItem{Title: item.Title, Link: item.Link, BookmarkCount: count}
}
