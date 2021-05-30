// create count variable = 0
  // condition rule key == "type"
    // looping through items
        //if ruleValue is equal to the index 0 of each item
        //increment count by one
  // condition rule key == "color"
    // looping through items
        //if ruleValue is equal to the index 1 of each item
        //increment count by one
  // condition rule key == "name"
    // looping through items
        //if ruleValue is equal to the index 2 of each item
        //increment count by one
 // return count 
 var countMatches = function(items, ruleKey, ruleValue) {
    let count= 0
    if(ruleKey === "type"){
        for(let i = 0; i < items.length; i++){
            if(items[i][0] === ruleValue){
                count++
            }
        }
    } else if (ruleKey === "color"){
        for(let i = 0; i < items.length; i++){
            if(items[i][1] === ruleValue){
                count++
            }
        }
    } else if(ruleKey === "name"){
        for(let i = 0; i < items.length; i++){
            if(items[i][2] === ruleValue){
                count++
            }
        }
    }
    return count
};