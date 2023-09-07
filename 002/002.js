let asciiCount = [];
const input = require("fs").readFileSync("/dev/stdin", "utf8").trim().split('');

for (const index in input) {
    const target = asciiCount.find((ascii) => ascii.char === input[index])

    if (target) {
        target.count = target.count + 1
    } else {
        asciiCount.push({char: input[index], count: 1})
    }
}

asciiCount.sort((s, t) => (s.char < t.char ? -1 : 1))
asciiCount.forEach((elem, _) => {
    console.log(`${elem.char} => ${elem.count}`)
})
