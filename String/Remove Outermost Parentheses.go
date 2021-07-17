var removeOuterParentheses = function(s) {
    let stack = []
    let result = ""
    
    for(let i=0; i < s.length; i++){
        if(s[i] === "("){   //"(" 
            stack.push(s[i]) //stack =["(","("]
            if(stack.length > 1){  
                result = result.concat(s[i])  //result = "("
            }
        }else{ 
            stack.pop() //stack ["("]
            if(stack.length > 0){ 
                result = result.concat(s[i]) //result = "()"
            }
        }
    }
    return result
};