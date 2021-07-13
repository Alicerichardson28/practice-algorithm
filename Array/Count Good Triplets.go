var countGoodTriplets = function(arr, a, b, c) {    
    let count=0
    
    for(let i = 0; i < arr.length-2; i++){ //3
        for(let j = i+1; j < arr.length; j++){ //0, 1, 1
            for(let k =j+1; k < arr.length; k++){ //1, 1, 9
                if (Math.abs(arr[i] - arr[j]) <= a && 
                    Math.abs(arr[j] - arr[k]) <= b && 
                    Math.abs(arr[i] - arr[k]) <= c ){
                   count++
                }
            }
        }
    }
   return count         
};