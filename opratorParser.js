import valueParser from './valueParser.js'

const resolvedObj={
    '+':(arr)=>arr.reduce((acc,curr)=>acc+curr),
    '-':(arr)=>arr.reduce((acc,curr)=>acc-curr),
    '*':(arr)=>arr.reduce((acc,curr)=>acc*curr),
    '/':(arr)=>arr.reduce((acc,curr)=>curr/acc),
    'mod':(arr)=>arr.reduce((acc,curr)=>curr%acc),
    'incf':(arr)=>arr[0]+arr[1],
    'decf':(arr)=>arr[0]-arr[1]
}

export default function operatorParser(input){
    if(!input.startsWith("("))return null;
    let opertator=input.slice(1,input.indexOf(" "));
    let rem=input.slice(opertator.length+1).trim();
    let arr=[]
    while(!rem.startsWith(")")){
        let val=valueParser(rem)
        if(!val)return null
        rem=val[1].trim()
        arr.push(val[0])
    }
    if(rem[0]!==')')return null
    return [resolvedObj[opertator](arr),rem.slice(1)]
}
