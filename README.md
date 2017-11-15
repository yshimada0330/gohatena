# gohatena

[![Build Status](https://travis-ci.org/yshimada0330/gohatena.svg?branch=master)](https://travis-ci.org/yshimada0330/gohatena)

Hatena Rss Search

## Installation
`go get github.com/yshimada0330/gohatena`

## Usage

### search text
http://b.hatena.ne.jp/search/text?q=test&mode=rss

```go
package main

import (
    "github.com/yshimada0330/gohatena"
    "fmt"
)

func main() {
  param := hatena.SearchParameter{Query: "テスト"}
  feed := hatena.SearchTextRss(&param)
  for _, item := range feed.Items {
    fmt.Printf("%d - %s - %s\n", item.BookmarkCount, item.Title, item.Link)
  }
}

```

### search domain
http://b.hatena.ne.jp/entrylist?url=example.com&mode=rss&sort=popular

```go
package main

import (
    "github.com/yshimada0330/gohatena"
    "fmt"
)

func main() {
  param := hatena.EntryListParameter{Threshold: 3, Sort: "popular"}
  feed := hatena.EntryListRssByUrl("https://example.com", &param)
  for _, item := range feed.Items {
    fmt.Printf("%d - %s - %s\n", item.BookmarkCount, item.Title, item.Link)
  }
}

```
