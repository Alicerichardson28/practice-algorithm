var longestCommonPrefix = function(strs) {
    //create prefix = ''
        //looping through strs at index[0]
            //set index[0][i] assign to character
            //itrate through strs[j] by using i for pointing to that index
                //compair if character !== strs[j][i]
                    //return prefix
                //else
                    //prefix = prefix concat character
    //return prefix
    
    let prefix = '';
        for(let i = 0; i < strs[0].length; i++){
            const character = strs[0][i]
        
            for(let j = 0; j < strs.length; j++){
                //when you loop, you using i to refer to that index // j itrate to strs but i use for point to that index
                // console.log(strs[j][i])
                if(character !== strs[j][i]){
                    return prefix
                }
            } 
            prefix = prefix + character  
        }
    return prefix 
};