package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"conference-scheduler/models" // Adjust if your module name is different
)

// removeTalk removes a talk from a slice of talks at a given index.
func removeTalk(talks []models.Talk, index int) []models.Talk {
	if index < 0 || index >= len(talks) {
		return talks
	}
	return append(talks[:index], talks[index+1:]...)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <input_file>")
	}
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file %s: %v", filePath, err)
	}
	defer file.Close()

	var talks []models.Talk
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if line == "" { // Skip empty lines
			continue
		}
		talk, err := models.NewTalk(line)
		if err != nil {
			log.Printf("Error parsing talk on line %d ('%s'): %v. Skipping.", lineNumber, line, err)
			continue
		}
		talks = append(talks, *talk)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	if len(talks) == 0 {
		fmt.Println("No talks to schedule.")
		return
	}

	conference := models.NewConference()
	currentTrack := models.NewTrack()

	for len(talks) > 0 {
		scheduledThisPass := false

		if currentTrack.IsFull() {
			conference.AddTrack(*currentTrack)
			currentTrack = models.NewTrack()
		}

		// Try to fill morning session
		if !currentTrack.IsMorningFull() {
			for i := len(talks) - 1; i >= 0; i-- { // Iterate backwards for safe removal
				talk := talks[i]
				if currentTrack.AddMorningTalk(talk) {
					talks = removeTalk(talks, i)
					scheduledThisPass = true
					// Found a talk for the morning, might want to look for more or move to afternoon
					// Original JS logic implies one talk per "round" for a session then re-evaluates.
					// For simplicity here, we'll try to fill as much as possible in one go.
					// To better mimic JS, we might break here and restart the outer loop pass.
					// However, filling greedily per session before moving to next track is also valid.
				}
			}
		}

		// Try to fill afternoon session
		if !currentTrack.IsAfternoonFull() { // Check if it got full from morning additions
			for i := len(talks) - 1; i >= 0; i-- { // Iterate backwards
				talk := talks[i]
				// Important: Ensure afternoon session doesn't start filling until morning is truly full.
				// The AddAfternoonTalk should respect capacity.
				if currentTrack.IsMorningFull() && currentTrack.AddAfternoonTalk(talk) {
					talks = removeTalk(talks, i)
					scheduledThisPass = true
				}
			}
		}
        
        // If a track is partially filled (e.g. morning) but couldn't be entirely filled,
        // and no more talks can be added in this pass, it means we should seal this track
        // and start a new one, if there are talks left.
        // This handles cases where remaining talks don't fit the current track's remaining capacity perfectly.
		if !scheduledThisPass && len(talks) > 0 {
            // If morning has talks, or afternoon has talks, then it's a valid track to add.
            // Avoid adding empty tracks if all talks were scheduled perfectly into prior tracks.
            if len(currentTrack.MorningTalks) > 0 || len(currentTrack.AfternoonTalks) > 0 {
                 conference.AddTrack(*currentTrack)
                 currentTrack = models.NewTrack()
            }
		}
	}

	// Add the last track if it has talks and hasn't been added yet
	if len(currentTrack.MorningTalks) > 0 || len(currentTrack.AfternoonTalks) > 0 {
		conference.AddTrack(*currentTrack)
	}

	conference.PrintSchedule()
}
```
