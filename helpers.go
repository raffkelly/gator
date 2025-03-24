package main

import (
	"database/sql"
	"log"
	"time"
)

func convertPubDate(pubDate string) sql.NullTime {
	if pubDate == "" {
		return sql.NullTime{Valid: false}
	}
	formats := []string{
		time.RFC1123,
		time.RFC1123Z,
		time.RFC822,
		time.RFC822Z,
		"Mon, 02 Jan 2006 15:04:05 -0700",
		"2006-01-02T15:04:05Z",      // ISO 8601
		"2006-01-02T15:04:05-07:00", // ISO 8601 with timezone
	}

	var publishedTime time.Time
	var err error

	for _, format := range formats {
		publishedTime, err = time.Parse(format, pubDate)
		if err == nil {
			publishedAt := sql.NullTime{
				Time:  publishedTime,
				Valid: true,
			}
			return publishedAt
		}
	}
	log.Printf("Could not parse publication date: %s\n", pubDate)
	return sql.NullTime{
		Valid: false,
	}
}
