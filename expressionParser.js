import valueParser, { symbolParser } from "./valueParser.js";
import specialForms from "./specialForms.js";

export default function expressionParser(input, env) {
  if (!input.startsWith("(")) return null;
  let arr = parametersParser(input);
  let operator = arr.shift();
  if (specialForms[operator]) {
    return [specialForms[operator](...arr, env), ""];
  }
  operator = valueParser(operator, env);
  if (operator[0] != null) return [expressionEvaluator(operator, arr, env), ""];

  throw new Error("Not a valid Operator");
}

export function parametersParser(input) {
  if (!input.startsWith("(")) return null;
  input = input.trim().slice(1);
  let arr = [];
  while (!input.startsWith(")") && input !== "") {
    let pos = bracketsParser(input);
    arr.push(input.slice(0, pos));
    input = input.slice(pos).trim();
  }
  if (input[0] !== ")") throw new Error("Unequal Paranthisis");
  return arr;
}

function bracketsParser(input) {
  input = input.trim();
  let count = 0;
  let i = 0;
  while (count > 0 || !(input[i] === " " || input[i] === ")")) {
    if (input[i] === "(") count++;
    if (input[i] === ")") count--;
    if (input[i] == undefined) throw new Error("Unequal Paranthisis");
    i++;
  }
  return i;
}

function expressionEvaluator(operator, params, env) {
  params = params.map((ele) => valueParser(ele, env)[0]);
  return operator[0](params);
}
