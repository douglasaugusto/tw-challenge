class Conference {

    constructor(tracks) {
        this.tracks = tracks;
        this.constants = Object.freeze({
            START_HOUR_MORNING: 9,
            START_HOUR_AFTERNOON: 13,
            START_MINUTE: 0,
            MIN_AFTERNOON_SESSION_MINUTES: 180,
            MAX_AFTERNOON_SESSION_MINUTES: 240
        });
    }

    getCurrentHour(date) {
        let hour = (date.getHours() < 10 ? "0" : "") + date.getHours();
        let min = (date.getMinutes() < 10 ? "0" : "") + date.getMinutes();
        let ampm = (date.getHours() < 12 ? "AM" : "PM");
        return `${hour}:${min}${ampm}`
    }

    printCurrentTalk(date, talk) {
        console.log(`${this.getCurrentHour(date)} ${talk.title}`);
        date.setMinutes(date.getMinutes() + talk.duration);
    }

    printNetworkingEvent(date) {
        console.log(`${this.getCurrentHour(date)} Networking Event`);
    }

    printSchedule() {
        let date = new Date();
        this.tracks.forEach((track, index) => {
            console.log(`Track ${index + 1}:`);
            date.setHours(this.constants.START_HOUR_MORNING);
            date.setMinutes(this.constants.START_MINUTE);
            track.morningTalks.forEach(talk => {
                this.printCurrentTalk(date, talk);
            });
            console.log("12:00PM Lunch");
            date.setHours(this.constants.START_HOUR_AFTERNOON);
            date.setMinutes(this.constants.START_MINUTE);
            track.afternoonTalks.forEach(talk => {
                this.printCurrentTalk(date, talk);
            });
            this.printNetworkingEvent(date);
            console.log();
        });
    }

}

module.exports = {
    Conference: Conference
}
