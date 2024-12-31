import { warn } from "node:console"
import { readFile } from "node:fs/promises"
let data = await readFile(import.meta.dirname + '/input.txt', 'utf8')
let testData = await readFile(import.meta.dirname + '/input_test.txt', 'utf8')

const mulMatcher = /mul\((\d+),(\d+)\)/g
function parseAll(rows: string) {
	let sum = 0
	for (const match of rows.matchAll(mulMatcher)) {
		sum += parseInt(match[1]) * parseInt(match[2])

	}
	return sum
}

console.log(parseAll(testData))
console.log(parseAll(data))

const mulMatcherV2 = /do\(\)|don\'t\(\)|mul\((\d+),(\d+)\)/g
function parseAllV2(input: string) {
	let sum = 0
	let enabled = true
	for (const match of input.matchAll(mulMatcherV2)) {
		if (match[0] === 'do()') {
			enabled = true
		} else if (match[0] === 'don\'t()') {
			enabled = false
		} else if (enabled && match[0].startsWith('mul')) {

			sum += parseInt(match[1]) * parseInt(match[2])
		}
	}

	return sum

}

console.log(parseAllV2("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"))
console.log(parseAllV2(data))
