const expect = require('chai').expect;
const Conference = require("../src/models/conference").Conference;
const Talk = require("../src/models/talk").Talk;

let conference = new Conference([]);

describe('Conference', function () {
    describe('getCurrentHour(date)', function () {
        it('should return the time in the correct format', function () {
            let date = new Date();
            date.setHours(10);
            date.setMinutes(30);
            expect(conference.getCurrentHour(date)).to.be.equal("10:30AM");
        });
    });
});
