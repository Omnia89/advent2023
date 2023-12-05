import { readFileSync } from "fs";
const content = readFileSync("../input/1.test.txt");

const isNumeric = (str) => {
  if (!str) return false;
  if (!isNaN(str)) return true;
  if (typeof str != "string") return false
  return !isNaN(str) && !isNaN(parseInt(str)) 
}

const checkCharacters = (line, index, before) => {
  let number = "";
  let currentIndex = before === "before" ? index - 1 : index + 1;
  while (isNumeric(line[currentIndex])) {
    number = before === "before" ? `${line[currentIndex]}${number}` : `${number}${line[currentIndex]}`;
    currentIndex = before === "before" ? currentIndex - 1 : currentIndex + 1;
  }
  return number;
  
}

const contentArray = content.toString().split("\n");
let sum = 0;

try {
  contentArray.forEach((line, lineIndex) => {
    line.split("").forEach((char, charIndex) => {
      if (char === "*") {
        let numbers = [];
        if (isNumeric(line[charIndex - 1])) {
          numbers.push(parseInt(checkCharacters(line, charIndex, "before")));
        }
        if (isNumeric(line[charIndex+1])) {
          numbers.push(parseInt(checkCharacters(line, charIndex, "after")));
        }
        const lineBefore = contentArray[lineIndex - 1];
        const lineAfter = contentArray[lineIndex + 1];
        if (lineBefore) {
          if (isNumeric(lineBefore[charIndex])) {
            if (isNumeric(lineBefore[charIndex - 1]) && isNumeric(lineBefore[charIndex + 1])) {
              let tempNumber = lineBefore[charIndex];
              tempNumber = `${checkCharacters(lineBefore, charIndex, "before")}${tempNumber}`;
              tempNumber = `${tempNumber}${checkCharacters(lineBefore, charIndex, "after")}`;
              numbers.push(parseInt(tempNumber));
            } else if (isNumeric(lineBefore[charIndex - 1]) || isNumeric(lineBefore[charIndex + 1])) {
              if (isNumeric(lineBefore[charIndex - 1])) {
                let tempNumber = lineBefore[charIndex];
                tempNumber = `${checkCharacters(lineBefore, charIndex, "before")}${tempNumber}`;
                numbers.push(parseInt(tempNumber));
              }
              if (isNumeric(lineBefore[charIndex + 1])) {
                let tempNumber = lineBefore[charIndex];
                tempNumber = `${tempNumber}${checkCharacters(lineBefore, charIndex, "after")}`;
                numbers.push(parseInt(tempNumber));
              }
            } else {
              numbers.push(parseInt(lineBefore[charIndex]));
            }
          } else {
            if (isNumeric(lineBefore[charIndex - 1])) {
              numbers.push(parseInt(checkCharacters(lineBefore, charIndex, "before")));
            }
            if (isNumeric(lineBefore[charIndex + 1])) {
              numbers.push(parseInt(checkCharacters(lineBefore, charIndex, "after")));
            }
          }
        }     
        if (lineAfter) {
          if (isNumeric(lineAfter[charIndex])) {
            if (isNumeric(lineAfter[charIndex - 1]) && isNumeric(lineAfter[charIndex + 1])) {
              let tempNumber = lineAfter[charIndex];
              tempNumber = `${checkCharacters(lineAfter, charIndex, "before")}${tempNumber}`;
              tempNumber = `${tempNumber}${checkCharacters(lineAfter, charIndex, "after")}`;
              numbers.push(parseInt(tempNumber));
            } else if (isNumeric(lineAfter[charIndex - 1]) || isNumeric(lineAfter[charIndex + 1])) {
              if (isNumeric(lineAfter[charIndex - 1])) {
                let tempNumber = lineAfter[charIndex];
                tempNumber = `${checkCharacters(lineAfter, charIndex, "before")}${tempNumber}`;
                numbers.push(parseInt(tempNumber));
              }
              if (isNumeric(lineAfter[charIndex + 1])) {
                let tempNumber = lineAfter[charIndex];
                tempNumber = `${tempNumber}${checkCharacters(lineAfter, charIndex, "after")}`;
                numbers.push(parseInt(tempNumber));
              }
            } else {
              numbers.push(parseInt(lineAfter[charIndex]));
            }
          } else {
            if (isNumeric(lineAfter[charIndex - 1])) {
              numbers.push(parseInt(checkCharacters(lineAfter, charIndex, "before")));
            }
            if (isNumeric(lineAfter[charIndex + 1])) {
              numbers.push(parseInt(checkCharacters(lineAfter, charIndex, "after")));
            }
          }
        }
        if (numbers.length === 2) {
          const multiplication = numbers[0] * numbers[1];
          sum += multiplication;
        }
        numbers = [];
      }
    })
  })
} catch (error) {
  console.log(error);
}

console.log(sum)