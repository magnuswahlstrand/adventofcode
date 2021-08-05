export {};

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

Number.prototype.isBetween = function (this: number, v1: number, v2: number) {
    return this >= v1 && this <= v2
}
