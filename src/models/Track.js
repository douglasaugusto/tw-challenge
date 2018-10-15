class Track {

    constructor(morningTalks, morningSessionCurrentDuration, afternoonTalks, afternoonSessionCurrentDuration) {
        this.morningTalks = morningTalks;
        this.morningSessionCurrentDuration = morningSessionCurrentDuration;
        this.afternoonTalks = afternoonTalks;
        this.afternoonSessionCurrentDuration = afternoonSessionCurrentDuration;
        this.constants = Object.freeze({
            MIN_MORNING_SESSION_MINUTES: 180,
            MIN_AFTERNOON_SESSION_MINUTES: 180,
            MAX_AFTERNOON_SESSION_MINUTES: 240
        });
    }

    timeRemainingToMorningSessionFull() {
        return this.constants.MIN_MORNING_SESSION_MINUTES - this.morningSessionCurrentDuration;
    }

    findTalkToMorningSession(talks, timeRemaining) {
        if (talks.some(talk => talk['duration'] === timeRemaining)) {
            return talks.findIndex(talk => talk['duration'] === timeRemaining);
        } else if (talks.some(talk => talk['duration'] < timeRemaining)) {
            return talks.findIndex(talk => talk['duration'] < timeRemaining);
        } else {
            return -1;
        }
    }

    isMorningSessionFull() {
        return this.constants.MIN_MORNING_SESSION_MINUTES == this.morningSessionCurrentDuration;
    }

    timeRemainingToAfternoonSessionFull() {
        let possibilities = [];
        possibilities.push(this.constants.MIN_AFTERNOON_SESSION_MINUTES - this.afternoonSessionCurrentDuration);
        possibilities.push(this.constants.MAX_AFTERNOON_SESSION_MINUTES - this.afternoonSessionCurrentDuration);
        return possibilities;
    }

    findTalkToAfternoonSession(talks, timeRemaining) {
        if (talks.some(talk => talk['duration'] === timeRemaining[0])) {
            return talks.findIndex(talk => talk['duration'] === timeRemaining[0]);
        } else if (talks.some(talk => talk['duration'] === timeRemaining[1])) {
            return talks.findIndex(talk => talk['duration'] === timeRemaining[1]);
        } else if (talks.some(talk => talk['duration'] < timeRemaining[0])) {
            return talks.findIndex(talk => talk['duration'] < timeRemaining[0]);
        } else if (talks.some(talk => talk['duration'] < timeRemaining[1])) {
            return talks.findIndex(talk => talk['duration'] < timeRemaining[1]);
        } else {
            return -1;
        }
    }

    isAfternoonSessionFull() {
        return this.constants.MIN_AFTERNOON_SESSION_MINUTES <= this.afternoonSessionCurrentDuration &&
            this.constants.MAX_AFTERNOON_SESSION_MINUTES >= this.afternoonSessionCurrentDuration;
    }

    isFull() {
        return this.isMorningSessionFull() && this.isAfternoonSessionFull();
    }

}

module.exports = {
    Track: Track
}
