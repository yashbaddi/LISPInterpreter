import valueParser from './valueParser.js'
import specialForms from './specialForms.js';
import { symbolObj } from './lisp.js';

export default function operatorParser(input){
    if(!input.startsWith("("))return null;
    let opertator=input.slice(1,input.indexOf(" "));
    let rem=input.slice(opertator.length+1).trim();
    let arr=[]
    if(specialForms[opertator]){
        return specialForms[opertator](rem);
    } 
    else if(symbolObj[opertator]){
        while(!rem.startsWith(")")){
            let val=valueParser(rem)
            if(!val)return null
            rem=val[1].trim()
            arr.push(val[0])
        }
    } 
    else{
        return null
    }
    if(rem[0]!==')')return null
    return [symbolObj[opertator](arr),rem.slice(1)]
    
}
