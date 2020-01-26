function minSum(arr) {
  const sortedArr = numSort(arr)
  let sum = 0

  while (sortedArr.length > 0) {
    sum = sum + sortedArr.shift() * sortedArr.pop()
  }
  return sum
}

const numSort = arr => {
  return arr.sort((a, b) => a - b)
}

console.log(minSum([5,4,2,3]), 22);
console.log(minSum([12,6,10,26,3,24]), 342);
console.log(minSum([9,2,8,7,5,4,0,6]), 74);

module.exports= minSum
