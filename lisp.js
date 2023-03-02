import valueParser from "./valueParser.js";
import * as readline from "node:readline";
import { globalEnv } from "./symbols.js";
import { linkSync } from "node:fs";

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

export default function lispint(input) {
  let val = valueParser(input);
  if (!val) return "Not a valid LISP Program";
  if (val[1] !== "") lispint(val[1].trim());
  return val[0];
}

rl.setPrompt("LISP > ");
rl.prompt();
rl.on("line", (input) => {
  if (input == "exit()") rl.close();
  console.log(lispint(input));
  rl.prompt();
});
