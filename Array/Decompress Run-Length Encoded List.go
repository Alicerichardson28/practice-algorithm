var decompressRLElist = function(nums) {
    //crate result = []
    //looping through nums; nums = 1; i < nums.lenght; i+=2
        //set freq, val with nums[i+1] from nums[i]
        //use while loop freq > 0
            // push val to result
    //return result
    const result = []
    let freq, val
    for(let i = 1; i < nums.length; i+=2){ 
        [freq, val] = [nums[i-1], nums[i]]
        while(freq > 0){
            result.push(val)
            freq--
        }
    }
    return result
};