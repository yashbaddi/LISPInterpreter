import {
  listParser,
  bracketsParser,
  expressionEvaluator,
} from "./expressionParser.js";
import { globalEnv } from "./symbols.js";
import valueParser, { whiteSpaceParser } from "./valueParser.js";

const env = globalEnv;

let specialForms = {
  define: (env, input) => {
    if (!input.startsWith(" ")) return null; //Space After Operator
    input = whiteSpaceParser(input);
    const pos = input.indexOf(" ");
    const val = valueParser(env, whiteSpaceParser(input.slice(pos)));
    if (val === null) return null;
    env[input.slice(0, pos)] = val[0];
    return val; //handle errors
  },
  quote: (env, input) => {
    input = whiteSpaceParser(input);
    let quotedOutput;
    if ((quotedOutput = listParser(input))) {
      return quotedOutput;
    }
    const pos = bracketsParser(input);
    return [input.slice(0, pos), input.slice(pos)];
  },
  if: (env, input) => {
    let condition, value;
    if (!input.startsWith(" ")) return null;
    input = whiteSpaceParser(input);
    [condition, input] = valueParser(env, input);
    input = whiteSpaceParser(input);

    if (condition) {
      [value, input] = valueParser(env, input);

      if (!input.startsWith(" ")) return null;
      input = whiteSpaceParser(input);
      input = input.slice(bracketsParser(input));

      return [value, input]; //Correct Way
    }

    input = input.slice(bracketsParser(input));
    [value, input] = valueParser(env, input);
    return [value, input];
  },
  lambda: (env, input) => {
    let paramsArr;
    [paramsArr, input] = listParser(input);
    input = whiteSpaceParser(input);
    if (!paramsArr) return null;
    const defnition = input.slice(0, bracketsParser(input));

    let func = (params) => {
      const localEnv = Object.create(env);
      for (let i = 0; i < params.length; i++) {
        localEnv[paramsArr[i]] = params[i];
      }
      return [valueParser(localEnv, defnition)[0], paramsArr.length];
    };
    return [func, input.slice(defnition.length)];
  },
};

export default specialForms;

specialForms["define"](env, " x 2)");
specialForms["if"](env, " #t (if #f 3 2) 7)");
