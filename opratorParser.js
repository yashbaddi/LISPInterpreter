import valueParser from './valueParser.js'


export default function operatorParser(input){
    if(!input.startsWith("("))return null;
    input=input.slice(1).trim();
    return (_add(input)||
            _subtract(input)||
            _product(input)||
            _division(input)||
            _mod(input)||
            _incf(input)||
            _decf(input))
}

function _add(input){
    if(!input.startsWith("+"))return null;
    let rem=input.slice(1).trim()
    let val=valueParser(rem)
    rem=val[1].trim()
    let result=val[0]
    while(!rem.startsWith(')')){
        val=valueParser(rem)
        rem=val[1].trim()
        result=result+val[0]
    }
    return [result,rem.slice(1)]
}

function _subtract(input){
    if(!input.startsWith("-"))return null;
    let rem=input.slice(1).trim()
    let val=valueParser(rem)
    rem=val[1].trim()
    let result=val[0]
    while(!rem.startsWith(')')){
        val=valueParser(rem)
        rem=val[1].trim()
        result=result-val[0]
    }
    return [result,rem.slice(1)]
}

function _product(input){
    if(!input.startsWith("*"))return null;
    let rem=input.slice(1).trim()
    let val=valueParser(rem)
    rem=val[1].trim()
    let result=val[0]
    while(!rem.startsWith(')')){
        val=valueParser(rem)
        rem=val[1].trim()
        result=result*val[0]
    }
    return [result,rem.slice(1)]
}

function _division(input){
    if(!input.startsWith("/"))return null;
    let rem=input.slice(1).trim()
    let val=valueParser(rem)
    rem=val[1].trim()
    let result=val[0]
    while(!rem.startsWith(')')){
        val=valueParser(rem)
        rem=val[1].trim()
        result=val[0]/result
    }
    return [result,rem.slice(1)]
}

function _mod(input){
    if(!(input.startsWith("mod")||input.startsWith("rem")))return null;
    let rem=input.slice(3).trim()
    let val=valueParser(rem)
    rem=val[1].trim()
    let result=val[0]
    while(!rem.startsWith(')')){
        val=valueParser(rem)
        rem=val[1].trim()
        result=val[0]%result
    }
    return [result,rem.slice(1)]
}

function _incf(input){
    if(!(input.startsWith("incf")))return null;
    let rem=input.slice(4).trim()
    let val1=valueParser(rem)
    rem=val1[1].trim()
    let val2=valueParser(rem)
    rem=val2[1].trim()
    let result=val1[0]+val2[0]
    return [result,rem.slice(1)]
}

function _decf(input){
    if(!(input.startsWith("incf")))return null;
    let rem=input.slice(4).trim()
    let val1=valueParser(rem)
    rem=val1[1].trim()
    let val2=valueParser(rem)
    rem=val2[1].trim()
    let result=val1[0]-val2[0]
    return [result,rem.slice(1)]
}
