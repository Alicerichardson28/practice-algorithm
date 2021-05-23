const numIdenticalPairs = (nums) => {
    // initilize goodPair = {}    
    // input array [nums]
    // looping array and put in hashmap  
        // if goodPair not exits 
        // adding in key and value
        // if already exits add 1    
    // return the value in hasmap with sum
    
    const goodPair = {};
    for(let i = 0; i < nums.length; i++){
        for(let j = i + 1; j < nums.length; j++){
           if(nums[i] === nums[j] && !goodPair.hasOwnProperty(nums[i])){
                goodPair[nums[i]] = 1
           } else if(nums[i] === nums[j] && goodPair.hasOwnProperty(nums[i])) {
               goodPair[nums[i]] = goodPair[nums[i]] + 1
           }
        }
    }
    
    let sum = 0
    for(const key in goodPair){
        sum = (goodPair[key]) + sum
    }
    return sum;
};
