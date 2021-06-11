var subtractProductAndSum = function(n) {
	let nArr = Array.from(n.toString())
	
	let resultMultiple = 1
	let resultPlus = 0
	
	for(let i = 0; i < nArr.length; i++){
	   resultMultiple = resultMultiple * parseInt(nArr[i])
	}
	  
	for(let i = 0; i < nArr.length; i++){
	   resultPlus = resultPlus + parseInt(nArr[i])
	}
	  
	  return resultMultiple - resultPlus 
	
  };