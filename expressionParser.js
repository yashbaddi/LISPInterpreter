import valueParser, { symbolParser } from "./valueParser.js";
import specialForms from "./specialForms.js";
import { globalEnv } from "./symbols.js";

export default function expressionParser(input, env) {
  if (!input.startsWith("(")) return null;
  // let operator = input.slice(1, input.indexOf(" "));
  // let rem = input.slice(operator.length + 1).trim();
  let arr = parametersParser(input);
  let operator = arr.shift();
  if (specialForms[operator]) {
    return [specialForms[operator](...arr, env), ""];
  }
  operator = valueParser(operator, env);

  return [expressionEvaluator(operator, arr, env), ""];

  throw new Error("Not a valid Operation");
}
function expressionEvaluator(operator, params, env) {
  params = params.map((ele) => valueParser(ele, env)[0]);
  return operator[0](params);
}
export function parametersParser(input) {
  if (!input.startsWith("(")) return null;
  input = input.trim().slice(1);
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
