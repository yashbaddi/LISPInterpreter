import opratorParser from './opratorParser.js'



export default function valueParser(input){
    input=input.trim()
    return (booleanParser(input)||
            numberParser(input)||
            StringParser(input)||
            opratorParser(input))

}

//Boolean Parser
function booleanParser(input){
    if (input.startsWith('#t')) return [true,input.slice(2)]
    if (input.startsWith('#f')) return [false,input.slice(2)]
    return null;
}

//Number Parser
function numberParser(input){
    let re=/^-?([1-9](\d)*|0)(\.(\d)+)?([eE][+-]?(\d)+)?/
    let n=input.match(re);
    if(!n)return null
    return [Number(n[0]),input.slice(n[0].length)]
}
//String Parser
function StringParser(input){
    input=input.trim()
    let a = new Set(['"', '\\', '/', 'b', 'f', 'n', 'r', 't', 'u']) 
    if(!(input.startsWith('"'))) return null;
    let i=1;
    while(input[i]!='"'){
        if(input.charCodeAt(i)===9||input.charCodeAt(i)===10) return null
        if(input[i]=='\\'){
            if(!a.has(input[i+1]))return null;
            i++;
        }
        i++
    }
    return [input.slice(1,i),input.slice(i+1)];
}