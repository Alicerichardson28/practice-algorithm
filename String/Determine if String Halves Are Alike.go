// s = string.length / 2     
    //firstHalf = slice string.length / 2 to half index
    //sencondHalf = slice secondhalf index    
    
    //create object vowelsFirstHalf = {'a': 0}
    //create object vowelsSencondHalf = {'a': 0}
    
        //looping firstHalf 
            //if have key increment value + 1                       
        //looping secondHalf
            //if have key increment value + 1 
    
    // looping vowelsFirstHalf value    
        // looping vowelsSencondHalf value
        //  compare value object
        //  if == return true
        //  if not == retrun false  

		var halvesAreAlike = function(s) {       
			//split half of string    
			let sHalf = s.length / 2 
			const firstHalf = s.slice(0, sHalf) 
			const secondHalf = s.slice(sHalf)
		
			
			//create object vowelsFirstHalf
			const vowelsFirstHalf = {
				'a': 0,
				'e': 0,
				'i': 0,
				'o': 0,
				'u': 0,
				'A': 0,
				'E': 0,
				'I': 0,
				'O': 0,
				'U': 0,
			}
			
			 //create object vowelsSecondHalf
			  const vowelsSecondHalf = {
				'a': 0,
				'e': 0,
				'i': 0,
				'o': 0,
				'u': 0,
				'A': 0,
				'E': 0,
				'I': 0,
				'O': 0,
				'U': 0,
			}
			   
			let sumFirstKey = 0
			let sumSecondKey = 0
			for(let i = 0; i < firstHalf.length; i++){
				 if(vowelsFirstHalf[firstHalf[i]] === 0){
					sumFirstKey++
				 }
			}
				
			for(let i = 0; i < secondHalf.length; i++){
				if(vowelsFirstHalf[secondHalf[i]] === 0){
					sumSecondKey++
				 }
			}
			return sumFirstKey === sumSecondKey  
		};
		
		