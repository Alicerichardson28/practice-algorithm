// Get 2 input string and array of integer
// create new empty array []
// create i = 0
// create j = 0
// create while loop if i less than s length and. j less than indices length
// set result indices j to s i 
// increment i by one and j by one
// return result using . join

var restoreString = function(s, indices) {
    const result = [];
    let i = 0;
    let j = 0;
    
    while(i < s.length && j < indices.length){
        result[indices[j]] = s[i]
        i++
        j++
    }
   return result.join('')
};