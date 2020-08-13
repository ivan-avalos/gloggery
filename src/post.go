package main

import (
	"fmt"
	"log"
	"regexp"
	"time"
)

type post struct {
	ModifiedTime time.Time
	Filename     string
	URL          string
	Date         time.Time
	Title        string
	ReadBody     func() string
}

func newPost(folder string, item folderItem, indexURL string) *post {
	geminiFilename := fmt.Sprintf("%v.gmi", item.Filename)

	url := fmt.Sprintf("%v/%v", indexURL, geminiFilename)

	date, title := parseFilename(item.Filename)

	readBody := func() string {
		return string(readFile(folder, item.Filename))
	}

	return &post{
		ModifiedTime: item.ModifiedTime,
		Filename:     geminiFilename,
		URL:          url,
		Date:         date,
		Title:        title,
		ReadBody:     readBody,
	}
}

var filenameRegex = regexp.MustCompile("^(\\d{4}-\\d{2}-\\d{2})-(.*)")

func parseFilename(filename string) (date time.Time, title string) {
	matches := filenameRegex.FindStringSubmatch(filename)
	if len(matches) == 0 {
		log.Fatalf("can't parse date from post filename %v", filename)
	}

	date, err := time.Parse("2006-01-02", matches[1])
	if err != nil {
		log.Fatalf("can't parse date from post filename %v", filename)
	}

	title = matches[2]
	return
}

func (p *post) ShouldBuild(geminiFolder string) bool {
	geminiModifiedTime, ok := findModifiedTime(geminiFolder, p.Filename)
	if !ok {
		return true
	}

	return geminiModifiedTime.Before(p.ModifiedTime)
}

func (p *post) ReadableDate() string {
	return p.Date.Format("2 January 2006")
}

func (p *post) ISODate() string {
	return p.Date.Format(time.RFC3339)
}

func (p *post) Body() string {
	return p.ReadBody()
}
