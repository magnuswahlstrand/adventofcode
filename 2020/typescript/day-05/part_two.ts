import "../lib/index";
import "./part_one";
import {findSeat, seatIDFromSeat} from "./part_one";
import {readFileSync} from "fs";
import path from "path";

export function part2(input: string): number {
    const seats = []
    for (const seadId of input.split('\n').map(findSeat).map(seatIDFromSeat)) {
        seats[seadId] = seadId
    }
    while (seats[0] === undefined) {
        seats.shift()
    }
    return seats.findIndex((seat) => seat === undefined) + seats[0]
}

let input = readFileSync(path.join(__dirname, './input.txt'), 'utf-8');
console.log("Answer part 2:", part2(input));
