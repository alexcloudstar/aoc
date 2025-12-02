import fs from 'fs';

const fileName = "part1.txt";

const contents = fs.readFileSync(fileName, {
    encoding: "utf8"
});


const contentsNumbers = contents.replaceAll("\n", "").split(",");
const invalidIds = [];

for(let i = 0; i < contentsNumbers.length; i++) {
    const lId = +contentsNumbers[i].split("-")[0]
    const rId = +contentsNumbers[i].split("-")[1]

    let prevNum = lId;

    for(let j = lId; j < rId; j++) {
        if(prevNum.toString().substring(0, prevNum.toString().split("").length / 2) ===
            prevNum.toString().substring(prevNum.toString().split("").length / 2)
        ) {
            invalidIds.push(prevNum);
        }


        prevNum = prevNum + 1;
    }

    if(rId.toString().substring(0, rId.toString().split("").length / 2) ===
        rId.toString().substring(rId.toString().split("").length / 2)
    ) {
        invalidIds.push(rId);
    }
}

const sum = invalidIds.reduce((acc, cum) => acc + cum, 0)

console.log(sum)
