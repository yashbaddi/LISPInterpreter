import { listParser, expressionEvaluator } from "./expressionParser.js";
import { globalEnv } from "./symbols.js";
import valueParser from "./valueParser.js";

const env = globalEnv;

let specialForms = {
  define: (env, key, val) => {
    env[key] = valueParser(env, val)[0];
    return env[key];
  },
  quote: (env, input) => {
    const quotedOutput = listParser(input) || input;
    return quotedOutput;
  },
  if: (env, condition, statement1, statement2) => {
    if (valueParser(env, condition)[0]) return valueParser(env, statement1)[0];
    return valueParser(env, statement2)[0];
  },
  lambda: (env, localparams, defnition) => {
    let paramsArr = listParser(localparams);
    if (paramsArr) return [null, "Error:Unequal Paratisis"];

    let func = (params) => {
      const localEnv = Object.create(env);
      for (let i = 0; i < params.length; i++) {
        localEnv[paramsArr[i]] = params[i];
      }
      return [valueParser(localEnv, defnition)[0], paramsArr.length];
    };
    return func;
  },
  "set!": (env, key, val) => {
    if (env[key]) {
      env[key] = valueParser(env, val)[0];
    } else return [null, "Error:Key Does not Exist"];
  },
  // eval: (env, input) => {
  //   const val = valueParser(env, input)[0];
  //   return expressionEvaluator(env, val[0], val.slice(1));
  // },
};

export default specialForms;

// console.log(specialForms["eval"](env, ["+", "1", "2"]));
