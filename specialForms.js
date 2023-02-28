import expressionParser from "./expressionParser.js";
import { symbolObj } from "./symbols.js";
import valueParser from "./valueParser.js";

let specialForms = {
  define: (key, val) => {
    symbolObj[key] = valueParser(val)[0];
    return 1;
  },
  quote: (input) => {
    return input;
  },
  if: (condition, statement1, statement2) => {
    if (valueParser(condition)) valueParser(statement1);
    else valueParser(statement2);
    return 1;
  },
  lambda: (params, defnition) => {
    // let keyParameters = [];
    // let i = 1;
    // while (input[i] !== ")") {
    //   let pos = input.indexOf(" ");
    //   keyParameters.push(input.slice(i, pos).trim());
    //   i = i + pos;
    // }
    // let func = (params) => {
    //   Object.assign(
    //     ...keyParameters.map((element, index) => ({ [element]: params[index] }))
    //   );
    //   valueParser(input);
    //   Object.keys(params).forEach((ele) => delete symbolObj[ele]);
    // };
    // return [func, input.slice(nextOperand(input))];
  },
  "set!": (input) => {},
};

export default specialForms;
