class Talk {

    constructor(title, duration) {
        this.title = title;
        this.duration = duration;
        this.constants = Object.freeze({
            LIGHTNING_SUBSTRING: "lightning",
            LIGHTNING_DURATION: 5
        });
    }

    isLightning(talk) {
        return talk.includes(this.constants.LIGHTNING_SUBSTRING);
    }

    getTitle(talk) {
        return this.isLightning(talk) ? talk.split(this.constants.LIGHTNING_SUBSTRING)[0] : talk.split(/([0-9]+)/)[0];
    }

    getDuration(talk) {
        return this.isLightning(talk) ? this.constants.LIGHTNING_DURATION : parseInt(talk.replace(/\D/g, ""));
    }

}

module.exports = {
    Talk: Talk
}
