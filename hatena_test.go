package hatena

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestSearchTextRss(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(rssHandler()))
  defer ts.Close()

  SEARCH_TEXT_URL = ts.URL

  param := SearchParameter{Query: "test", Sort: "popular", Threshold: 5, DateBegin: "2017-11-01", DateEnd: "2017-12-01", Safe: true}
  feed := SearchTextRss(&param)

  if feed.Items[0].Title != "タイトル1" {
    t.Fatalf("Error Title")
  }
  
  if feed.Items[1].Title != "タイトル2" {
    t.Fatalf("Error Title")
  }
}

func TestEntryListRssByUrl(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(rssHandler()))
  defer ts.Close()

  ENTRY_LIST_URL = ts.URL
  param := EntryListParameter{Threshold: 3, Sort: "recent"}
  feed := EntryListRssByUrl("https://example.com/", &param)

  if feed.Items[0].Title != "タイトル1" {
    t.Fatalf("Error Title")
  }
}

func rssHandler() func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`
<?xml version="1.0" encoding="UTF-8"?>
<rdf:RDF
 xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
>
  <channel rdf:about="http://b.hatena.ne.jp/search/text?safe=on&amp;sort=popular&amp;q=test&amp;date_begin=2017-10-01&amp;users=3">
    <title>本文「test」を検索（人気順） - はてなブックマーク</title>
    <link>http://b.hatena.ne.jp/search/text?safe=on&amp;sort=popular&amp;q=test&amp;date_begin=2017-10-01&amp;users=3</link>
    <description>本文「test」を検索（人気順） - はてなブックマーク</description>
    <items>
      <rdf:Seq>
        <rdf:li rdf:resource="http://hoge.com/entry/2017/10/31/001" />
        <rdf:li rdf:resource="http://hoge.com/entry/2017/10/03/002" />
      </rdf:Seq>
    </items>
  </channel>
  <item rdf:about="http://hoge.com/entry/2017/10/31/001">
    <title>タイトル1</title>
    <link>http://hoge.com/entry/2017/10/31/001</link>
    <description>あいうえお...</description>
    <content:encoded>&lt;blockquote cite=&quot;http://hoge.com/entry/2017/10/31/001&quot; title=&quot;タイトル1 - Blog&quot;&gt;&lt;cite&gt;&lt;img src=&quot;http://cdn-ak.favicon.st-hatena.com/?url=http%3A%2F%2Fhoge.com%2Fentry%2F2017%2F10%2F31%2F001&quot; alt=&quot;&quot; /&gt; &lt;a href=&quot;http://hoge.com/entry/2017/10/31/001&quot;&gt;タイトル1 - Blog&lt;/a&gt;&lt;/cite&gt;&lt;p&gt;&lt;a href=&quot;http://hoge.com/entry/2017/10/31/001&quot;&gt;&lt;img src=&quot;https://cdn-ak-scissors.b.st-hatena.com/image/square/beb285bcb4cb5e1b3aeff81722d55bde91590eea/height=90;version=1;width=120/https%3A%2F%2Fcdn-ak.f.st-hatena.com%2Fimages%2Ffotolife%2Ft%2Ftadashi-nemoto0713%2F20171026%2F20171026161553.png&quot; alt=&quot;タイトル1 - Blog&quot; title=&quot;タイトル1 - Blog&quot; class=&quot;entry-image&quot; /&gt;&lt;/a&gt;&lt;/p&gt;&lt;p&gt;2017 - 10 - 31 タイトル1 QA-SET Test Automation Slack こんにちは、ほげほげ...&lt;/p&gt;&lt;p&gt;&lt;a href=&quot;http://b.hatena.ne.jp/entry/http://hoge.com/entry/2017/10/31/001&quot;&gt;&lt;img src=&quot;http://b.hatena.ne.jp/entry/image/http://hoge.com/entry/2017/10/31/001&quot; alt=&quot;はてなブックマーク - タイトル1 - Blog&quot; title=&quot;はてなブックマーク - タイトル1 - Blog&quot; border=&quot;0&quot; style=&quot;border: none&quot; /&gt;&lt;/a&gt; &lt;a href=&quot;http://b.hatena.ne.jp/append?http://hoge.com/entry/2017/10/31/001&quot;&gt;&lt;img src=&quot;http://b.hatena.ne.jp/images/append.gif&quot; border=&quot;0&quot; alt=&quot;はてなブックマークに追加&quot; title=&quot;はてなブックマークに追加&quot; /&gt;&lt;/a&gt;&lt;/p&gt;&lt;/blockquote&gt;</content:encoded>
    <dc:date>2017-10-31T12:30:47+09:00</dc:date>
    <dc:subject>テクノロジー</dc:subject>
    <hatena:bookmarkcount>875</hatena:bookmarkcount>
  </item>
  <item rdf:about="http://hoge.com/entry/2017/10/03/002">
    <title>タイトル2</title>
    <link>http://hoge.com/entry/2017/10/03/002</link>
    <description>かきくけこ...</description>
    <content:encoded>&lt;blockquote cite=&quot;http://hoge.com/entry/2017/10/03/002&quot; title=&quot;タイトル2 - Blog&quot;&gt;&lt;cite&gt;&lt;img src=&quot;http://cdn-ak.favicon.st-hatena.com/?url=http%3A%2F%2Fhoge.com%2Fentry%2F2017%2F10%2F03%2F002&quot; alt=&quot;&quot; /&gt; &lt;a href=&quot;http://hoge.com/entry/2017/10/03/002&quot;&gt;タイトル2 - Blog&lt;/a&gt;&lt;/cite&gt;&lt;p&gt;&lt;a href=&quot;http://hoge.com/entry/2017/10/03/002&quot;&gt;&lt;img src=&quot;http://cdn-ak.b.st-hatena.com/entryimage/345691636-1506991929.jpg&quot; alt=&quot;タイトル2 - Blog&quot; title=&quot;タイトル2 - Blog&quot; class=&quot;entry-image&quot; /&gt;&lt;/a&gt;&lt;/p&gt;&lt;p&gt;2017 - 10 - 03 タイトル2 QA-SET Test Automation ほげほげを...&lt;/p&gt;&lt;p&gt;&lt;a href=&quot;http://b.hatena.ne.jp/entry/http://hoge.com/entry/2017/10/03/002&quot;&gt;&lt;img src=&quot;http://b.hatena.ne.jp/entry/image/http://hoge.com/entry/2017/10/03/002&quot; alt=&quot;はてなブックマーク - タイトル2 - Blog&quot; title=&quot;はてなブックマーク - タイトル2 - Blog&quot; border=&quot;0&quot; style=&quot;border: none&quot; /&gt;&lt;/a&gt; &lt;a href=&quot;http://b.hatena.ne.jp/append?http://hoge.com/entry/2017/10/03/002&quot;&gt;&lt;img src=&quot;http://b.hatena.ne.jp/images/append.gif&quot; border=&quot;0&quot; alt=&quot;はてなブックマークに追加&quot; title=&quot;はてなブックマークに追加&quot; /&gt;&lt;/a&gt;&lt;/p&gt;&lt;/blockquote&gt;</content:encoded>
    <dc:date>2017-10-03T09:51:28+09:00</dc:date>
    <dc:subject>テクノロジー</dc:subject>
    <hatena:bookmarkcount>235</hatena:bookmarkcount>
  </item>
</rdf:RDF>
    `))
  }
}