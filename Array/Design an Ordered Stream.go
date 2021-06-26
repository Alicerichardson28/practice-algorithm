var OrderedStream = function(n) {
    //how to make space array with contructor
    this.resultArr = new Array(n + 1)
    //for keep tracking pointer
    this.pointer = 1
};

/** 
 * @param {number} idKey 
 * @param {string} value
 * @return {string[]}
 */
OrderedStream.prototype.insert = function(idKey, value) {
    //declare output = []
    //insert value to this.resultArr at index idKey
    //while value from resultArr at pointer index !== undefined
        //push value from that pointer into output 
        //and then increase pointer by one
    //return output
    let output = [];
    this.resultArr[idKey] = value;
    while(this.resultArr[this.pointer] !== undefined){
        output.push(this.resultArr[this.pointer])
        this.pointer++
    }
    console.log(output)
    return output
};

/** 
 * Your OrderedStream object will be instantiated and called as such:
 * var obj = new OrderedStream(n)
 * var param_1 = obj.insert(idKey,value)
 */