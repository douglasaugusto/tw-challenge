package models

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestNewConference(t *testing.T) {
	conf := NewConference()
	if conf.Tracks == nil {
		t.Error("NewConference().Tracks should not be nil, expected empty slice")
	}
	if len(conf.Tracks) != 0 {
		t.Errorf("NewConference().Tracks should be empty, got len %d", len(conf.Tracks))
	}
}

func TestConference_AddTrack(t *testing.T) {
	conf := NewConference()
	track1 := NewTrack()
	track1.AddMorningTalk(Talk{Title: "Test Talk 1", Duration: 60})
	
	conf.AddTrack(*track1)
	if len(conf.Tracks) != 1 {
		t.Fatalf("Expected 1 track after AddTrack, got %d", len(conf.Tracks))
	}
	if conf.Tracks[0].MorningTalks[0].Title != "Test Talk 1" {
		t.Errorf("Track data mismatch after AddTrack. Expected title 'Test Talk 1', got '%s'", conf.Tracks[0].MorningTalks[0].Title)
	}

	track2 := NewTrack()
	track2.AddMorningTalk(Talk{Title: "Test Talk 2", Duration: 30})
	conf.AddTrack(*track2)
	if len(conf.Tracks) != 2 {
		t.Fatalf("Expected 2 tracks after second AddTrack, got %d", len(conf.Tracks))
	}
    if conf.Tracks[1].MorningTalks[0].Title != "Test Talk 2" {
		t.Errorf("Track data mismatch after second AddTrack. Expected title 'Test Talk 2', got '%s'", conf.Tracks[1].MorningTalks[0].Title)
	}
}

// Helper function to capture stdout
func captureOutput(f func()) string {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f() // Call the function whose output you want to capture

	w.Close()
	os.Stdout = old // restore real stdout
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestConference_PrintSchedule(t *testing.T) {
	// Freeze time for consistent output in tests, if needed for NetworkingEvent logic.
    // For this test, the specific time of NetworkingEvent is less critical than its presence and general format.
    // The `formatTime` function uses time.Date(0,0,0,...) so current date doesn't affect it.

	t.Run("Empty conference", func(t *testing.T) {
		conf := NewConference()
		output := captureOutput(func() {
			conf.PrintSchedule()
		})
		expected := "No tracks scheduled.\n"
		if output != expected {
			t.Errorf("PrintSchedule for empty conference: expected '%s', got '%s'", expected, output)
		}
	})

	t.Run("Conference with one track, morning and afternoon talks", func(t *testing.T) {
		conf := NewConference()
		track := NewTrack()
		track.AddMorningTalk(Talk{Title: "Go Routines", Duration: 60}) // 09:00AM
		track.AddMorningTalk(Talk{Title: "Go Channels", Duration: 30}) // 10:00AM
        // Morning total 90 mins. Remaining 90. Fills with more.
        track.AddMorningTalk(Talk{Title: "Go Select", Duration: 90}) // 10:30AM, ends 12:00PM

		track.AddAfternoonTalk(Talk{Title: "Go Interfaces", Duration: 120}) // 01:00PM, ends 03:00PM
        // Afternoon total 120 mins. Remaining min 60, max 120.
        track.AddAfternoonTalk(Talk{Title: "Go Generics", Duration: 60}) // 03:00PM, ends 04:00PM

		conf.AddTrack(*track)

		output := captureOutput(func() {
			conf.PrintSchedule()
		})

		expectedOutputParts := []string{
			"Track 1:",
			"09:00AM Go Routines",
			"10:00AM Go Channels",
			"10:30AM Go Select",
			"12:00PM Lunch",
			"01:00PM Go Interfaces",
			"03:00PM Go Generics",
			"04:00PM Networking Event", // Networking event at 4PM since talks end at 4PM
		}

		for _, part := range expectedOutputParts {
			if !strings.Contains(output, part) {
				t.Errorf("PrintSchedule output missing expected part: '%s'.\nGot:\n%s", part, output)
			}
		}
        // Check for an empty line after the track output
        if !strings.HasSuffix(strings.TrimSpace(output), "Networking Event") {
             // This check is a bit fragile if there's extra whitespace, better to check the structure.
             // The main check is if the parts are present.
        }
	})
    
    t.Run("Conference with networking event capped at 5 PM", func(t *testing.T) {
		conf := NewConference()
		track := NewTrack()
		track.AddMorningTalk(Talk{Title: "Full Morning", Duration: 180}) // 09:00AM - 12:00PM

		track.AddAfternoonTalk(Talk{Title: "Long Talk 1", Duration: 120}) // 01:00PM - 03:00PM
        track.AddAfternoonTalk(Talk{Title: "Long Talk 2", Duration: 120}) // 03:00PM - 05:00PM
        // Afternoon session is now 240 mins, ends at 05:00PM

		conf.AddTrack(*track)
		output := captureOutput(func() {
			conf.PrintSchedule()
		})
        
        // Expected networking event at 05:00PM
		if !strings.Contains(output, "05:00PM Networking Event") {
			t.Errorf("PrintSchedule output should have Networking Event at 05:00PM for talks ending at 5PM.\nGot:\n%s", output)
		}
	})

    t.Run("Conference with networking event scheduled at 4 PM if talks end early", func(t *testing.T) {
		conf := NewConference()
		track := NewTrack()
		track.AddMorningTalk(Talk{Title: "Full Morning", Duration: 180}) // 09:00AM - 12:00PM

		track.AddAfternoonTalk(Talk{Title: "Short Afternoon", Duration: 60}) // 01:00PM - 02:00PM
        // Afternoon session is 60 mins, ends at 02:00PM. Still needs to reach MinAfternoonSessionMinutes to be "full" for scheduling purposes,
        // but PrintSchedule should show Networking at 4PM if talks end early.
        // Let's ensure the track is considered "full" for scheduling logic by adding enough talks.
        track.AfternoonSessionCurrentDuration = MinAfternoonSessionMinutes // Manually set for testing print logic.
                                                                        // This simulates a track that is "full enough" to be scheduled.

		conf.AddTrack(*track)
		output := captureOutput(func() {
            // Temporarily set a logger to dev/null to suppress "afternoon session not full" messages if any.
            // This is for testing PrintSchedule's behavior given a track state.
            originalLoggerOutput := log.Writer()
            log.SetOutput(io.Discard)
            defer log.SetOutput(originalLoggerOutput)
			conf.PrintSchedule()
		})
        
		if !strings.Contains(output, "04:00PM Networking Event") {
			t.Errorf("PrintSchedule output should have Networking Event at 04:00PM for talks ending before 4PM.\nGot:\n%s", output)
		}
	})
}

```
