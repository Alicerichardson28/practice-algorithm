var isValid = function(s) {
    const allBrackets = {
        '(': ')',
        '[': ']',
        '{': '}'
    }
    
    let stack = []
    const allOpenBracket = ['(', '[', '{']
    for(let i = 0; i < s.length; i++){
        if(allOpenBracket.includes(s[i])){
            stack.push(s[i])
        } else {
            let lastElementInStack = stack.pop()
            if(allBrackets[lastElementInStack] !== s[i]) return false
        }
    }
    return stack.length ? false : true
};