import valueParser, { symbolParser } from "./valueParser.js";
import specialForms from "./specialForms.js";

export default function expressionParser(env, input) {
  if (!input.startsWith("(")) return null;
  const arr = listParser(input);
  const operator = arr.shift();
  if (specialForms[operator]) {
    return [specialForms[operator](env, ...arr), ""];
  }
  const operatorFunction = valueParser(env, operator)[0];
  if (operatorFunction && typeof operatorFunction === "function") {
    return [expressionEvaluator(env, operatorFunction, arr), ""];
  }
  return null;
}

export function listParser(input) {
  if (!input.startsWith("(")) return null;
  input = input.trim().slice(1);
  const arr = [];
  while (!input.startsWith(")") && input !== "") {
    let pos = bracketsParser(input);
    arr.push(input.slice(0, pos));
    input = input.slice(pos).trim();
  }
  return arr;
}

function bracketsParser(input) {
  input = input.trim();
  let count = 0;
  let i = 0;
  while (
    count > 0 ||
    !(input[i] === " " || input[i] === undefined || input[i] === ")")
  ) {
    if (input[i] === "(") count++;
    if (input[i] === ")") count--;
    i++;
  }
  return i;
}

export function expressionEvaluator(env, operatorFunction, params) {
  params = params.map((ele) => valueParser(env, ele)[0]);
  return operatorFunction(params);
}
