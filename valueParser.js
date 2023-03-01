import expressionParser, { parametersParser } from "./expressionParser.js";
import { globalEnv } from "./symbols.js";

export default function valueParser(input, env = globalEnv) {
  input = input.trimStart();
  return (
    expressionParser(input, env) ||
    symbolParser(input, env) ||
    booleanParser(input) ||
    numberParser(input) ||
    StringParser(input)
  );
}

//Boolean Parser
function booleanParser(input) {
  if (input.startsWith("#t")) return [true, input.slice(2)];
  if (input.startsWith("#f")) return [false, input.slice(2)];
  return null;
}

//Number Parser
function numberParser(input) {
  let re = /^-?([1-9](\d)*|0)(\.(\d)+)?([eE][+-]?(\d)+)?/;
  let n = input.match(re);
  if (!n) return null;
  return [Number(n[0]), input.slice(n[0].length)];
}
//String Parser
function StringParser(input) {
  input = input.trim();
  let a = new Set(['"', "\\", "/", "b", "f", "n", "r", "t", "u"]);
  if (!input.startsWith('"')) return null;
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

function symbolParser(input, env) {
  if (!env[input]) return null;
  return [env[input], input.slice(input.length)];
}
