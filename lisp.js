import valueParser from "./valueParser.js"
import * as readline from 'node:readline';

const rl=readline.createInterface({
    input: process.stdin,
    output: process.stdout,
  });


export let symbolObj={
    '+':(arr)=>arr.reduce((acc,curr)=>acc+curr),
    '-':(arr)=>arr.reduce((acc,curr)=>acc-curr),
    '*':(arr)=>arr.reduce((acc,curr)=>acc*curr),
    '/':(arr)=>arr.reduce((acc,curr)=>curr/acc),
    'mod':(arr)=>arr.reduce((acc,curr)=>curr%acc),
    'incf':(arr)=>arr[0]+arr[1],
    'decf':(arr)=>arr[0]-arr[1],
    '>':(arr)=>arr[0]>arr[1],
    '>=':(arr)=>arr[0]>=arr[1],
    '<':(arr)=>arr[0]<arr[1],
    '<=':(arr)=>arr[0]<=arr[1],
    '=':(arr)=>arr[0]===arr[1]
};

function lispint(input){
    let val=valueParser(input)
    if(!val)return "Not a valid LISP Program";
    if(val[1]!=='')lispint(val[1].trim())
    return val[0]
}


// console.log(lispint('(define x (+ 2 3)) (if (= x 3) (define x 3))'))
// console.log(lispint('x '))

// console.log(symbolObj)


rl.setPrompt("LISP > ")
rl.prompt()
rl.on('line', (input) => {
    if(input=='exit()')rl.close()
    console.log(lispint(input));
    rl.prompt()
});



// rl.on('SIGCONT', () => {
//     // `prompt` will automatically resume the stream
    
//   });
  

// rl.on('SIGINT', () => {
//     rl.question('Are you sure you want to exit? ', (answer) => {
//       if (answer.match(/^y(es)?$/i)) rl.pause();
//     });
// });