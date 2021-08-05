import "../lib/index";
import {readFileSync} from "fs";
import path from "path";

const matcher = /[FBLR]/g

export function findSeat(ticket: string) {
    const binaryTicket = ticket.replaceAll(matcher, (c) => (c == 'B' || c == 'R') ? '1' : '0')
    let row = binaryTicket.slice(0, -3).toInt(2) ?? 0
    let column = binaryTicket.slice(-3).toInt(2) ?? 0
    return {row, column};
}

export function seatIDFromSeat(t: { column: number; row: number }): number {
    return t.row * 8 + t.column;
}

export function calculateSeatId(ticket: string): number {
    return seatIDFromSeat(findSeat(ticket))
}

export function part1(input: string): number {
    return Math.max(...input.split('\n').map(calculateSeatId))
}


let input = readFileSync(path.join(__dirname, './input.txt'), 'utf-8');

console.log("Answer part 1:", part1(input));
