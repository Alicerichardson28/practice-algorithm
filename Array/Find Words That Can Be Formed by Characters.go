var countCharacters = function(words, chars) { 
    let output = 0
    let charArr = new Array(26).fill(0);
    
    for(let i = 0; i < chars.length; i++){
        charArr[chars[i].charCodeAt(0) - "a".charCodeAt(0)]++
    }
    for(let i = 0; i < words.length; i++){
        let tempArr = charArr.slice();
        let charCount = 0
        for(let j = 0; j < words[i].length; j++){
            let curr = words[i][j].charCodeAt(0) - "a".charCodeAt(0);
            if(tempArr[curr] > 0){
                charCount++
                tempArr[curr]--
            }
        }
        if(words[i].length === charCount){
        output += charCount;
        }
    
    }
    return output
};