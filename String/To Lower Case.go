var toLowerCase = function(s) {
    if(s === s.lowercase) return s
    let result = ""
    for(let i = 0; i < s.length; i++){
        if(s.charCodeAt(i) >= 65 && s.charCodeAt(i) <= 90){
            const convertToAscii = s.charCodeAt(i)+32
            const covertToChar = String.fromCharCode(convertToAscii)
            result = result+covertToChar
        }else{
            result = result + s[i]
        }
    }
    return result
};