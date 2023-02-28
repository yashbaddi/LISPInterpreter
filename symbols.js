import valueParser from "./valueParser.js";

export let symbolObj = {
  "+": (arr) => arr.reduce((acc, curr) => acc + curr, 0),
  "-": (arr) => arr.reduce((acc, curr) => acc - curr),
  "*": (arr) => arr.reduce((acc, curr) => acc * curr, 1),
  "/": (arr) => arr.reduce((acc, curr) => curr / acc),
  mod: (arr) => arr.reduce((acc, curr) => curr % acc),
  incf: (arr) => arr[0] + arr[1],
  decf: (arr) => arr[0] - arr[1],
  ">": (arr) => arr[0] > arr[1],
  ">=": (arr) => arr[0] >= arr[1],
  "<": (arr) => arr[0] < arr[1],
  "<=": (arr) => arr[0] <= arr[1],
  "=": (arr) => arr[0] === arr[1],
};

// export let symbolObj = {
//   "+": (arr) => arr.reduce((acc, curr) => acc + valueParser(curr)[0], 0),
//   "-": (arr) => arr.reduce((acc, curr) => acc - valueParser(curr)[0]),
//   "*": (arr) => arr.reduce((acc, curr) => acc * valueParser(curr)[0], 1),
//   "/": (arr) => arr.reduce((acc, curr) => valueParser(curr)[0] / acc),
//   mod: (arr) => arr.reduce((acc, curr) => valueParser(curr)[0] % acc),
//   incf: (arr) => valueParser(arr[0])[0] + valueParser(arr[1])[0],
//   decf: (arr) => valueParser(arr[0])[0] - valueParser(arr[1])[0],
//   ">": (arr) => valueParser(arr[0])[0] > valueParser(arr[1])[0],
//   ">=": (arr) => valueParser(arr[0])[0] >= valueParser(arr[1])[0],
//   "<": (arr) => valueParser(arr[0])[0] < valueParser(arr[1])[0],
//   "<=": (arr) => valueParser(arr[0])[0] <= valueParser(arr[1])[0],
//   "=": (arr) => valueParser(arr[0])[0] === valueParser(arr[1])[0],
// };
