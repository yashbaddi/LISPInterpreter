import valueParser from "./valueParser.js";
import specialForms from "./specialForms.js";
import { globalEnv } from "./symbols.js";

export default function expressionParser(input, env) {
  if (!input.startsWith("(")) return null;
  let operator = input.slice(1, input.indexOf(" "));
  let rem = input.slice(operator.length + 1).trim();
  let arr = parametersParser(rem);

  if (specialForms[operator]) {
    return [specialForms[operator](...arr, env), rem.slice(1)];
  }
  if (env[operator]) {
    arr = arr.map((ele) => valueParser(ele, env)[0]);
    return [env[operator](arr), rem.slice(1)];
  }

  throw new Error("Not a valid Operation");
}

export function parametersParser(input) {
  let arr = [];
  while (!input.startsWith(")")) {
    let pos = bracketsParser(input);
    arr.push(input.slice(0, pos));
    input = input.slice(pos).trim();
  }
  if (input[0] !== ")") return null;
  return arr;
}

function bracketsParser(input) {
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
