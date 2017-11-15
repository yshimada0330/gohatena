package hatena

import (
  "net/url"
  "strconv"
)

var ENTRY_LIST_URL string = "http://b.hatena.ne.jp/entrylist"

type EntryListParameter struct {
  Sort string
  Threshold int
}

func EntryListRssByUrl(domain string, param *EntryListParameter) (*RssFeed){
  values := url.Values{}
  values.Set("mode", "rss")
  values.Set("url", domain)

  if param.Sort != "" {
    values.Set("sort", param.Sort)
  }
  
  if param.Threshold > 0 {
    values.Set("threshold", strconv.Itoa(param.Threshold))
  }

  query := ENTRY_LIST_URL + "?" + values.Encode()
  return createFeed(query)
}
