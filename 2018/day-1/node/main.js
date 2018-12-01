var fs = require("fs");

var contents = fs.readFileSync("../input.txt", "utf8");

console.log("-----------------------------");
console.log("Advent of Code - Day 1 - Node");

let freq = 0;
contents.split("\n").forEach(number => {
  freq += parseInt(number);
});

console.log("Final frequency is:", freq);
