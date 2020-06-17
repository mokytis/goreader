package feed

import "encoding/xml"

type RSS struct {
	XMLName  xml.Name  `xml:"rss"`
	Version  float64   `xml:"version,attr"`
	Channels []Channel `xml:"channel"`
}

type Channel struct {
	XMLName       xml.Name `xml:"channel"`
	Title         string   `xml:"title"`
	Link          string   `xml:"link"`
	Description   string   `xml:"description"`
	Generator     string   `xml:"generator"`
	Copyright     string   `xml:"copyright"`
	LastBuildDate string   `xml:lastBuildDate"`
	Items         []Item   `xml:"item"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Guid        string   `xml:"guid"`
	Description string   `xml:"description"`
	Content     string   `xml:"content"`
	ContentType string   `xml:"string"`
}

func Parse(xmlbytes []byte) RSS {
	var rss RSS
	xml.Unmarshal(xmlbytes, &rss)
	return rss
}
