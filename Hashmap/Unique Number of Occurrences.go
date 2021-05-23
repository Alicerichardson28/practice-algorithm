var uniqueOccurrences = function(arr) {
    // create store = {}
    // initilize value = 0 to be able to increase
    // looping array and store in the hash map, key and value
        //  checking if the key is not in the store already adding to store {}
        //  if already in the store then + 1
    // looping store {} to get value
        //  check if value the number unique return true
        //  if not uniqe return flase 
    
    const store = {}
    let value = 1
    
    for(let i = 0; i < arr.length; i++){ //[1]
        // adding key and value
        if(!store.hasOwnProperty(arr[i])){
            store[arr[i]] = value
        } else {
            store[arr[i]] = store[arr[i]] + 1
        }  
    }
    
    const checkUniqueValue = {}
    for(const key in store){
        if(!checkUniqueValue.hasOwnProperty(store[key])){
            checkUniqueValue[store[key]] = checkUniqueValue[store[value]]
        } else {
            return false
        }
    }
    return true;
};
