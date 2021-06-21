var sortSentence = function(s) {
    const sArr = s.split(' ');
    const result = new Array(sArr.length)
  
    for(let i = 0; i < sArr.length; i++){
        const lastIndex = sArr[i].length - 1
        const lastChar = sArr[i][lastIndex]
        const newWord = sArr[i].slice(0, lastIndex)
     
        result[parseInt(lastChar)-1] = newWord
    }
    return result.join(' ');
};
    