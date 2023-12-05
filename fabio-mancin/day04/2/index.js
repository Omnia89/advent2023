import { readFileSync } from "fs";
const content = readFileSync("../input/input.txt");

const contentArray = content.toString().split("\n");
let sum = 0;
const regex = / (.*): (.*)/;

try {
  const mapped = contentArray.reduce((acc, line) => {
    const game = line.match(regex);
    const gameNumber = game[1].trim();
    const gameData = game[2].split(" | ");
    const winningNumbers = gameData[0].split(" ").filter(Boolean);
    const foundNumbers = gameData[1].split(" ").filter(Boolean);
    return {
      ...acc,
      [gameNumber]: {
        winningNumbers,
        foundNumbers,
        q: 1,
      }
    }
  }, {})
  
  Object.entries(mapped).forEach(([gameNumber, gameData]) => {
    const matches = new Array(gameData.q).fill([]);
    let numberOfMatches = 0;
    gameData.foundNumbers.forEach((number) => {
      if (gameData.winningNumbers.includes(number)) {
        numberOfMatches += 1;
      }
    });
    matches.forEach((_, index) => {
      matches[index] = new Array(numberOfMatches).fill(true);
    })
    matches.forEach((match) => {
      match.forEach((_, numberIndex) => {
        mapped[parseInt(gameNumber) + numberIndex + 1].q += 1
      })
    })
  })
  const totalQuantity = Object.values(mapped).reduce((acc, game) => {
    return acc + game.q;
  }, 0)
  sum += totalQuantity;
} catch (error) {
  console.log(error);
}

console.log(sum)