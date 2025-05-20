# Conference Track Management (Go Version)

This program schedules conference talks into tracks. It is a Go implementation of the original JavaScript version.

## Running the Application

To run the application, provide the path to a text file containing the list of talks:

```bash
# Make sure you have Go installed (version 1.18 or higher recommended).
# Run the application, replacing talks.txt with your input file if different:
go run main.go talks.txt
```

## Running Tests

To run the unit tests for the application:

```bash
# From the root directory of the project:
go test ./...
```

## Project Structure

-   `main.go`: The main application entry point. Handles file reading, scheduling logic, and output.
-   `models/`: This package contains the Go structs and methods for:
    -   `talk.go`: Defines the `Talk` struct and parsing logic.
    -   `track.go`: Defines the `Track` struct and session management logic.
    -   `conference.go`: Defines the `Conference` struct and schedule printing logic.
-   `talks.txt`: An example input file. (Note: This file might not exist yet, but it's good to mention as an example.)

## Original Problem Description (for context)

The original problem involved reading a list of talks with their durations and organizing them into tracks. Each track has a morning session and an afternoon session.
- Morning sessions are 3 hours long (9am - 12pm).
- Afternoon sessions are 3-4 hours long (1pm - 4pm or 5pm).
- Lunch is at 12pm.
- Networking event is scheduled after the last talk of the afternoon session, ideally between 4pm and 5pm. If talks finish before 4pm, networking is at 4pm. If talks finish after 5pm, networking is at 5pm.
- "lightning" talks are 5 minutes long.
```
