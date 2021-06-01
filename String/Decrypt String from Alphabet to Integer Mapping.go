var freqAlphabets = function(s) {
    //create object of a-z and set value = number{a: 1, b: 2}
    //create outputChar = ""
    //looping from the back of string and check 
        //if found # and two digit number 
        //check in the object a - z value and concat that key into outputChar
        //else just check that number match to value object a-z or not 
        //concat that key into outputChar
    //return outputChar
    const aTozChar = {
        a : 1,
        b : 2,
        c : 3,
        d : 4,
        e : 5,
        f : 6,
        g : 7,
        h : 8,
        i : 9,
        j : 10,
        k : 11,
        l : 12,
        m : 13,
        n : 14,
        o : 15,
        p : 16,
        q : 17,
        r : 18,
        s : 19,
        t : 20,
        u : 21,
        v : 22,
        w : 23,
        x : 24,
        y : 25,
        z : 26  
    }
    
  
    let newArr = []
    for(let i = 0; i < s.length; i++){
        let temp = ""
        if(s[i+2] === '#'){
          newArr.push(temp.concat(s[i],s[i+1]))
          i+=2
        } else {
            newArr.push(s[i])
        }
    } 

    
    let newChar = []
    for(let i = 0; i < newArr.length; i++){
        for(let key in aTozChar){
            if(parseInt(newArr[i]) === aTozChar[key]){
                newChar.push(key)
            }
        }
    }
    return newChar.join('')
};