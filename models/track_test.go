package models

import (
	"reflect"
	"testing"
)

// Helper to create a talk for tests
func newTestTalk(title string, duration int) Talk {
	return Talk{Title: title, Duration: duration}
}

func TestTrack_MorningSession(t *testing.T) {
	t.Run("TimeRemainingMorning and IsMorningFull", func(t *testing.T) {
		track := NewTrack()
		if track.TimeRemainingMorning() != MinMorningSessionMinutes {
			t.Errorf("Expected initial morning remaining time to be %d, got %d", MinMorningSessionMinutes, track.TimeRemainingMorning())
		}
		if track.IsMorningFull() {
			t.Error("New track's morning session should not be full")
		}

		track.AddMorningTalk(newTestTalk("Test Talk 1", 60))
		if track.TimeRemainingMorning() != MinMorningSessionMinutes-60 {
			t.Errorf("Expected morning remaining time to be %d, got %d", MinMorningSessionMinutes-60, track.TimeRemainingMorning())
		}

		track.AddMorningTalk(newTestTalk("Test Talk 2", 120))
		if track.TimeRemainingMorning() != 0 {
			t.Errorf("Expected morning remaining time to be 0, got %d", track.TimeRemainingMorning())
		}
		if !track.IsMorningFull() {
			t.Error("Track's morning session should be full")
		}

		// Try adding another talk when full
		if track.AddMorningTalk(newTestTalk("Test Talk 3", 30)) {
			t.Error("Should not be able to add talk to a full morning session")
		}
	})

	t.Run("AddMorningTalk", func(t *testing.T) {
		track := NewTrack()
		talk1 := newTestTalk("Talk 1", 60)
		talk2 := newTestTalk("Talk 2", MinMorningSessionMinutes-30) // Too long with talk1
		talk3 := newTestTalk("Talk 3", MinMorningSessionMinutes)    // Fills session exactly

		if !track.AddMorningTalk(talk1) {
			t.Error("Failed to add talk1 to morning session")
		}
		if len(track.MorningTalks) != 1 || !reflect.DeepEqual(track.MorningTalks[0], talk1) {
			t.Errorf("Morning talks not updated correctly after adding talk1. Got %v", track.MorningTalks)
		}
		if track.MorningSessionCurrentDuration != 60 {
			t.Errorf("Morning duration not updated correctly. Got %d", track.MorningSessionCurrentDuration)
		}

		if track.AddMorningTalk(talk2) { // This talk should not fit
			t.Error("Added talk2 which should not have fit in the morning session")
		}
		if len(track.MorningTalks) != 1 { // Ensure talk2 was not added
			t.Errorf("Morning talks should still have 1 talk. Got %v", track.MorningTalks)
		}

		// Reset track and add talk3
		track = NewTrack()
		if !track.AddMorningTalk(talk3) {
			t.Error("Failed to add talk3 which exactly fills the morning session")
		}
		if !track.IsMorningFull() {
			t.Error("Morning session should be full after adding talk3")
		}
	})
}

func TestTrack_AfternoonSession(t *testing.T) {
	t.Run("TimeRemainingAfternoon and IsAfternoonFull", func(t *testing.T) {
		track := NewTrack()
		minRem, maxRem := track.TimeRemainingAfternoon()
		if minRem != MinAfternoonSessionMinutes || maxRem != MaxAfternoonSessionMinutes {
			t.Errorf("Expected initial afternoon remaining time to be min %d, max %d, got min %d, max %d",
				MinAfternoonSessionMinutes, MaxAfternoonSessionMinutes, minRem, maxRem)
		}
		if track.IsAfternoonFull() {
			t.Error("New track's afternoon session should not be full")
		}

		track.AddAfternoonTalk(newTestTalk("Test Talk A", 100))
		minRem, maxRem = track.TimeRemainingAfternoon()
		if minRem != MinAfternoonSessionMinutes-100 || maxRem != MaxAfternoonSessionMinutes-100 {
			t.Errorf("Expected afternoon remaining time to be min %d, max %d, got min %d, max %d",
				MinAfternoonSessionMinutes-100, MaxAfternoonSessionMinutes-100, minRem, maxRem)
		}
		if track.IsAfternoonFull() { // Not full yet, below min threshold
			t.Error("Afternoon session should not be full yet")
		}

		track.AddAfternoonTalk(newTestTalk("Test Talk B", 80)) // Total 180
		if !track.IsAfternoonFull() {
			t.Error("Afternoon session should be full (at min duration)")
		}

		track.AddAfternoonTalk(newTestTalk("Test Talk C", 60)) // Total 240
		if !track.IsAfternoonFull() {
			t.Error("Afternoon session should be full (at max duration)")
		}
		minRem, maxRem = track.TimeRemainingAfternoon()
		// Corrected expectations for minRem based on TimeRemainingAfternoon logic
		// If current duration is 240, MinAfternoonSessionMinutes - 240 = 180 - 240 = -60
		// MaxAfternoonSessionMinutes - 240 = 240 - 240 = 0
		// The TimeRemainingAfternoon function in track.go does not cap negative minRemaining at 0.
		expectedMinRem := MinAfternoonSessionMinutes - 240
		expectedMaxRem := MaxAfternoonSessionMinutes - 240
		if minRem != expectedMinRem || maxRem != expectedMaxRem {
			t.Errorf("Expected afternoon remaining time to be min %d, max %d, got min %d, max %d",
				expectedMinRem, expectedMaxRem, minRem, maxRem)
		}


		// Try adding another talk when at max capacity
		if track.AddAfternoonTalk(newTestTalk("Test Talk D", 30)) {
			t.Error("Should not be able to add talk when afternoon session is at max capacity")
		}
        
        // Test case: adding a talk that would exceed MaxAfternoonSessionMinutes but not MinAfternoonSessionMinutes
        track = NewTrack()
        track.AddAfternoonTalk(newTestTalk("Large Talk", MaxAfternoonSessionMinutes - 10)) // 230 mins
        if !track.IsAfternoonFull() { // Should be full because 230 is between Min (180) and Max (240)
            t.Error("Afternoon should be considered full if it meets min criteria and is within max")
        }
        if track.AddAfternoonTalk(newTestTalk("Too Large Talk", 15)) { // Takes it to 245 mins
            t.Error("Should not add talk that makes session exceed MaxAfternoonSessionMinutes")
        }

	})

	t.Run("AddAfternoonTalk", func(t *testing.T) {
		track := NewTrack()
		talkA := newTestTalk("Talk A", 60)
		// Talk that would make it exceed max if session already had MinAfternoonSessionMinutes
		talkB_tooLarge := newTestTalk("Talk B too large", (MaxAfternoonSessionMinutes - MinAfternoonSessionMinutes) + 10) // e.g. 240-180=60, so 70
		talkC_fitsMax := newTestTalk("Talk C fits max", MaxAfternoonSessionMinutes)
        talkD_fitsMin := newTestTalk("Talk D fits min", MinAfternoonSessionMinutes)


		if !track.AddAfternoonTalk(talkA) {
			t.Error("Failed to add talkA to afternoon session")
		}
		if len(track.AfternoonTalks) != 1 || !reflect.DeepEqual(track.AfternoonTalks[0], talkA) {
			t.Errorf("Afternoon talks not updated correctly. Got %v", track.AfternoonTalks)
		}

        track.AfternoonSessionCurrentDuration = MinAfternoonSessionMinutes // Simulate session has some talks (total 180)
        // talkB_tooLarge has duration (240-180)+10 = 70.
        // Current duration 180 + 70 = 250, which is > MaxAfternoonSessionMinutes (240)
		if track.AddAfternoonTalk(talkB_tooLarge) {
			t.Errorf("Added talkB_tooLarge which should have made session exceed MaxAfternoonSessionMinutes. Current duration: %d, Talk duration: %d", track.AfternoonSessionCurrentDuration, talkB_tooLarge.Duration)
		}
        
        track = NewTrack()
        if !track.AddAfternoonTalk(talkD_fitsMin) {
            t.Error("Failed to add talkD_fitsMin")
        }
        if !track.IsAfternoonFull() {
            t.Error("Track should be full after adding talkD_fitsMin")
        }
		if track.AfternoonSessionCurrentDuration != MinAfternoonSessionMinutes { // Check exact duration
			t.Errorf("Duration should be MinAfternoonSessionMinutes, got %d", track.AfternoonSessionCurrentDuration)
		}


        track = NewTrack()
        if !track.AddAfternoonTalk(talkC_fitsMax) {
            t.Error("Failed to add talkC_fitsMax")
        }
        if !track.IsAfternoonFull() {
            t.Error("Track should be full after adding talkC_fitsMax")
        }
        if track.AfternoonSessionCurrentDuration != MaxAfternoonSessionMinutes {
            t.Errorf("Duration should be MaxAfternoonSessionMinutes, got %d", track.AfternoonSessionCurrentDuration)
        }
	})
}

func TestTrack_IsFull(t *testing.T) {
	track := NewTrack()
	if track.IsFull() {
		t.Error("New track should not be full")
	}

	// Fill morning session
	track.AddMorningTalk(newTestTalk("Morning Filler", MinMorningSessionMinutes))
	if !track.IsMorningFull() {
		t.Error("Morning session should be full")
	}
	if track.IsFull() {
		t.Error("Track should not be full if only morning is full")
	}

	// Fill afternoon session to minimum
	track.AddAfternoonTalk(newTestTalk("Afternoon Min Filler", MinAfternoonSessionMinutes))
	if !track.IsAfternoonFull() { // IsAfternoonFull checks >= Min && <= Max
		t.Errorf("Afternoon session should be full (at min). Current duration: %d", track.AfternoonSessionCurrentDuration)
	}
	if !track.IsFull() {
		t.Error("Track should be full if morning is full and afternoon is at min duration")
	}

	// Add more to afternoon up to max
    track = NewTrack() // Reset track
    track.AddMorningTalk(newTestTalk("Morning Filler", MinMorningSessionMinutes))
	track.AddAfternoonTalk(newTestTalk("Afternoon Max Filler", MaxAfternoonSessionMinutes))
	if !track.IsAfternoonFull() {
		t.Errorf("Afternoon session should be full (at max). Current duration: %d", track.AfternoonSessionCurrentDuration)
	}
	if !track.IsFull() {
		t.Error("Track should be full if morning is full and afternoon is at max duration")
	}
}
```
