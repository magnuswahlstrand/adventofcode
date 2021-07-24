import {entries} from './input'

let testEntries: number[];

testEntries = [
    1721,
    979,
    366,
    299,
    675,
    1456
]

function part1(entries: number[], target: number): number {
    for (let i = 0; i < entries.length; i++) {
        const first_number = entries[i];
        for (let j = i + 1; j < entries.length; j++) {
            const second_number = entries[j];
            if (first_number + second_number == target) {
                return first_number * second_number
            }
        }
    }
    throw 'Number not found';
}

function part2(entries: number[], target: number): number {
    for (let i = 0; i < entries.length; i++) {
        const first_number = entries[i];
        for (let j = i + 1; j < entries.length; j++) {
            const second_number = entries[j];
            for (let k = i; k < entries.length; k++) {
                if (k == i || k == j) {
                    continue
                }

                const third_number = entries[k];
                if (first_number + second_number + third_number == target) {
                    return first_number * second_number * third_number
                }
            }
        }
    }
    throw 'Number not found';
}

part1(testEntries, 2020);
console.log("Answer part 1:", part1(entries, 2020));

part2(testEntries, 2020);
console.log("Answer part 2:", part2(entries, 2020));
