package models

import (
	"testing"
)

func TestNewTalk(t *testing.T) {
	tests := []struct {
		name         string
		line         string
		wantTitle    string
		wantDuration int
		wantErr      bool
	}{
		{
			name:         "Regular talk with duration",
			line:         "Writing Fast Tests Against Enterprise Rails 60min",
			wantTitle:    "Writing Fast Tests Against Enterprise Rails",
			wantDuration: 60,
			wantErr:      false,
		},
		{
			name:         "Lightning talk",
			line:         "Rails for Python Developers lightning",
			wantTitle:    "Rails for Python Developers",
			wantDuration: 5,
			wantErr:      false,
		},
		{
			name:         "Talk with numbers in title",
			line:         "Ruby on Rails 3: The Way of the Warrior 30min",
			wantTitle:    "Ruby on Rails 3: The Way of the Warrior",
			wantDuration: 30,
			wantErr:      false,
		},
		{
			name:         "Another regular talk",
			line:         "Sit Down and Write 30min",
			wantTitle:    "Sit Down and Write",
			wantDuration: 30,
			wantErr:      false,
		},
		{
			name:    "Empty line",
			line:    "",
			wantErr: true,
		},
		{
			name:    "Missing duration",
			line:    "A talk without duration",
			wantErr: true,
		},
		{
			name:    "Invalid duration format",
			line:    "A talk with invalid duration 60minutes",
			wantErr: true,
		},
		{
			name:         "Lightning talk with extra spaces",
			line:         "  Overdoing it in Python lightning  ",
			wantTitle:    "Overdoing it in Python",
			wantDuration: 5,
			wantErr:      false,
		},
        {
            name: "Talk title ending with number without min suffix",
            line: "Common Ruby Errors 45", // This should be an error as "45" is not "45min"
            wantErr: true,
        },
        {
            name: "Talk with only duration",
            line: "60min",
            wantErr: true, // Title would be empty
        },
        {
            name: "Talk with only lightning",
            line: "lightning",
            wantErr: true, // Title would be empty
        },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTalk, err := NewTalk(tt.line)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewTalk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if gotTalk.Title != tt.wantTitle {
					t.Errorf("NewTalk() gotTitle = %v, want %v", gotTalk.Title, tt.wantTitle)
				}
				if gotTalk.Duration != tt.wantDuration {
					t.Errorf("NewTalk() gotDuration = %v, want %v", gotTalk.Duration, tt.wantDuration)
				}
			}
		})
	}
}
```
