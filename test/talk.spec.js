const expect = require('chai').expect;
const Talk = require("../src/models/talk").Talk;

let talk = new Talk('', 0);

describe('Talk', function() {
    describe('isLightning(talk)', function() {
        it('should return false when it does not include the word lightning', function() {
            expect(talk.isLightning("Accounting-Driven Development")).to.be.equal(false);
        });
        it('should return true when it does have include word lightning', function() {
            expect(talk.isLightning("Rails for Python Developers lightning")).to.be.equal(true);
        });
    });
    describe('getTitle(talk)', function() {
        it('should return the title without the word lightning', function() {
            expect(talk.getTitle("Rails for Python Developers lightning")).to.be.equal("Rails for Python Developers ");
        });
        it('should return the title without the duration', function() {
            expect(talk.getTitle("Sit Down and Write 30min")).to.be.equal("Sit Down and Write ");
        });
    });
    describe('getDuration(talk)', function() {
        it('should return 5 minutes when lightning', function() {
            expect(talk.getDuration("Rails for Python Developers lightning")).to.be.equal(5);
        });
        it('should return a positive integer', function() {
            expect(talk.getDuration("Pair Programming vs Noise 45min")).to.be.equal(45);
        });
    });
});
