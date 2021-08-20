var areOccurrencesEqual = function(s) {
    // create empty result object {}
    // iterate through s then store into result object 
        //if already have it increment
        //if don't have before store as a key
    //then check if all the keys value have same value or not
        //yes return true
        //no return false
    
    const result = {}
    for(let i = 0; i < s.length; i++){
        if(!result.hasOwnProperty(s[i])){
            result[s[i]] = 1
        } else {
            result[s[i]]++
        }
    }
    
    let valueResult = Object.values(result)
   
 
    for(let i = 0; i < valueResult.length -1; i++){
        console.log(valueResult[i] !== valueResult[i+1])
        
        if(valueResult[i] !== valueResult[i+1]){
            return false
        }
    }
   return true
    
    
 //     let valueResult = Object.values(result)
//     const target = valueResult[0]
   
 
//     for(let i = 0; i < valueResult.length -1; i++){
//         console.log(valueResult[i] !== valueResult[i+1])
        
//         if(valueResult[i] !== valueResult[i+1]){
//             return false
//         }
//     }
//    return true
    
    
    // const target = result[s[0]]
    // for(let key in result){
    //     if(result[key] !== target){
    //         return false
    //     }
    // }
    // return true
};