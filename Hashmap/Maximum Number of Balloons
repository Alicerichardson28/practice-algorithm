///////////////////////////////
// if S.length less than 7 return 0

// create wordBallon object
  // B : 1
  // A : 2
  // L : 0
  // O : 0
  // N : 0

// looping to S 
  // check if S[i] have any key in wordBallon object
  // if have increment 1

// create result = 0
// looping wordBallon object for check value
  // if value l and o divide by 2 
  // if value have same number return that number 

var maxNumberOfBalloons = function(text) {
    if(text.length < 7) return 0
    
    const wordBalloon = {
        b : 0,
        a : 0,
        l : 0,
        o : 0,
        n : 0
    }
    
    for(let i = 0; i < text.length; i++){
        if(wordBalloon.hasOwnProperty(text[i])){
            wordBalloon[text[i]] = wordBalloon[text[i]] + 1
        }
    }
  
    for(const key in wordBalloon){
        if(key === "l" || key === "o"){
            wordBalloon[key] = wordBalloon[key]/ 2
        }
    }
    
    
    const result = Object.values(wordBalloon)
    const finalResult = Math.min(...result)
    return Math.trunc(finalResult)
};

