function collectOddValues2(arr) {
    let newArr = [];
    if (arr.length === 0) {
        return newArr;
    }

    if (arr[0] % 2 !== 0 ) {
        newArr.push(arr[0]);
    }

    newArr = newArr.concat(collectOddValues2(arr.slice(1)));
    return newArr;
}

let result = collectOddValues2([1,2,3,5,6,8]);
console.log("result : ", result);
