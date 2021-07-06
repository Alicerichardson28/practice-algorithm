var isPalindrome = function(s) {
    //using ascii code to check charecters and number 
        //push charecters and change to lowercase and number into new array assign new array name newSArray
    //2 pointer 
    //iterate through newSArray set i = 0 and j = newSArray.lenght-1; i < j; i++; j--
        //if newSArray[i] !== newSArray[j]
            //return false
    //return true
 
    
    const newSArray = []
    for(let i = 0; i < s.length; i++){
        if(s[i].charCodeAt(0) >= 97 && s[i].charCodeAt(0) <= 122){
            newSArray.push(s[i].toLowerCase())
        } else if(s[i].charCodeAt(0) >= 48 && s[i].charCodeAt(0) <= 57){
            newSArray.push(s[i])
        } else if(s[i].charCodeAt(0) >= 65 && s[i].charCodeAt(0) <= 90) {
            newSArray.push(s[i].toLowerCase())
        }
    }
    

    for(let i = 0, j = newSArray.length - 1; i < j; i++, j--){
        if(newSArray[i] !== newSArray[j]){
            return false
        }
    }
    return true
};
