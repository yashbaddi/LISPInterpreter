import expressionParser, { listParser } from "./expressionParser.js";
import { globalEnv } from "./symbols.js";

export default function valueParser(env = globalEnv, input) {
  input = whiteSpaceParser(input);
  return (
    expressionParser(env, input) ||
    numberParser(input) ||
    symbolParser(env, input) ||
    booleanParser(input) ||
    StringParser(input)
  );
}

//Symbol Parser
export function symbolParser(env, input) {
  // const symbol = input.slice(0, input.indexOf(" "));
  input = whiteSpaceParser(input);
  let i = 0;
  while (!(input[i] === " " || input[i] === ")" || input[i] === undefined)) {
    i++;
  }
  const symbol = input.slice(0, i);
  if (env[symbol] == undefined) return null;
  return [env[symbol], input.slice(symbol.length)];
}

//Number Parser
function numberParser(input) {
  let re = /^-?([1-9](\d)*|0)(\.(\d)+)?([eE][+-]?(\d)+)?/;
  let n = input.match(re);
  if (!n) return null;
  return [Number(n[0]), input.slice(n[0].length)];
}

//Boolean Parser
function booleanParser(input) {
  if (input.startsWith("#t")) return [true, input.slice(2)];
  if (input.startsWith("#f")) return [false, input.slice(2)];
  return null;
}

//String Parser
function StringParser(input) {
  input = whiteSpaceParser(input);
  let a = new Set(['"', "\\", "/", "b", "f", "n", "r", "t", "u"]);
  if (!input.startsWith('"')) return [null, " Error: Unknown Identifier"];
  let i = 1;
  while (input[i] != '"') {
    if (input.charCodeAt(i) === 9 || input.charCodeAt(i) === 10) return null;
    if (input[i] == "\\") {
      if (!a.has(input[i + 1])) return null;
      i++;
    }
    i++;
  }
  return [input.slice(1, i), input.slice(i + 1)];
}

//WhiteSpace Parser
export function whiteSpaceParser(input) {
  let i = 0;
  while (/^\s/.test(input[i])) {
    i++;
  }
  return input.slice(i);
}
