const fs = require('fs');
const file = process.argv[2];

const Talk = require("./models/talk").Talk;
const Track = require("./models/track").Track;
const Conference = require("./models/conference").Conference;

fs.readFile(file, 'utf8', function (err, data) {

    if (err) throw err;

    // Load Talks
    let records = data.split("\n");
    let talks = [];
    records.forEach(function (currentTalk) {
        let talk = new Talk('', 0);
        talk.title = talk.getTitle(currentTalk);
        talk.duration = talk.getDuration(currentTalk);
        talks.push(talk);
    });

    // Fill Tracks
    let track = new Track([], 0, [], 0);
    let conference = new Conference([]);
    while (talks.length != 0) {
        if (!track.isFull()) {
            if (!track.isMorningSessionFull()) {
                let index = track.findTalkToMorningSession(talks, track.timeRemainingToMorningSessionFull());
                track.morningTalks.push(talks[index]);
                track.morningSessionCurrentDuration += talks[index].duration;
                talks.splice(index, 1);
            } else if (!track.isAfternoonSessionFull()) {
                let index = track.findTalkToAfternoonSession(talks, track.timeRemainingToAfternoonSessionFull());
                track.afternoonTalks.push(talks[index]);
                track.afternoonSessionCurrentDuration += talks[index].duration;
                talks.splice(index, 1);
            }
        } else {
            conference.tracks.push(track);
            track = new Track([], 0, [], 0);
        }
    };

    // Print Conference Schedule
    conference.printSchedule();

});
