import { parametersParser } from "./expressionParser.js";
import { globalEnv } from "./symbols.js";
import valueParser from "./valueParser.js";

const env = globalEnv;

let specialForms = {
  define: (key, val, env) => {
    env[key] = valueParser(val)[0];
  },
  quote: (input, env) => {
    return input[0];
  },
  if: (condition, statement1, statement2, env) => {
    if (valueParser(condition, env)[0]) return valueParser(statement1, env)[0];

    return valueParser(statement2, env)[0];
  },
  lambda: (localparams, defnition, env) => {
    let paramsArr = parametersParser(localparams);

    let func = (params) => {
      let localEnv = Object.create(env);
      for (let i = 0; i < params.length; i++) {
        localEnv[paramsArr[i]] = params[i];
      }
      return valueParser(defnition, localEnv)[0];
    };
    return func;
  },
  "set!": (key, val, env) => {
    if (env[key]) {
      env[key] = valueParser(val)[0];
    } else throw new Error(key + " is not defined");
  },
};

export default specialForms;
