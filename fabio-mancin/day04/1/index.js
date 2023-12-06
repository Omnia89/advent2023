import { readFileSync } from "fs";
const content = readFileSync("../input/input.txt");

const contentArray = content.toString().split("\n");
let sum = 0;
const regex = /: (.*)/;

try {
  contentArray.forEach((line) => {
    let matches = 0;
    const game = line.match(regex)[1].split(" | ");
    const winningNumbers = game[0].split(" ").filter(Boolean);
    const foundNumbers = game[1].split(" ").filter(Boolean);
    foundNumbers.forEach((number) => {
      if (winningNumbers.includes(number)) {
        if (matches === 0) {
          matches = 1;
        } else {
          matches *= 2;
        }
      }
    });
    sum += matches;
  })
} catch (error) {
  console.log(error);
}

console.log(sum)