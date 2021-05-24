var numJewelsInStones = function(jewels, stones) {
    // accept 2 input jewels and stones
    // create empthy object {} for keep tracking key 
    // lowercase and Uppercase matter 
    // looping to jewels and store as a key in object
    // looping to stones and check if match the key or not if yes increment by 1 if not out of the loop
    // loop value in the object and sum
    // return sum   
    
    const jewelsObject = {};
    for(let i = 0; i < jewels.length; i++){
        console.log('looping string', jewels[i]);
        if(!jewelsObject.hasOwnProperty(jewels[i])){
            jewelsObject[jewels[i]] = 0
        } 
    }
    
    for(let i = 0; i < stones.length; i++){
        for(let key in jewelsObject){
                if(stones[i] === key){
                    jewelsObject[key] = jewelsObject[key] + 1
            }
        }
    }
    let result = 0
    for(let key in jewelsObject){
        result = jewelsObject[key] + result
    }
    
    return result;
};