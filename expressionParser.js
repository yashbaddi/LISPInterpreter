import valueParser, { symbolParser, whiteSpaceParser } from "./valueParser.js";
import specialForms from "./specialForms.js";

export default function expressionParser(env, input) {
  if (!input.startsWith("(")) return null;

  //get the list of un evaluated operator and operands
  const parsedList = listParser(input);
  if (!parsedList) return [null, "Error:Unequal Paratisis"];
  const arr = parsedList[0];
  input = parsedList[1];

  //get the operator
  const operator = arr.shift();

  //Check for the special forms
  if (specialForms[operator]) {
    return [specialForms[operator](env, ...arr), input.slice(1)];
  }

  //Apply the operator on the operands
  const operatorFunction = valueParser(env, operator)[0];
  if (operatorFunction && typeof operatorFunction === "function") {
    return [expressionEvaluator(env, operatorFunction, arr), input.slice(1)];
  }

  return [null, "Error:Not a proper Operator"];
}

export function listParser(input) {
  //This works like a quote function in lisp
  if (!input.startsWith("(")) return null;
  input = whiteSpaceParser(input.slice(1));
  const arr = [];

  //Check for the each element in the array
  while (!input.startsWith(")") && input !== "") {
    input = whiteSpaceParser(input);
    const pos = bracketsParser(input);

    if (!pos == null) return null;

    arr.push(input.slice(0, pos));
    input = input.slice(pos);
  }

  return [arr, input.slice(1)];
}

function bracketsParser(input) {
  let count = 0;
  let i = 0;
  while (count > 0 || !(input[i] === " " || input[i] === ")")) {
    if (input[i] === "(") count++;
    if (input[i] === ")") count--;
    if (input[i] === undefined) return null;
    i++;
  }
  return i;
}

export function expressionEvaluator(env, operatorFunction, params) {
  params = params.map((ele) => valueParser(env, ele)[0]);
  const evaluatedExpression = operatorFunction(params);
  if (
    evaluatedExpression[1] === params.length ||
    evaluatedExpression[1] === undefined
  ) {
    return evaluatedExpression[0];
  }

  return [
    null,
    "Error:Required Params:" +
      evaluatedExpression[1] +
      " Given:" +
      params.length,
  ];
}

console.log(listParser("( sad ( if ( = 1 1 ) 1 2 ) ( jkh  ( kjh dsa ) oiu ))"));
// [ 'sad', 'asd', '( jkh  ( kjh dsa ) oiu )']
