
function productOfArray(arr) {

    if (arr.length === 0) {
        return 1;
    }

    let num = arr[0];
    let arr2 = arr.slice(1);

    console.log("num : ", num);
    console.log("arr2 : ", arr2);

    return num * productOfArray(arr2);
}

console.log("result : ", productOfArray([1,2,3]));
console.log("result : ", productOfArray([1,2,3,10]));