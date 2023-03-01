import expressionParser, { parametersParser } from "./expressionParser.js";
import { globalEnv } from "./symbols.js";
import valueParser from "./valueParser.js";

const env = globalEnv;

let specialForms = {
  define: (key, val, env) => {
    env[key] = valueParser(val)[0];
    return 1;
  },
  quote: (input, env) => {
    return input;
  },
  if: (condition, statement1, statement2, env) => {
    if (valueParser(condition, env)[0]) return valueParser(statement1, env)[0];
    else return valueParser(statement2, env)[0];
    return 1;
  },
  lambda: (localparams, defnition, env) => {
    // let keyParameters = [];
    // let i = 1;
    // while (input[i] !== ")") {
    //   let pos = input.indexOf(" ");
    //   keyParameters.push(input.slice(i, pos).trim());
    //   i = i + pos;
    // }
    let paramsArr = parametersParser(localparams.slice(1));

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
    }
  },
};

export default specialForms;
