var fs = require("fs");

console.log("-----------------------------");
console.log("Advent of Code - Day 2 - Node");

var contents = fs.readFileSync("../input.txt", "utf8");
var ids = contents.split("\n");

function checksum(ids) {
  var twos = 0;
  var threes = 0;

  for (let i = 0; i < ids.length; i++) {
    const id = ids[i];
    charactersInID = {};
    twoIdenticalFound = false;
    threeIdenticalFound = false;

    for (let j = 0; j < id.length; j++) {
      charactersInID[id[j]] = (charactersInID[id[j]] || 0) + 1;
    }

    // Find characters that are either double or tripple
    for (const key of Object.keys(charactersInID)) {
      if (!twoIdenticalFound && charactersInID[key] == 2) {
        twos++;
        twoIdenticalFound = true; // Increment only once per ID
      }

      if (!threeIdenticalFound && charactersInID[key] == 3) {
        threes++;
        threeIdenticalFound = true; // Increment only once per ID
      }
    }
  }
  console.log(
    "Words with two characters: ",
    twos,
    "and three characters: ",
    threes
  );
  return twos * threes;
}

console.log("Checksum of IDs is:", checksum(ids));
