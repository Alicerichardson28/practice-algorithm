var shuffle = function(nums, n) {
    //create result = []
    //nums length devide by 2 set that to numsLengthDevineByTwo
    //looping the nums1 array until numsLengthDivineByTwo.length
        //put nums1 array in to result array
    //looping the nums array, set i = numsLengthDivineByTwo.length
        //put nums2 array in to odd position i + 1
    //return the result array  
//     if(nums.length > n * 2) return nums
//     let firstHalfNums = []
//     let secondHalfNums = []
//     const result = [];
//     for(let i = 0; i < n; i++){
//         firstHalfNums.push(nums[i])
//     }
    
//     for(let j = n; j < nums.length; j++){
//         secondHalfNums.push(nums[j])
//     }
    
//     for(let i = 0; i < nums.length / 2; i++){
//         let currentFirstHalf = firstHalfNums[i]
//         let currentSecoundHalf = secondHalfNums[i]
        
//         result.push(currentFirstHalf)
//         result.push(currentSecoundHalf)
//     }
   
//     return result
    
    
//     second solution using slice
    
    let firstHalfNums = nums.slice(0, n);
    let secondHalfNums = nums.slice(n);
    
    let result = [];
    
    for(let i = 0; i < nums.length / 2; i++){
        let currentFirstHalf = firstHalfNums[i]
        let currentSecoundHalf = secondHalfNums[i]
     
        result.push(currentFirstHalf)
        result.push(currentSecoundHalf)
    }
    
    return result
};