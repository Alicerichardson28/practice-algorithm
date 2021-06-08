var findOcurrences = function(text, first, second) { 		
		const result = [];
		const arrText = text.split(' ');
	 
	  
			for(let i = 0; i < arrText.length -2 ; i++){
				console.log(arrText.length - 2)
				if(arrText[i] === first && arrText[i+1] === second){
						result.push(arrText[i+2])
				}
		}
		return result
	};