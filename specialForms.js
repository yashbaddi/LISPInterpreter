import opratorParser from './opratorParser.js'
import { symbolObj } from './lisp.js'
import valueParser from './valueParser.js';

let specialForms={
    'define':(input)=>{
        let pos=input.indexOf(" ")
        let val=valueParser(input.slice(pos))
        symbolObj[input.slice(0,pos)]=val[0]
        return [0,val[1].slice(1)]
    },
    'quote':(input)=>{
        let i=nextOperand(input)
        return [input.slice(0,i),input.slice(i+1)]
    },
    'if':(input)=>{
        let condn=valueParser(input)
        
        if(condn[0]){
            return valueParser(condn[1])
        }else{
            return[0,condn[1].slice(nextOperand(input))]
        }
    },
    'lambda':(input)=>{

    },
    'set!':(input)=>{

    }
    
};
function nextOperand(input){
    let parCheck=[];
    parCheck.push('(')
    let i=0
    do{
        if(input[i]==='(')parCheck.push('(')
        if(input[i]===')')parCheck.pop();
        i++
    }while(parCheck.length!==0)
    return i;
}

export default specialForms;
