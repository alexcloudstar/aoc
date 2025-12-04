import fs from "fs";

const fileName = "input.txt";
const contents = fs.readFileSync(fileName, "utf8");

const content = contents.split("\n").filter(file => file);

const greatestNumbers = []

const getGreatestNumber = (idx) => {
    const numbers = content[idx].split("").map(Number)

    let firstBest = numbers[0]
    let bestJolt = -Infinity

    for(let i = 1; i < numbers.length; i++) {
        const second = numbers[i]

        const currVal = firstBest * 10 + second

        if(currVal > bestJolt) {
            bestJolt = currVal
        }

        if(second > firstBest) {
            firstBest = second
        }
    }

    greatestNumbers.push(bestJolt)
}

for(let i = 0; i < content.length; i++) {
    getGreatestNumber(i)
}

console.log(greatestNumbers.reduce((acc, cum) => acc + cum, 0))
