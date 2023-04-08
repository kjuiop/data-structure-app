
const isOdd = val => val % 2 !== 0;

function someRecursive(arr, callback) {
    if (arr.length === 0) {
        return false;
    }
    if (callback(arr[0])) {
        return true;
    }

    arr.shift();

    return someRecursive(arr, callback);
}

console.log("result : ", someRecursive([1,2,3,4], isOdd));
console.log("result : ", someRecursive([4,6,8,9], isOdd));
console.log("result : ", someRecursive([4,6,8], isOdd));
console.log("result : ", someRecursive([4,6,8], val => val > 10));
