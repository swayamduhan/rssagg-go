package models

import (
	"encoding/xml"
	"time"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title			string	 `xml:"title"`
	Link		 	string   `xml:"link"`
	Description		string  `xml:"description"`
	Items			[]Item	 `xml:"item"`
}

type Item struct {
	Title 		string 		`xml:"title"`
	Link 		string 		`xml:"link"`
	PublishDate	RSSPubDate	`xml:"pubDate"`
	Description string 	`xml:"description"`
}

type RSSPubDate struct {
	time.Time
}

func (r *RSSPubDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var pubDateStr string
	if err := d.DecodeElement(&pubDateStr, &start); err != nil {
		return err
	}

	// Try multiple formats
	layouts := []string{
		time.RFC1123,  
		time.RFC1123Z, 
		time.RFC3339,  
		"Mon, 2 Jan 2006 15:04:05 MST",
	}

	var parsedTime time.Time
	var err error
	for _, layout := range layouts {
		parsedTime, err = time.Parse(layout, pubDateStr)
		if err == nil {
			r.Time = parsedTime
			return nil
		}
	}

	return err
}