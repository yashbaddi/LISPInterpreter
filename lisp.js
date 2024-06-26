import valueParser from "./valueParser.js";
import * as readline from "node:readline";
import { globalEnv } from "./symbols.js";

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

export default function lispint(input) {
  let val = valueParser(globalEnv, input);
  if (val[0] === null) return val[1];
  if (val[1] !== "") console.log(lispint(val[1].trim()));

  return val[0];
}

// console.log(lispint("(define 2x 2)"));
console.log(lispint("(define x 2"));

rl.setPrompt("LISP > ");
rl.prompt();
rl.on("line", (input) => {
  console.log(lispint(input));
  rl.prompt();
});

rl.on("SIGINT", () => rl.pause());
