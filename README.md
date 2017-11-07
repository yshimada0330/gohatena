# gohatena
 Hatena Rss Search
 
# usage

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
