var replaceDigits = function(s) {
    // input string
    // if length === 1 return string    
    // using let storeArr = s.split("") //['a', '1', 'c', '1', 'e', '1']
    // looping through storeArr
        //if found the letter store that letter in the newArr         
    // change element in newArr1 from number to char 
    // using join("")
    // return string  
    

    // for (let i = 0; i < s.length; i++) {
    //   console.log(s.charCodeAt(i));
    // }
    if(s.length === 1) return s
    let storeS = s.split('');

    for(let i = 1; i < storeS.length; i+=2){
        // console.log("index", i, "value",storeS[i])
        // storeS[i] = String.fromCharCode(storeS[i-1].charCodeAt() + parseInt(storeS[i])); 
        // +storeS[i]  type coercion      
        storeS[i] = String.fromCharCode(storeS[i-1].charCodeAt() +  +storeS[i]); 
        // console.log('line 43',storeS[i])
    }
    return storeS.join('')
};