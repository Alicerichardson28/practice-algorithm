    //create stringX = ""
    //using toString change number into string
    //iterate to number in the stringfrom back 
        //store value into stringX
    
    //change stringX to number
    //compair between stringX and x 
    // if the value match return true
    // if doesn't match return false


var isPalindrome = function(x) { 
    const stringX = x.toString();
    let result = ""
    
    for(let i = stringX.length-1; i >= 0; i--){
        result = result + stringX[i]
    }
    return parseInt(result) === x
};
