import { readFile } from "node:fs/promises"
import { fileURLToPath } from "node:url"
let data = await readFile(import.meta.dirname + '/input.txt', 'utf8')
console.log(import.meta.filename);
console.log(import.meta.dirname);
// let data = `3   4
// 4   3
// 2   5
// 1   3
// 3   9
// 3   3`
const rows = data.split('\n')
const left: number[] = []
const right: number[] = []
const rightOccurances: Record<number, number> = {}
rows.filter(Boolean).forEach(row => {
	const [a, b] = row.split(/\s+/).map(x => parseInt(x, 10))
	// console.log(b)
	left.push(a)
	right.push(b)
	rightOccurances[b] = (rightOccurances[b] ?? 0) + 1
})
left.sort()
right.sort()

let sum = 0
let sum2 = 0
for (let i = 0; i < left.length; i++) {
	/* console.log(right[i]) */
	const diff = Math.abs(left[i] - right[i])
	/* console.log(diff) */
	sum += diff
	sum2 += (rightOccurances[left[i]] ?? 0) * left[i]
}

console.log(sum)
console.log(sum2)
