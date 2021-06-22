var maximumWealth = function(accounts) {   
    let total = 0;
    const output = [];
    
    for(let i = 0; i < accounts.length; i++){
        for(let j = 0; j < accounts[i].length; j++){
            total = total + accounts[i][j]
        }
        output.push(total)
        total = 0
    }
    return Math.max(...output)
};