const fs = require("node:fs")
// Part 1
function total(left, right) {
    left.sort()
    right.sort()
    indexes = left.length
    total = 0

    for (let i = 0; i < indexes; i++) {
        total += Math.abs(left[i] - right[i])
    }
    console.log(total)
}

// Part 2
function score(left, right) {
    score = 0
    rightMapped = {}
    right.forEach(element => {
        if (rightMapped.hasOwnProperty(element)) {
            rightMapped[element] += 1
        } else {
            rightMapped[element] = 1
        }
    })
    left.forEach(element => {
        if (rightMapped.hasOwnProperty(element)) {
            score += element * rightMapped[element]
        }
    })
    console.log(score)
}

function main() {
    const left = []
    const right = []
    fs.readFile('./inputday1.txt', 'utf8', (error, data) => {
        if (error) {
            console.log(error)
            return
        }
        data.split(/\r?\n/).forEach(line => {
            const numbers = line.split("   ")
            if (numbers.length == 2) {
                left.push(parseInt(numbers[0]))
                right.push(parseInt(numbers[1]))
            }
        })

        total(left, right)
        score(left, right)
    })
}

main()
