var merge = function(nums1, m, nums2, n) {  
    if(nums1 === 0) return;
    else {
        for(let i=m; i < nums1.length; i++){
            nums1[i] = nums2.pop()
        }
    }
    nums1.sort((a,b) => a - b)
};