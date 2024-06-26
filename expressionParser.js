import valueParser, { symbolParser, whiteSpaceParser } from "./valueParser.js";
import specialForms from "./specialForms.js";

export default function expressionParser(env, input) {
  if (!input.startsWith("(")) return null;

  input = input.slice(1);
  input = whiteSpaceParser(input);

  //Check For Special Forms
  const operatorMatch = /^(define|if|quote|lambda)/.exec(input);
  if (operatorMatch) {
    const value = specialForms[operatorMatch[0]](
      env,
      input.slice(operatorMatch[0].length)
    );
    //Error Handling
    if (value == null)
      return [null, "Error:Wrong " + operatorMatch[0] + " Expression"];

    return [value[0], value[1].slice(1)];
  }

  //Not a speical Forms?
  //Get the operator
  const parseredOperator =
    expressionParser(env, input) || symbolParser(env, input);
  if (!parseredOperator || typeof parseredOperator[0] !== "function")
    return [null, "Error Invalid Operator"];
  const operator = parseredOperator[0];
  input = parseredOperator[1];

  //Get the operands
  const operands = [];
  while (!input.startsWith(")")) {
    const operandValue = valueParser(env, input);
    operands.push(operandValue[0]);
    input = operandValue[1];
    if (!input.startsWith(" ") && !input.startsWith(")"))
      return [null, "Invalid Spaces"];
    input = whiteSpaceParser(input);
  }

  //Call the evaluator
  return [expressionEvaluator(env, operator, operands), input.slice(1)];
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

export function listParser(input) {
  //Takes a string as input and returns a list of strings which includes
  //different elements. It oprates just like lisp's Quote.
  input = whiteSpaceParser(input);
  if (!input.startsWith("(")) return null;
  input = whiteSpaceParser(input).slice(1);
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

// console.log(listParser("( sad ( if ( = 1 1 ) 1 2 ) ( jkh  ( kjh dsa ) oiu ))"));
// [ 'sad', 'asd', '( jkh  ( kjh dsa ) oiu )']
// console.log(bracketsParser("hey there this is yash"));
