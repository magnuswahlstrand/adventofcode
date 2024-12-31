import { readFile } from "node:fs/promises"
let data = await readFile(import.meta.dirname + '/input.txt', 'utf8')
let testData = await readFile(import.meta.dirname + '/input_test.txt', 'utf8')

function isValid(line: number[]): boolean {
	let isAsc = false
	let isDesc = false
	for (let i = 0; i < line.length - 1; i++) {
		const diff = line[i] - line[i + 1]
		const absDiff = Math.abs(diff)

		if (absDiff < 1 || absDiff > 3) return false

		if (diff >= 1) {
			if (isDesc) return false
			isAsc = true
		}
		if (diff <= -1) {
			if (isAsc) return false
			isDesc = true
		}
	}

	return true
}
function isValidTest(input: number[], expected: boolean): any {
	console.log(input, expected == isValid(input))
}
function isValidTest2(input: number[], expected: boolean): any {
	console.log(input, expected == isValid2(input))
}
function parse(arg0: string): number[] {
	return arg0.split(' ').map(v => parseInt(v))
}
function isValid2(input: number[]) {
	for (let i = 0; i < input.length; i++) {
		if (isValid([...input.slice(0, i), ...input.slice(i + 1)])) return true

	}
	return false
}

const res = part1(testData.trim())
console.log(res)
isValidTest(parse('7 6 4 2 1'), true)
isValidTest(parse('1 2 7 8 9'), false)
isValidTest(parse('9 7 6 2 1'), false)
isValidTest(parse('1 3 2 4 5'), false)
isValidTest(parse('8 6 4 4 1'), false)
isValidTest(parse('1 3 6 7 9'), true)
function part1(testData: string) {
	const j = testData.split('\n')
	const rows = j.filter(v => isValid(parse(v)))

	return rows.length
}

const resPart1 = part1(data.trim())

console.log(resPart1)
isValidTest2(parse('7 6 4 2 1'), true)
isValidTest2(parse('1 2 7 8 9'), false)
isValidTest2(parse('9 7 6 2 1'), false)
isValidTest2(parse('1 3 2 4 5'), true)
isValidTest2(parse('8 6 4 4 1'), true)
isValidTest2(parse('1 3 6 7 9'), true)

function part2(testData: string) {
	const j = testData.split('\n')
	const rows = j.filter(v => isValid2(parse(v)))

	return rows.length
}
const resPart2 = part2(data.trim())
console.log(resPart2)
