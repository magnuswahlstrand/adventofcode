import {readFileSync} from 'fs';

// Derived from https://github.com/AlexAegis/advent-of-code/blob/master/solutions/typescript/2020/04/part_one.ts

let input = readFileSync('./input.txt', 'utf-8');

export enum RelevantField {
    byr = 'byr',
    iyr = 'iyr',
    eyr = 'eyr',
    hgt = 'hgt',
    hcl = 'hcl',
    ecl = 'ecl',
    pid = 'pid',
}

export type Passport = Record<RelevantField, string>;

const isPassport = (passport: Partial<Passport>): passport is Passport =>
    Object.values(RelevantField).every((pf) => {
        return Object.keys(passport).find((k) => {
            return k === pf
        })
    });

function parsePassport(passportString: string): Passport {
    return passportString.split(' ').reduce((collector: Passport, current: string) => {
        const [key, value] = current.split(':')
        collector[key as RelevantField] = value
        return collector
    }, {} as Passport)
}

function part1(input: string): number {
    return input
        .replace(/(?<!\n)\n(?!\n)/g, ' ')
        .split('\n\n')
        .map(parsePassport)
        .filter(isPassport).length
}

declare global {
    interface String {
        toInt(radix?: number): number | undefined;
    }

    interface Number {
        isBetween(v1: number, v2: number): boolean
    }
}

String.prototype.toInt = function (this: string, radix = 10): number | undefined {
    const n = parseInt(this, radix)
    return isNaN(n) ? undefined : n
}

Number.prototype.isBetween = function (this: number, v1, v2: number) {
    return this >= v1 && this <= v2
}

const color = new Set(['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']);

function hasValidValues(p: Passport) {

    let byr = p.byr.toInt()?.isBetween(1920, 2002) ?? false
    let iyr = p.iyr.toInt()?.isBetween(2010, 2020) ?? false
    let eyr = p.eyr.toInt()?.isBetween(2020, 2030) ?? false

    let hgt: boolean
    if (p.hgt.endsWith('cm')) {
        hgt = p.hgt.slice(0, -2).toInt()?.isBetween(150, 193) ?? false
    } else if (p.hgt.endsWith('in')) {
        hgt = p.hgt.slice(0, -2).toInt()?.isBetween(59, 76) ?? false
    } else {
        hgt = false
    }

    let hcl = /^#[0-9a-f]{6}$/.test(p.hcl)
    let clr = color.has(p.ecl)
    let pid = /^[0-9]{9}$/.test(p.pid)
    // console.log(byr, iyr, eyr, hcl, hgt, clr, pid)
    return byr && iyr && eyr && hcl && hgt && clr && pid
}

function part2(input: string): number {
    return input
        .replace(/(?<!\n)\n(?!\n)/g, ' ')
        .split('\n\n')
        .map(parsePassport)
        .filter(isPassport)
        .filter(hasValidValues).length
}

const testInput = `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`

console.log("Answer part 1:", part1(input));
console.log("Answer part 2:", part2(input));
