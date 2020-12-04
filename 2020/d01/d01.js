const fs = require('fs');

const inputData = () => {
  const raw = fs.readFileSync('input', 'utf8');
  return raw.split(/\r?\n/).map(x => parseInt(x)).filter(Boolean);
}

const partOne = (entries) => {
  console.log("Part One\n----");
  entries.forEach((a, i) => {
    entries.slice(i).forEach(b => {
      if (a + b == 2020) {
        console.log(`${a} + ${b} = ${a + b}`);
        console.log(`${a} * ${b} = ${a * b}`);
      }
    });
  });
}

const partTwo = (entries) => {
  console.log("Part Two\n----");
  entries.forEach((a, i) => {
    entries.slice(i).forEach((b, j) => {
      entries.slice(j).forEach(c => {
        if (a + b + c == 2020) {
          console.log(`${a} + ${b} + ${c} = ${a + b + c}`);
          console.log(`${a} * ${b} * ${c} = ${a * b * c}`);
        }
      });
    });
  });
}

partOne(inputData());
partTwo(inputData());
