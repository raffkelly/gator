package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}
	req.Header.Set("User-Agent", "gator")
	client := http.Client{
		Timeout: 10 * time.Second, // Set a 10-second timeout
	}
	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, nil
	}
	bodyXml := RSSFeed{}
	err = xml.Unmarshal(body, &bodyXml)
	if err != nil {
		return &bodyXml, err
	}
	bodyXml.Channel.Title = html.UnescapeString(bodyXml.Channel.Title)
	bodyXml.Channel.Description = html.UnescapeString(bodyXml.Channel.Description)
	for i, _ := range bodyXml.Channel.Item {
		bodyXml.Channel.Item[i].Title = html.UnescapeString(bodyXml.Channel.Item[i].Title)
		bodyXml.Channel.Item[i].Description = html.UnescapeString(bodyXml.Channel.Item[i].Description)
	}
	return &bodyXml, nil
}
