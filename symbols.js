export const globalEnv = {
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
