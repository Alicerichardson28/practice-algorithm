var countConsistentStrings = function(allowed, words) {
    //create allowedChar in the object
        //iterate through allowed and put allowed[i] as a key into allowedChar object
    //create variable name countConsistent = 0
    //iterate through words set i = 0;  words length - 1
        //if words[i][i] have the same key of allowedChar && words[i][i].length === 0
        //else 
            //skip that words[i]
    //return countConsistent
    
    const allowedChar = {};
    let countConsistent = 0;
    
    for(let i = 0; i < allowed.length; i++){
        if(!allowedChar.hasOwnProperty(allowed[i])){
            allowedChar[allowed[i]] = 0
        }
    }
    
    //{a: 'a', b: 'b'}
    
    for(let i = 0; i < words.length; i++){ //ad
        for(let j = 0; j < words[i].length; j++){ //a
            // if(allowedChar.hasOwnProperty(words[i][j])){ //boolean
            //     // if(words[i][j] === allowedChar[words[i][j]] && j === words[i].length-1)
            //     if(j === words[i].length-1)
            //         countConsistent++
            // } else {
            //     break
            // }
            // if(!allowedChar.hasOwnProperty(words[i][j])){
            //     break
            // } else if(j === words[i].length-1) {
            //     countConsistent++
            // }
            
            //return value include truthy or falsy
            //make sure value not 0 or false 
            // console.log(allowedChar[words[i][j]])
            if(allowedChar[words[i][j]] === 0){
                 if(j === words[i].length-1){
                     countConsistent++
                 }
            } else {
                 break
            }   
        }
      }
    return countConsistent
};