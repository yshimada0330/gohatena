package hatena

import (
  "net/url"
  "strconv"
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

  path := SEARCH_TEXT_URL + "?" + values.Encode()
  return createFeed(path)
}
