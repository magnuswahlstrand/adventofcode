import {readFileSync} from 'fs';

let input = readFileSync('./input.txt', 'utf-8');

interface PassportPolicy {
    low: number,
    high: number,
    character: string,
    passport: string,
}

let matcher = /^(\d+)-(\d+) (\w): (\w+)$/;

function parseLine(line: string): PassportPolicy {
    const [, lows, highs, character, passport] = line.match(matcher) ?? []
    const low = parseInt(lows, 10)
    const high = parseInt(highs, 10)
    return {low, high, character, passport}
}

function part1(input: string): number {
    return input.split('\n').map(parseLine).filter((p: PassportPolicy) => {
        const count = p.passport.split(p.character).length - 1
        return count >= p.low && count <= p.high
    }).length
}

function part2(input: string): number {
    return input.split('\n').map(parseLine).filter((p: PassportPolicy) => {
        return (p.passport.charAt(p.low - 1) == p.character) !== (p.passport.charAt(p.high - 1) == p.character)
    }).length
}

console.log("Answer part 1:", part1(input));
console.log("Answer part 2:", part2(input));
