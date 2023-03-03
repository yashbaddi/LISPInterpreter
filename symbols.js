import valueParser from "./valueParser.js";

export let globalEnv = {
  "+": (arr) => arr.reduce((acc, curr) => acc + curr, 0),
  "-": (arr) => arr.reduce((acc, curr) => acc - curr),
  "*": (arr) => arr.reduce((acc, curr) => acc * curr, 1),
  "/": (arr) => arr.reduce((acc, curr) => curr / acc),
  mod: (arr) => arr.reduce((acc, curr) => curr % acc),
  incf: (arr) => arr[0] + arr[1],
  decf: (arr) => arr[0] - arr[1],
  pi: Math.PI,
  ">": (arr) => arr[0] > arr[1],
  ">=": (arr) => arr[0] >= arr[1],
  "<": (arr) => arr[0] < arr[1],
  "<=": (arr) => arr[0] <= arr[1],
  "=": (arr) => arr[0] === arr[1],
  abs: (arr) => Math.abs(arr[0]),
  pow: (arr) => Math.pow(arr[0], arr[1]),
  "equal?": (arr) => arr[0] === arr[1],
  list: (arr) => arr,
  cons: (arr) => {
    if (Array.isArray(arr[1])) return [arr[0], ...arr[1]];
    return [arr[0], arr[1]];
  },
  car: (arr) => {
    if (arr[0].length > 0) return arr[0][0];
    return null;
  },
  cdr: (arr) => {
    if (arr[0].length > 0) return arr[0].slice(1);
    return null;
  },
  map: (arr) => arr[1].map((ele) => arr[0]([ele])),
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
