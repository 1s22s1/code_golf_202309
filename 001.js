const input = require("fs").readFileSync("/dev/stdin", "utf8");

console.log(input.replaceAll(/\s+/g, ''));
