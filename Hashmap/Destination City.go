var destCity = function(paths) {
    //create empty object and assign to departured
        //iterate through paths set i = 0
            //check if paths[i][0] doesn't have key in departured object
                //push it to departured object as a key
        //iterate through paths set i = 0
            //check if paths[i][1] don't have any match key in the departured object 
                //return paths[i][1]
    const departured = {};
    
    for(let i = 0; i < paths.length; i++){
        departured[paths[i][0]] = true
    }
    
    for(let i = 0; i < paths.length; i++){
        if(!departured[paths[i][1]]){
            return paths[i][1]
        }
    }
};
