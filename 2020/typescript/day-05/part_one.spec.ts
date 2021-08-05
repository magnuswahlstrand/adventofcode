import {expect} from 'chai';
import {calculateSeatId} from './part_one';

describe('2020 - Day 5 - Part One', () => {
    it('should get the seat id of the 1st example', () => {
        expect(calculateSeatId('BFFFBBFRRR')).to.equal(567);
    });
    it('should get the seat id of the 2nd example', () => {
        expect(calculateSeatId('FFFBBBFRRR')).to.equal(119);
    });
    it('should get the seat id of the 3rd example', () => {
        expect(calculateSeatId('BBFFBBFRLL')).to.equal(820);
    });
});
