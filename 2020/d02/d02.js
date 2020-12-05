const fs = require('fs');

const inputData = () => {
  const raw = fs.readFileSync('input', 'utf8');
  const rawList = raw.split(/\r?\n/);
  rawList.pop();
  return rawList.map(x => x.split(": "));
}

const partOne = (passwords) => {
  return passwords.filter(password => {
    const rule = password[0];
    const phrase = password[1];
    const min = rule.substring(0, rule.indexOf("-"));
    const max = rule.substring(rule.indexOf("-") + 1, rule.indexOf(" "));
    const req = rule.substr(rule.indexOf(" ") + 1);
    const occurances = phrase.split('').filter(x => x == req.valueOf()).length;
    return occurances >= min && occurances <= max;
  }).length;
}

const partTwo = (passwords) => {
  return passwords.filter(password => {
    const rule = password[0];
    const phrase = password[1];
    const firstSpot = rule.substring(0, rule.indexOf("-"));
    const secondSpot = rule.substring(rule.indexOf("-") + 1, rule.indexOf(" "));
    const req = rule.substr(rule.indexOf(" ") + 1);
    let matches = 0;
    if (phrase[firstSpot - 1] == req) { matches++; }
    if (phrase[secondSpot - 1] == req) { matches++; }
    return matches == 1;
  }).length;
}

const allPasswords = inputData();
console.log(`Part One: ${partOne(allPasswords)}`);
console.log(`Part Two: ${partTwo(allPasswords)}`);
