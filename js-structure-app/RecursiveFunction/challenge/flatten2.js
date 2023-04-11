
function flatten2(oldArr){
    let newArr = []
    for(let i = 0; i < oldArr.length; i++){
        if(Array.isArray(oldArr[i])){
            newArr = newArr.concat(flatten2(oldArr[i]))
        } else {
            newArr.push(oldArr[i])
        }
    }
    return newArr;
}

console.log("result : ", flatten2([1, 2, 3, [4, 5] ]));
console.log("result : ", flatten2([1, [2, [3, 4], [[5]]]]));
console.log("result : ", flatten2([1],[2],[3]));
console.log("result : ", flatten2([[[[1], [[[2]]], [[[[[[[3]]]]]]]]]]));
