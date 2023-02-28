import valueParser from "./valueParser.js";
import specialForms from "./specialForms.js";
import { symbolObj } from "./symbols.js";

export default function expressionParser(input) {
  if (!input.startsWith("(")) return null;
  let operator = input.slice(1, input.indexOf(" "));
  let rem = input.slice(operator.length + 1).trim();
  let arr = [];
  while (!rem.startsWith(")")) {
    let pos = parantisisSkip(rem);
    arr.push(rem.slice(0, pos));
    rem = rem.slice(pos).trim();
  }
  if (rem[0] !== ")") return null;

  if (specialForms[operator]) {
    return [specialForms[operator](...arr), rem.slice(1)];
  }

  if (symbolObj[operator]) {
    arr = arr.map((ele) => valueParser(ele)[0]);
    return [symbolObj[operator](arr), rem.slice(1)];
  }

  throw new Error("Not a valid Operation");
}

function parantisisSkip(input) {
  input = input.trim();
  let count = 0;
  let i = 0;
  while (count > 0 || !(input[i] === " " || input[i] === ")")) {
    if (input[i] === "(") count++;
    if (input[i] === ")") count--;
    i++;
  }
  return i;
}

// while (!rem.startsWith(")")) {
//   let val = valueParser(rem);
//   if (!val) return null;
//   rem = val[1].trim();
//   arr.push(val[0]);
// }
