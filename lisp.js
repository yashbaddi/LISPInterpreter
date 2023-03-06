import valueParser from "./valueParser.js";
import * as readline from "node:readline";
import { globalEnv } from "./symbols.js";

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

export default function lispint(input) {
  let val = valueParser(globalEnv, input);
  if (!val) return "Not a valid LISP Program";
  if (val[1] !== "") lispint(val[1].trim());
  return val[0];
}

rl.setPrompt("LISP > ");
rl.prompt();
rl.on("line", (input) => {
  console.log(lispint(input));
  rl.prompt();
});

rl.on("SIGINT", () => rl.pause());
