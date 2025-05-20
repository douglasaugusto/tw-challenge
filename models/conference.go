package models

import (
	"fmt"
	"time"
)

const (
	StartHourMorning    = 9
	StartMinute         = 0
	StartHourAfternoon  = 13
	LunchString         = "12:00PM Lunch"
	NetworkingEventString = " Networking Event" // Note: The JS version has a leading space in the time for networking.
)

type Conference struct {
	Tracks []Track
}

// NewConference initializes a new Conference object
func NewConference() *Conference {
	return &Conference{
		Tracks: make([]Track, 0),
	}
}

// formatTime formats a time.Time object into HH:MM AM/PM string.
func formatTime(t time.Time) string {
	return t.Format("03:04PM")
}

// PrintSchedule outputs the conference schedule to the console.
func (c *Conference) PrintSchedule() {
	if len(c.Tracks) == 0 {
		fmt.Println("No tracks scheduled.")
		return
	}

	for i, track := range c.Tracks {
		fmt.Printf("Track %d:\n", i+1)

		// Morning Session
		currentTime := time.Date(0, 0, 0, StartHourMorning, StartMinute, 0, 0, time.UTC)
		for _, talk := range track.MorningTalks {
			fmt.Printf("%s %s\n", formatTime(currentTime), talk.Title)
			currentTime = currentTime.Add(time.Duration(talk.Duration) * time.Minute)
		}

		// Lunch
		fmt.Println(LunchString)

		// Afternoon Session
		currentTime = time.Date(0, 0, 0, StartHourAfternoon, StartMinute, 0, 0, time.UTC)
		for _, talk := range track.AfternoonTalks {
			fmt.Printf("%s %s\n", formatTime(currentTime), talk.Title)
			currentTime = currentTime.Add(time.Duration(talk.Duration) * time.Minute)
		}

		// Networking Event
        // Ensure networking event is not scheduled before 4 PM, but can be as late as 5 PM.
        // Based on JS logic, it's simply after the last talk.
        // The original JS code prints the time of the networking event.
        // If afternoon session ends at 4:30 PM, it prints 04:30PM Networking Event.
        // If it ends at 5:00 PM, it prints 05:00PM Networking Event.
        if currentTime.Hour() < 16 { // Before 4 PM
            currentTime = time.Date(0,0,0, 16, 0,0,0, time.UTC) // Schedule it at 4 PM if talks end earlier
        } else if currentTime.Hour() >= 17 && currentTime.Minute() > 0 { // After 5 PM (strictly after 5:00 PM)
             currentTime = time.Date(0,0,0, 17, 0,0,0, time.UTC) // Cap at 5PM
        }
		fmt.Printf("%s%s\n", formatTime(currentTime), NetworkingEventString)
		fmt.Println() // Empty line between tracks
	}
}

// AddTrack adds a new track to the conference.
func (c *Conference) AddTrack(track Track) {
    c.Tracks = append(c.Tracks, track)
}
```
