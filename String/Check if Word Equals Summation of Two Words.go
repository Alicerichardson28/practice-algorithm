var isSumEqual = function(firstWord, secondWord, targetWord) {
    const letterValues = {
        a: 0,
        b: 1,
        c: 2,
        d: 3,
        e: 4,
        f: 5,
        g: 6,
        h: 7,
        i: 8,
        j: 9
    }

    let valuesOfFirstWord = ""
    let valuesOfSecondWord = ""
    let valuesOfTargetWord = ""
    
    for(let i = 0; i < firstWord.length; i++){
        if(letterValues.hasOwnProperty([firstWord[i]])){
           valuesOfFirstWord = valuesOfFirstWord + letterValues[firstWord[i]]
        }      
    }
    
     for(let i = 0; i < secondWord.length; i++){
        if(letterValues.hasOwnProperty([secondWord[i]])){
           valuesOfSecondWord = valuesOfSecondWord + letterValues[secondWord[i]]
        }      
    }
     for(let i = 0; i < targetWord.length; i++){
        if(letterValues.hasOwnProperty([targetWord[i]])){
           valuesOfTargetWord = valuesOfTargetWord + letterValues[targetWord[i]]
        }      
    }
   return parseInt(valuesOfFirstWord) + parseInt(valuesOfSecondWord) === parseInt(valuesOfTargetWord)
};