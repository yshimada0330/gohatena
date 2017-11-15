package hatena

import (
  "strconv"
  "github.com/mmcdole/gofeed"
)

type RssFeed struct {
  Items []*FeedItem
}

type FeedItem struct {
  Title string
  Link string
  BookmarkCount int
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
