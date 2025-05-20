package models

const (
	MinMorningSessionMinutes    = 180
	MinAfternoonSessionMinutes  = 180
	MaxAfternoonSessionMinutes  = 240
)

type Track struct {
	MorningTalks                 []Talk
	AfternoonTalks               []Talk
	MorningSessionCurrentDuration int
	AfternoonSessionCurrentDuration int
}

// NewTrack initializes a new Track object
func NewTrack() *Track {
	return &Track{
		MorningTalks:                 make([]Talk, 0),
		AfternoonTalks:               make([]Talk, 0),
		MorningSessionCurrentDuration: 0,
		AfternoonSessionCurrentDuration: 0,
	}
}

// TimeRemainingMorning returns the remaining time in minutes for the morning session.
func (t *Track) TimeRemainingMorning() int {
	return MinMorningSessionMinutes - t.MorningSessionCurrentDuration
}

// IsMorningFull checks if the morning session has reached its maximum capacity.
func (t *Track) IsMorningFull() bool {
	return t.MorningSessionCurrentDuration >= MinMorningSessionMinutes
}

// AddMorningTalk attempts to add a talk to the morning session.
// Returns true if the talk was added, false otherwise.
func (t *Track) AddMorningTalk(talk Talk) bool {
	if t.IsMorningFull() {
		return false
	}
	if talk.Duration <= t.TimeRemainingMorning() {
		t.MorningTalks = append(t.MorningTalks, talk)
		t.MorningSessionCurrentDuration += talk.Duration
		return true
	}
	return false
}

// TimeRemainingAfternoon returns the minimum and maximum remaining time in minutes for the afternoon session.
func (t *Track) TimeRemainingAfternoon() (minRemaining int, maxRemaining int) {
	minRemaining = MinAfternoonSessionMinutes - t.AfternoonSessionCurrentDuration
	maxRemaining = MaxAfternoonSessionMinutes - t.AfternoonSessionCurrentDuration
	return
}

// IsAfternoonFull checks if the afternoon session has reached its capacity.
// It's full if current duration is within the min and max limits.
func (t *Track) IsAfternoonFull() bool {
	return t.AfternoonSessionCurrentDuration >= MinAfternoonSessionMinutes &&
		t.AfternoonSessionCurrentDuration <= MaxAfternoonSessionMinutes
}

// AddAfternoonTalk attempts to add a talk to the afternoon session.
// Returns true if the talk was added, false otherwise.
// Considers that the talk should not make the session exceed MaxAfternoonSessionMinutes.
func (t *Track) AddAfternoonTalk(talk Talk) bool {
    // Cannot add if already over max 
	if t.AfternoonSessionCurrentDuration >= MaxAfternoonSessionMinutes {
		return false
	}
    
    _, maxRem := t.TimeRemainingAfternoon() // maxRem could be negative if current duration > MaxAfternoonSessionMinutes, but previous check handles this.

	if talk.Duration <= maxRem { // Check against maxRem to ensure we don't exceed MaxAfternoonSessionMinutes
		t.AfternoonTalks = append(t.AfternoonTalks, talk)
		t.AfternoonSessionCurrentDuration += talk.Duration
		return true
	}
	return false
}

// IsFull checks if both morning and afternoon sessions are full.
// Note: Afternoon session is considered "full" if it meets at least the minimum duration and not more than maximum.
func (t *Track) IsFull() bool {
	return t.IsMorningFull() && t.IsAfternoonFull()
}
```
