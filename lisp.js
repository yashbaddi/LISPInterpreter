import valueParser from "./valueParser.js"
import operatorParser from "./opratorParser.js"

function lispint(input){
    let val=valueParser(input)||operatorParser(input)
    if(!val)return "Not a valid LISP Program";
    if(val[1]==='')return val[0]
}

console.log(lispint('(incf (+ 2 1) 3'))