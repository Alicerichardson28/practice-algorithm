var singleNumber = function(nums) {
    //looping number and store in object 
    //looping to obeject and find the value == 1 
    
	let result;
    for(let i = 0; i < nums.length; i++){
        result = result ^ nums[i];
    }
    return result
    // const numsObject = {}
    // for(let i = 0; i < nums.length; i++){
    //     if(!numsObject.hasOwnProperty(nums[i])){
    //         numsObject[nums[i]] = 1
    //     } else {
    //         numsObject[nums[i]]++
    //     }
    // }
    // for(const key in numsObject){
    //     if(numsObject[key] == 1) return key
    // }
};