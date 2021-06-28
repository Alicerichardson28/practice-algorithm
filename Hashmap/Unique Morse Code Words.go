var uniqueMorseRepresentations = function(words) {
    const morseCode = {
       "a": ".-",
       "b": "-...",
       "c": "-.-.",
       "d": "-..",
       "e": ".",
       "f": "..-.",
       "g": "--.",
       "h": "....",
       "i": "..",
       "j": ".---",
       "k": "-.-",
       "l": ".-..",
       "m": "--",
       "n": "-.",
       "o": "---",
       "p": ".--.",
       "q": "--.-",
       "r": ".-.",
       "s": "...",
       "t": "-",
       "u": "..-",
       'v': "...-",
       "w": ".--",
       "x": "-..-",
       "y": "-.--",
       "z": "--.."
    }
    const translateToMorseCode = [];
    let tempWord = "";
    
    for(let i = 0; i < words.length; i++){
        for(let j = 0; j < words[i].length; j++){
            tempWord = tempWord + morseCode[words[i][j]]
        }
        translateToMorseCode.push(tempWord)
        tempWord = ""
    }
    
    const uniqueMorseCode = [...new Set(translateToMorseCode)]
    return uniqueMorseCode.length
   
};