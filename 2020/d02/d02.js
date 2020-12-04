const fs = require('fs');

const inputData = () => {
  const raw = fs.readFileSync('input', 'utf8');
  const rawList = raw.split(/\r?\n/);
  return rawList.map(x => x.split(": "));
}

const validPasswords = (passwords) => {
  passwords.pop();
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
const allPasswords = inputData();
console.log(`Number of valid passwords: ${validPasswords(allPasswords)}`);
