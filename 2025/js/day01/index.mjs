import fs from "fs";

let pos = 50;
let clicks = 0;

const contents = fs.readFileSync("part1.txt")
const lines = contents.toString("utf-8").split("\n");

for (let i = 0; i < lines.length; i++) {
    const direction = lines[i][0];
    const num = +lines[i].substring(1) % 100;

    if (direction === "L") {
        pos = pos - num;
    }

    if (direction === "R") {
        pos = pos + num;
    }

    if (pos < 0) {
        pos = pos + 100;
    }

    if (pos > 99) {
        pos = pos % 100;
    }


    if (pos === 0) {
        clicks++;
    }
}

console.log("Clicks", clicks);
