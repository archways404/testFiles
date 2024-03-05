const vow = new Set('a', 'A', 'e', 'i', 'o', 'u');

const seta = vow.has('a');
const setA = vow.has('A');

console.log(seta);
console.log(setA);
