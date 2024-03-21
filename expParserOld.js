export default function expressionParser(env, input) {
  if (!input.startsWith("(")) return null;

  // input = input.slice(1)
  // result = symbolParser(input)
  // input = result[1]
  // operator = input[0]
  // operands = []
  //
  // while(input[0] !== ')') {
  // const result = valueParser(input)
  // input = result[1]
  // operands.push(result[0])
  // }
  // if (input[0] !== ')') return null

  // return [operator(...operands), input.slice(1)]

  const arr = listParser(input);
  if (!arr) return [null, "Error:Unequal Paratisis"];

  const operator = arr.shift();
  if (specialForms[operator]) {
    return [specialForms[operator](env, ...arr), ""];
  }

  const operatorFunction = valueParser(env, operator)[0];
  if (operatorFunction && typeof operatorFunction === "function") {
    return expressionEvaluator(env, operatorFunction, arr);
  }

  return [null, "Error:Not a proper Operator"];
}
