import { warn } from "node:console"
import { readFile } from "node:fs/promises"

const dataTest = (await readFile(import.meta.dirname + '/input_test.txt', 'utf8')).trimEnd()
const data = (await readFile(import.meta.dirname + '/input.txt', 'utf8')).trimEnd()

const patterns = [
	[[0, 0], [0, 1], [0, 2], [0, 3]],
	[[0, 0], [1, 1], [2, 2], [3, 3]],
	[[0, 0], [1, 0], [2, 0], [3, 0]],
	[[0, 3], [1, 2], [2, 1], [3, 0]],
]

patterns.push(patterns[0].toReversed())
patterns.push(patterns[1].toReversed())
patterns.push(patterns[2].toReversed())
patterns.push(patterns[3].toReversed())

function findXmas(input: string) {
	let sum = 0
	const g = input.split('\n').map(row => row.split(''))
	for (let y = 0; y < g.length; y++) {
		for (let x = 0; x < g[0].length; x++) {
			for (const [[x0, y0], [x1, y1], [x2, y2], [x3, y3]] of patterns) {
				if (
					g.at(y + y0)?.at(x + x0) +
					g.at(y + y1)?.at(x + x1) +
					g.at(y + y2)?.at(x + x2) +
					g.at(y + y3)?.at(x + x3) === 'XMAS'
				) {
					sum++
				}
			}
		}
	}
	return sum
}
const matches = [
	'MMSS',
	'SMMS',
	'SSMM',
	'MSSM'
]
function findXmasV2(input: string) {
	let sum = 0
	const g = input.split('\n').map(row => row.split(''))
	for (let y = 0; y < g.length; y++) {
		for (let x = 0; x < g[0].length; x++) {
			const c1 = g.at(y)?.at(x)
			const c2 = g.at(y)?.at(x + 2)
			const c3 = g.at(y + 2)?.at(x + 2)
			const c4 = g.at(y + 2)?.at(x)
			const c = c1 + c2 + c3 + c4
			const m = g.at(y + 1)?.at(x + 1)

			if (m === 'A' && matches.includes(c)) sum++
		}
	}
	return sum
}
console.log(findXmas("XMASXXMASM"))
const test = `XXXX
MMMM
AAAA
SSSS
AAAA
MMMM
XXXX`

console.log(findXmas(test))
console.log(findXmas(data))

console.log(findXmasV2(dataTest))
console.log(findXmasV2(test))
console.log(findXmasV2(data))
