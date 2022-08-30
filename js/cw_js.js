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

