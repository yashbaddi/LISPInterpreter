import {
  listParser,
  bracketsParser,
  expressionEvaluator,
} from "./expressionParser.js";
import { globalEnv } from "./symbols.js";
import valueParser from "./valueParser.js";

const env = globalEnv;

let specialForms = {
  define: (env, input) => {
    input = input.trim();
    const pos = input.indexOf(" ");
    const val = valueParser(env, input.slice(pos));
    env[input.slice(0, pos)] = val[0];
    return [val[0], val[1].slice(1)];
  },
  quote: (env, input) => {
    input = input.trim();
    let quotedOutput;
    if ((quotedOutput = listParser(input))) {
      return quotedOutput;
    }
    const pos = bracketsParser(input);
    return [input.slice(0, pos), input.slice(pos + 1)];
  },
  if: (env, input) => {
    let condition;
    [condition, input] = valueParser(env, input);
    if (condition) {
      return valueParser(env, input);
    }
    input = input.trim().slice(bracketsParser(input.trim()));
    return valueParser(env, input);
  },
  lambda: (env, input) => {
    let paramsArr;
    [paramsArr, input] = listParser(input);
    input = input.trim();
    if (!paramsArr) return [null, "Error:Unequal Paratisis"];
    const defnition = input.slice(0, bracketsParser(input));

    let func = (params) => {
      const localEnv = Object.create(env);
      for (let i = 0; i < params.length; i++) {
        localEnv[paramsArr[i]] = params[i];
      }
      return [valueParser(localEnv, defnition)[0], paramsArr.length];
    };
    return [func, input.slice(defnition.length + 1)];
  },
  // "set!": (env, input) => {
  //   if (env[key]) {
  //     env[key] = valueParser(env, val)[0];
  //   } else return [null, "Error:Key Does not Exist"];
  // },
  // eval: (env, input) => {
  //   const val = valueParser(env, input)[0];
  //   return expressionEvaluator(env, val[0], val.slice(1));
  // },
};
// console.log(specialForms["quote"](env, "(this is the form of god)"));

export default specialForms;

// console.log(specialForms["eval"](env, ["+", "1", "2"]));
