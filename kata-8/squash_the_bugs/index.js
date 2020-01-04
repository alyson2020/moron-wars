function findLongest(str) {
  const spl = str.split(" ");
  let longest = 0

  for (let i = 0; i < spl.length; i++) {
    if (spl[i].length > longest) {
      longest = spl[i].length
    }
  }
  return longest
}

console.log(findLongest("The quick white fox jumped around the massive dog"), 7);
console.log(findLongest("Take me to tinseltown with you"), 10);
console.log(findLongest("Sausage chops"), 7);
console.log(findLongest("Wind your body and wiggle your belly"), 6);
console.log(findLongest("Lets all go on holiday"), 7);
