var smallerNumbersThanCurrent = function(nums) {
    // if array length less than 2 return array    
    // input array of nums 
    // create smallNums {} 
    // first loop through array start from index 0
        //second loop through array start from index 1
        // check if first loop > second loop
        // if don't have this key in the object
        // if key have a same name store as a new key       
        // store first loop index as a key and set value = 1
        // else 
        // increment value + 1 
    // create result array    
    // looping object to get the value
    // when get the value store in array // can use built in method Object.value
    // return output result array of nums   
    
//     if(nums < 2) return nums
//     let smallNums = {}
    
//     for(let i = 0; i < nums.length; i++){ //8, 1, 2
//         // console.log('looping i',nums[i])
//         for(let j = nums.length - 1; j >= 0; j--){ 
//            if(nums[i] > nums[j] && !smallNums.hasOwnProperty(nums[i])){
//                 smallNums[nums[i]] = 1
//             } else if(nums[i] < nums[j] && !smallNums.hasOwnProperty(nums[i])){
//                 smallNums[nums[i]] = 0
//             } else if(nums[i] > nums[j] && smallNums.hasOwnProperty(nums[i])){
//                 smallNums[nums[i]]++
//         }
//     }
//     console.log(smallNums)
    
//     } 
    
    let resultArr = []
    let count = 0
    
    for(let i = 0; i < nums.length; i++){
        for(let j = 0; j < nums.length; j++)
        // for(let j = nums.length - 1; j >= 0; j--)
            if(nums[j] < nums[i]){
                count++
            }
        // console.log(count)
        resultArr.push(count)
        count = 0
    }
    return resultArr
};

     