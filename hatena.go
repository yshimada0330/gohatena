package hatena

import (
  "net/url"
  "strconv"
  "github.com/mmcdole/gofeed"
)

const SearchTextUrl = "http://b.hatena.ne.jp/search/text"

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
  values.Set("q", param.Query)
  values.Set("sort", param.Sort)
  values.Set("threshold", strconv.Itoa(param.Threshold))
  values.Set("date_begin", param.DateBegin)
  values.Set("date_end", param.DateEnd)
//  values.Set("safe", param.Safe)
  values.Set("mode", "rss")
  url := SearchTextUrl + "?" + values.Encode()
  return feed(url)
}

func feed(url string) (*RssFeed){
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
