const expect = require('chai').expect;
const Track = require("../src/models/track").Track;
const Talk = require("../src/models/talk").Talk;

let track = new Track([], 0, [], 0);

describe('Track', function () {
    describe('timeRemainingToMorningSessionFull()', function () {
        it('should return the difference between the available time and the time already occupied for the morning session', function () {
            track.morningSessionCurrentDuration = 120;
            expect(track.timeRemainingToMorningSessionFull()).to.be.equal(60);
        });
    });
    describe('findTalkToMorningSession()', function () {
        it('should return the index of the talk that best fits to complete the morning session', function () {
            let talks = [new Talk('Lua for the Masses', 30), new Talk('Communicating Over Distance', 60)];
            expect(track.findTalkToMorningSession(talks, 60)).to.be.equal(1);
        });
    });
    describe('isMorningSessionFull()', function () {
        it('should return that the morning session is not full', function () {
            track.morningSessionCurrentDuration = 120;
            expect(track.isMorningSessionFull()).to.be.equal(false);
        });
        it('should return that the morning session is full', function () {
            track.morningSessionCurrentDuration = 180;
            expect(track.isMorningSessionFull()).to.be.equal(true);
        });
    });
    describe('timeRemainingToAfternoonSessionFull()', function () {
        it('should return the difference between the minimum available time and the time already occupied for the afternoon session', function () {
            track.afternoonSessionCurrentDuration = 150;
            expect(track.timeRemainingToAfternoonSessionFull()[0]).to.be.equal(30);
        });
        it('should return the difference between the maximum available time and the time already occupied for the afternoon session', function () {
            track.afternoonSessionCurrentDuration = 150;
            expect(track.timeRemainingToAfternoonSessionFull()[1]).to.be.equal(90);
        });
    });
    describe('findTalkToAfternoonSession()', function () {
        it('should return the index of the talk that best fits to complete the minimum afternoon session', function () {
            let talks = [new Talk('Lua for the Masses', 60), new Talk('Communicating Over Distance', 30)];
            expect(track.findTalkToAfternoonSession(talks, [30, 90])).to.be.equal(1);
        });
        it('should return the index of the talk that best fits to complete the maximium afternoon session', function () {
            let talks = [new Talk('Lua for the Masses', 60), new Talk('Communicating Over Distance', 30)];
            expect(track.findTalkToAfternoonSession(talks, [-30, 30])).to.be.equal(1);
        });
    });
    describe('isAfternoonSessionFull()', function () {
        it('should return that the afternoon session is not full', function () {
            track.afternoonSessionCurrentDuration = 120;
            expect(track.isAfternoonSessionFull()).to.be.equal(false);
        });
        it('should return that the afternoon session is full', function () {
            track.afternoonSessionCurrentDuration = 240;
            expect(track.isAfternoonSessionFull()).to.be.equal(true);
        });
    });
    describe('isFull()', function () {
        it('should return that the track is not full', function () {
            track.morningSessionCurrentDuration = 180;
            track.afternoonSessionCurrentDuration = 0;
            expect(track.isFull()).to.be.equal(false);
        });
        it('should return that the track is full', function () {
            track.morningSessionCurrentDuration = 180;
            track.afternoonSessionCurrentDuration = 240;
            expect(track.isFull()).to.be.equal(true);
        });
    });
});
