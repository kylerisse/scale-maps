var mc
var maps
var aps

function setup() {

    background(255)
    createCanvas(1920, 1080)
    frameRate(1)
    timer = 0

    maps = []
    mc = 0
    ballroom = loadImage('images/map_ballroom.png')
    maps.push(ballroom)
    expo = loadImage('images/map_expo.png')
    maps.push(expo)
    conupper = loadImage('images/map_conupper.png')
    maps.push(conupper)
    conlower = loadImage('images/map_conlower.png')
    maps.push(conlower)

    notifier = ""

    loadJSON("/api/aps", drawAPs)

}

function drawAPs(payload) {
    aps = []
    for (let i = 0; i < payload.aps.length; i++) {
        aps.push(payload.aps[i])
    }
}

function draw() {

    if (timer == 1) {
        image(maps[mc], 0, 0)
    }

    if (timer >= 10) {
        loadJSON("/api/aps", drawAPs)
        timer = 0
        mc++
    }

    if (mc >= maps.length) {
        mc = 0
    }

    text(timer, 1900, 20)
    timer++

    if ((aps !== undefined) && (aps.length > 0)) {
        for (let i = 0; i < aps.length; i++) {
            if (aps[i].map == parseInt(mc) + 1) {
                fill(240, 24, 48)
                textSize(14)
                text(aps[i].name, aps[i].x, aps[i].y)
            }
        }
    }

}

function mousePressed() {
    console.log("click " + mouseX + ", " + mouseY)
    textSize(16)
    text("click " + mouseX + ", " + mouseY, mouseX, mouseY)
}
