var repeatedNTimes = function(nums) {
   
    const result ={};
    for(let i = 0; i < nums.length; i++){
        if(!result.hasOwnProperty(nums[i])){
            result[nums[i]] = 1
        }else{
            result[nums[i]]++
        }
    }

   
    let nResult = ''
    let max = 0
    
    for(const key in result){
        if( max < result[key]){
            max = result[key]
            nResult = key
        }
    }
    return parseInt(nResult)
};