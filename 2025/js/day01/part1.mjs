import fs from "fs";

let pos = 50;
let clicks = 0;

const contents = fs.readFileSync("part2.txt")
const lines = contents.toString("utf-8").split("\n");

for (let i = 0; i < lines.length; i++) {
    const direction = lines[i][0];
    const num = +lines[i].substring(1);
    const start = pos;

    let hits = 0;


    if (direction === "R") {
        hits = Math.floor((start + num) / 100)
        pos = (((start + num) % 100) + 100) % 100
    }

    if (direction === "L") {
        if(start === 0) {
            hits = Math.floor(num / 100)
        }

        if (start !== 0 && num >= start) {
            hits = Math.floor((num - start) / 100) + 1
        }

        pos = (((start - num) % 100) + 100) % 100
    }

    clicks += hits
}

console.log("Clicks", clicks);
