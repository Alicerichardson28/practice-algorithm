//create variable assign to output array
//iterate through primeInput, set i = 2 
  //if primeInput[i] < 2 
      //return false
  //use modulus will give the reminder
  //if primeInput[i] % i === 0
  //return false
  //else 
  //push primeInput[i] to output 
//return output.sort()


const input = [10,7,41,18,31,2,4,5,11,19,25,17,13,1,0,15]
// const output = [2,5,7,11,13,17,19,31,41]
const filterPrime = (input) => {
  const output = [];
  for(let i = 0; i < input.length; i++){
    if(input[i] < 2){ //0,1
      continue
    }

    if(input[i] % 2 === 0){  //10,18,2,4
      if(input[i] === 2){ //2
        output.push(input[i])
      }
    } else { //7,41,31,5,11,19,25,17,13,15
      let flag = true
      for(let j = 2; j < input[i]; j++){ //7

        if(input[i] % j === 0){
          flag = false
          break
        } 
      }
      if(flag){
        output.push(input[i])
      }
    }
  }
  return output.sort((a,b) => a - b)
}

filterPrime(input)
