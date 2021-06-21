var sumOddLengthSubarrays = function(arr) {  
    if(arr.length === 1)return arr
    let subArray = []
    for(let i=0; i < arr.length; i++){
        for(let j = i; j < arr.length; j++){
            subArray.push(arr.slice(i, j+1))
        }
    }
    let sum = 0
        for(let i = 0; i < subArray.length; i++){
            let individualArr = subArray[i]
            if(individualArr.length%2 == 1){
                for(let j=0; j < individualArr.length; j++){
                   sum = sum + individualArr[j]
                    console.log(sum)
                }
            }
        }
    return sum
};