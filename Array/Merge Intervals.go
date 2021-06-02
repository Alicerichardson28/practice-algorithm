var merge = function(intervals) {
    //looping through the first index and compair which one is small
    //set the smaller number = A and set the bigger number to B
    //nested for loop 
    //looping through i
      //looping through j 
    
    if(intervals.length === 1)return intervals;
    
    intervals.sort((intervalA, intervalB) => intervalA[0] - intervalB[0]);
    
    let lastStart = intervals[0][0];
    let lastEnd = intervals[0][1];
    const newIntervals = []
    
    for(let i =0; i < intervals.length; i++){
        const curInterval = intervals[i];
        const curStart = curInterval[0]
        const curEnd = curInterval[1]
        
        if(curStart <= lastEnd){
            //overlap
            lastEnd = Math.max(lastEnd, curEnd);
        } else {
            newIntervals.push([lastStart, lastEnd]);
            lastStart = curStart;
            lastEnd = curEnd;
        }
    }
    newIntervals.push([lastStart, lastEnd]);
    return newIntervals;
};