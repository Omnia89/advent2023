import { readFileSync } from "fs";
const content = readFileSync("../input/input.txt");

let lowestLocation = null;
const contentArray = content.toString().split("\n");
const steps = [];
let currentDestination = null;
let currentOrigin = null;
const map = {};

const seeds = contentArray[0].split(": ")[1].split(" ");
const couplesOfSeeds = [];
for (let i = 0; i < seeds.length; i += 2) {
  couplesOfSeeds.push([parseInt(seeds[i]), parseInt(seeds[i]) + (parseInt(seeds[i + 1]) - 1)]);
}

const lineToOriginAndDestination = (line) => {
  const originAndDestination = line.split(" ")[0].split("-");
  if (!steps.includes(originAndDestination[0])) steps.push(originAndDestination[0]);
  if (!steps.includes(originAndDestination[1])) steps.push(originAndDestination[2]);
  return [originAndDestination[0], originAndDestination[2]];
}
const lineToValues = (line) => line.split(" ");
const createKeyInObjectIfNotExists = (object, key) => {
  if (!object[key]) {
    object[key] = [];
  }
}
const handleDefinitionLine = (line) => {
  const firstCharacterIsADigit = !isNaN(parseInt(line[0]));
  if (firstCharacterIsADigit) return;
  [currentOrigin, currentDestination] = lineToOriginAndDestination(line);
  createKeyInObjectIfNotExists(map, currentOrigin);
  createKeyInObjectIfNotExists(map[currentOrigin], currentDestination);
}
const handleValueLine = (line) => {
  const firstCharacterIsADigit = !isNaN(parseInt(line[0]));
  if (!firstCharacterIsADigit) return;
  const [destinationRangeStart, sourceRangeStart, rangeLength] = lineToValues(line);
  const ranges = [
    parseInt(sourceRangeStart),
    parseInt(sourceRangeStart) + (parseInt(rangeLength) - 1 ),
    parseInt(destinationRangeStart),
    parseInt(destinationRangeStart) + (parseInt(rangeLength) - 1 ),
  ];
  map[currentOrigin][currentDestination].push(ranges);
}
const findInObjectOrReturnSame = (array, key) => {
  let keyCopy = key;
  array.forEach(range => {
    if (key >= range[0] && key <= range[1]) {
      keyCopy = range[2] + (key - range[0]);
    }
  })
  return keyCopy;
}
const processSeed = (seed, map) => {
  return steps.reduce((acc, step, i) => {
    if (i === steps.length - 1) return acc;
    return findInObjectOrReturnSame(map[step][steps[i + 1]], acc);
  }, seed);
};
const reset = () => {
  currentDestination = null;
  currentOrigin = null;
}

try {
  for (let i = 1; i < contentArray.length - 1 ; i++) {
    const currentLine = contentArray[i];
    if (currentLine === "") {
      reset();
      continue;
    }
    handleDefinitionLine(currentLine);
    handleValueLine(currentLine);
  }

  couplesOfSeeds.forEach(couple => {
    let i = couple[0];
    while (i <= couple[1]) {
      const location = processSeed(i, map);
      if (!lowestLocation || location < lowestLocation) {
        lowestLocation = location;
      }
      i++;
    }
  })
  
  
} catch (error) {
  console.log(error);
}

console.log(lowestLocation)
