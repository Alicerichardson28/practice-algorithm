var checkZeroOnes = function(s) {
    let countOnes = 0 
    let countZeros = 0 
    
    let longestOnes = 0 
    let longestZeros = 0

    
    for(let i = 0; i < s.length; i++){ 
        if(s[i] === "1"){  
            countOnes++ 
            longestZeros = Math.max(longestZeros, countZeros)
            countZeros = 0
        } else if(s[i] == "0"){ 
            countZeros++ 
            longestOnes = Math.max(longestOnes, countOnes) 
            countOnes = 0
        } 
    }
    longestZeros = Math.max(longestZeros, countZeros)
    longestOnes = Math.max(longestOnes, countOnes)
    
    return longestOnes > longestZeros
    
};