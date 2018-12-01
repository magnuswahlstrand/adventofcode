var fs = require("fs");

console.log("-----------------------------");
console.log("Advent of Code - Day 1 - Node");

function findDuplicateFrequency(filename) {
  var freq = 0;
  const visitedFrequencies = new Set([freq]);

  var lines = fs.readFileSync(filename, "utf8").split("\n");

  while (true) {
    for (let i = 0; i < lines.length; i++) {
      const line = lines[i];
      freq += parseInt(line);

      // Check if frequency already visited
      if (visitedFrequencies.has(freq)) {
        return freq;
      }

      // If not, add it to list
      visitedFrequencies.add(freq);
    }
  }
}

freq = findDuplicateFrequency("../input.txt");
console.log("First duplicate frequency is:", freq);
