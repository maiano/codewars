// Regex validate PIN code

function validatePIN (pin) {
  return /^(\d{4}|\d{6})$/.test(pin);
}

// Reverse words

function reverseWords(str) {
  return str.split(' ').map(word => [...word].reverse().join('')).join(' ');
}

// Shortest Word

function findShort(s){
  return Math.min(...s.split(' ').map(it => it.length));
}

// Exes and Ohs

function XO(str) {
  let xo = 0;
  for (const char of str.toLowerCase()) {
    if (char === 'x') xo++;
    else if (char === 'o') xo--;
  }
  return !xo;
}

// Highest and Lowest

function highAndLow(numbers){
  numbers = numbers.split(' ');
  return `${Math.max(...numbers)} ${Math.min(...numbers)}`;
}

// Disemvowel Trolls

function disemvowel(str) {
  const vowels = 'aeiouAEIOU';  
  return str.split('').filter(el => !vowels.includes(el)).join('');
}

// disemvowel = str => str.replace(/[aeiou]/gi,'');

// Isograms

function isIsogram(str){
  str = str.toLowerCase();
  return str.split('').every((el,i) => str.indexOf(el) === i);
}

// function isIsogram(str){
//   return !str.match(/([a-z]).*\1/i);
// }

// Digits explosion

function explode(s) {
  return [...s].map(el => el.repeat(el)).join('');
}

// const explode = s => s.replace(/\d/g, d => d.repeat(d));

// Human readable duration format

function formatDuration (seconds) {
  if (seconds === 0) return 'now';
  const time = { year: 31536000, day: 86400, hour: 3600, minute: 60, second: 1 };
  let str = [];
  for (var key in time) {
    if (time[key] <= seconds) {
      let temp = ~~(seconds/time[key]);
      str.push(temp += temp > 1 ? ' ' + key + 's' : ' ' + key);
      seconds = seconds % time[key];
    }
  }
  return str.length > 1 ? str.join(', ').replace(/,([^,]*)$/,' and'+'$1') : str[0];
}

// Duplicate Encoder

const duplicateEncode = word =>
    word.toLowerCase()
    .split('')
    .map ((el, _, array) => array.filter(e => e === el).length > 1 ? ')' : '(')
    .join('');

// function duplicateEncode(word) {
//   word = word.toLowerCase();
//   return word.replace(/./g, m => word.indexOf(m) == word.lastIndexOf(m) ? '(' : ')');
// }

// Multiples of 3 or 5

function solution(number){
     if (number < 0) {
          return 0;
     } else {
         return [...Array(number).keys()]
         .filter(it => it % 3 == 0 || it % 5 == 0)
         .reduce((prev, curr) => prev + curr, 0);
     }
}

// N-th Fibonacci

function nthFibo(n) {
  let [prev, curr] = [0, 1];
  for (let i = 1; i < n; i++) {
    [prev, curr] = [curr, prev + curr];
    }
  return prev;
}

// Javascript Mathematician

const calculate = (...a) => (...b) => [].concat(a, b).reduce((a, b) => a + b);

// Array Deep Count

function deepCount(a){
  return a
  .reduce((acc, val) => {
     return acc + (Array.isArray(val) ? deepCount(val) : 0);
   }, a.length);
}

// Length of missing array

// function getLengthOfMissingArray(arrayOfArrays) {
//   const len = (arrayOfArrays || [])
//     .map(a => a ? a.length : 0)
//     .sort((a, b) => a - b)
  
//   if (len.includes(0)) {
//     return 0
//   }
//   for (let i = 0; i < len.length - 1; i++) {
//     if (len[i] + 1 !== len[i + 1]) {
//       return len[i] + 1
//     }
//   }
//   return 0
// }

// Pair of gloves

function numberOfPairs(gloves) {
  const glovesColor = new Set(gloves);
  let pairs = 0;
  for (const i of glovesColor) {
    let count = 0;
    gloves.forEach(element => {
      if (element === i) count++
    });
    pairs += count >> 1
  }
  return pairs
}

// Sorting by bits

const sortByBit = arr => arr.sort((a, b) => {
  const countBits = d => [...d.toString(2)].map(el => Number(el)).reduce((x, y) => x + y);
  return countBits(a) - countBits(b) || a - b
})

// const sortByBit = arr => arr
//   .sort((a, b) => {
//     const countBits = d => d && d % 2 + countBits(d >> 1);
//     return countBits(a) - countBits(b) || a - b
//   }
// )



