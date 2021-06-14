var interpret = function(command) {
    if (command.length === 1) return command
    
    const result = []
    
    for(let i = 0; i < command.length; i++){
        if(command[i] === "("){
                if(command[i + 1] === ")"){
                    result.push("o")
                    i+=1
                } else {
                    result.push("al")
                    i+=3
                }
        } else {
            result.push("G")
        }
    }
    return result.join('')
};