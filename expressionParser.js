import valueParser, { symbolParser } from "./valueParser.js";
import specialForms from "./specialForms.js";

export default function expressionParser(env, input) {
  if (!input.startsWith("(")) return null;

  input = input.slice(1);

  //Check In Special Forms
  let operator;
  if ((operator = /^(define|if|quote|lambda)/.exec(input))) {
    return specialForms[operator[0]](env, input.slice(operator[0].length));
  }

  //If not a speical Forms
  //Get the operator
  const parseredOperator =
    expressionParser(env, input) || symbolParser(env, input);
  if (!parseredOperator) return [null, "Error Invalid Operator"];
  operator = parseredOperator[0];
  input = parseredOperator[1];

  //Get the operands
  const operands = [];
  while (!input.startsWith(")")) {
    const operandValue = valueParser(env, input);
    operands.push(operandValue[0]);
    input = operandValue[1].trim();
  }

  //Call the evaluator
  return [expressionEvaluator(env, operator, operands), input.slice(1)];
}

export function listParser(input) {
  //Takes a string as input and returns a list of strings which includes
  //different elements. It oprates just like lisp's Quote.
  input = input.trim();
  if (!input.startsWith("(")) return null;
  input = input.trim().slice(1);
  const arr = [];
  while (!input.startsWith(")") && input !== "") {
    const pos = bracketsParser(input);

    if (!pos) return null;

    arr.push(input.slice(0, pos));
    input = input.slice(pos).trim();
  }
  if (input == "") return null;

  return [arr, input.slice(1)];
}

export function bracketsParser(input) {
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

// console.log(listParser("( sad ( if ( = 1 1 ) 1 2 ) ( jkh  ( kjh dsa ) oiu ))"));
// [ 'sad', 'asd', '( jkh  ( kjh dsa ) oiu )']
