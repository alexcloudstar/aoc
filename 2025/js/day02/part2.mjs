import fs from 'fs';

const fileName = 'part2.txt';

const contents = fs.readFileSync(fileName, {
  encoding: 'utf8',
});

const eachNumArr = contents.replaceAll('\n', '').split(',');

const invalidIds = [];

function isInvalid(num) {
  const s = num.toString();
  const L = s.length;

  for (let b = 1; b <= L / 2; b++) {
    if (L % b !== 0) continue;

    const block = s.slice(0, b);
    const repeated = block.repeat(L / b);

    if (repeated === s) return true;
  }

  return false;
}

for (let k = 0; k < eachNumArr.length; k++) {
  const ids = eachNumArr[k].split('-');
  const lNum = +ids[0];
  const rNum = +ids[1];

  for (let n = lNum; n <= rNum; n++) {
    if (isInvalid(n)) {
      invalidIds.push(n);
    }
  }
}

const sum = invalidIds.reduce((acc, num) => acc + num, 0);
console.log(sum);
