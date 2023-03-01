import valueParser from "./valueParser.js";
import * as readline from "node:readline";
import { globalEnv } from "./symbols.js";

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

function lispint(input) {
  let val = valueParser(input);
  if (!val) return "Not a valid LISP Program";
  if (val[1] !== "") lispint(val[1].trim());
  return val[0];
}
// lispint("(define twice (lambda (x) (* 2 x)))");
// lispint("(define repeat (lambda (f) (lambda (x) (f (f x)))))");
// console.log(lispint("((repeat twice) 10)"));
rl.setPrompt("LISP > ");
rl.prompt();
rl.on("line", (input) => {
  if (input == "exit()") rl.close();
  console.log(lispint(input));
  rl.prompt();
});
