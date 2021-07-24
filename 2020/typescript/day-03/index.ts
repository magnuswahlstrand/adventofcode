import {readFileSync} from 'fs';

let input = readFileSync('./input.txt', 'utf-8');

function part1(input: string): number {
    return traverse(input, 3, 1)
}

function traverse(input: string, dx: number, dy: number): number {
    let x = 0;
    let res = input.split('\n')
        .filter((row, y) => {
            return y % dy === 0
        }).filter((row, y) => {
            const isTree = row[x] === '#'
            x = (x + dx) % row.length
            return isTree
        }).length
    return res
}

function part2(input: string): number {
    return traverse(input, 1, 1) *
        traverse(input, 3, 1) *
        traverse(input, 5, 1) *
        traverse(input, 7, 1) *
        traverse(input, 1, 2)
}

const testInput = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

console.log("Answer part 1:", part1(input));
console.log("Answer part 2:", part2(input));
