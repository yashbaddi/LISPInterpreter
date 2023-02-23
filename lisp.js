import valueParser from "./valueParser.js"
import operatorParser from "./opratorParser.js"

function lispint(input){
    let val=valueParser(input)
    if(!val)return "Not a valid LISP Program";
    
    if(val[1]==='')return val[0]
}

function mulitlineLISP(input){


}

console.log(lispint('(decf (+ 2 3) 3)'))