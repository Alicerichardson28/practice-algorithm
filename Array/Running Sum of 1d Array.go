var runningSum = function(nums) { 
    
    //  if(nums.length === 1) return nums
    // for(let i = 1; i < nums.length; i++){
    //     nums[i] = nums[i] + nums[i-1]
    // }
    // return nums
    
//      if(nums.length === 1) return nums
//      let result = []
//      let temp = 0
     
//      for(let number of nums){
//          temp = number + temp
//          result.push(temp)
//      }
//   return result
    
    // let result = []
    // nums.reduce((acc, num) => {
    //     result.push(acc + num)
    //     return acc+num
    // }, 0)
    // return result
    
    nums.reduce((acc, num, index, array) => {
        array[index] = acc + num
        return acc+num
    }, 0)
    return nums
};