let asciiCount = {};
const input = require("fs").readFileSync("/dev/stdin", "utf8").trim().split('');

for (const index in input) {
    if (input[index] in asciiCount) {
        asciiCount[input[index]]++;
    } else {
        asciiCount[input[index]] = 1;
    }
}

console.log(asciiCount)
