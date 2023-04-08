
function flatten(arr) {

    if (arr.length <= 1 && !Array.isArray(arr[0])) {
        return arr;
    }

    let result = [];
    if (Array.isArray(arr[0])) {
        flatten(arr[0]);
        arr.shift();
        return arr;
    }

    result.push(arr[0]);
    arr.shift();

    return result.concat(flatten(arr));
}

console.log("result : ", flatten([1, 2, 3, [4, 5] ]));
// console.log("result : ", flatten([1, [2, [3, 4], [[5]]]]));
// console.log("result : ", flatten([1],[2],[3]));
// console.log("result : ", flatten([[[[1], [[[2]]], [[[[[[[3]]]]]]]]]]));
