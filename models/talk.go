package models

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	lightningSubstring = "lightning"
	lightningDuration  = 5
)

type Talk struct {
	Title    string
	Duration int // in minutes
}

// NewTalk creates a new Talk instance from a raw string line.
// It parses the title and duration, handling "lightning" talks.
func NewTalk(line string) (*Talk, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, errors.New("input line is empty")
	}

	var title string
	var duration int

	if strings.Contains(line, lightningSubstring) { // Changed from HasSuffix to Contains to match JS version more closely
		duration = lightningDuration
		// Extract title by removing "lightning" and trimming spaces
		// This assumes "lightning" might not always be at the very end,
		// and aims to get the text part, then cleans it up.
		title = strings.TrimSpace(strings.Replace(line, lightningSubstring, "", 1))
	} else {
		// Use regex to find the duration like "45min" at the end of the string
		re := regexp.MustCompile(`(\d+)min$`)
		matches := re.FindStringSubmatch(line)

		if len(matches) < 2 {
			return nil, errors.New("could not parse duration from talk: " + line)
		}

		var err error
		duration, err = strconv.Atoi(matches[1])
		if err != nil {
			// This case should ideally not happen if regex matches a number
			return nil, errors.New("could not convert duration to integer: " + matches[1])
		}
		// The title is everything before the duration match
		title = strings.TrimSpace(line[:len(line)-len(matches[0])])
	}
    
    if title == "" {
        // If after processing, title is empty, it's an error.
        // This can happen if the line was just "lightning" or "60min".
        return nil, errors.New("parsed title is empty for line: " + line)
    }

	return &Talk{Title: title, Duration: duration}, nil
}
