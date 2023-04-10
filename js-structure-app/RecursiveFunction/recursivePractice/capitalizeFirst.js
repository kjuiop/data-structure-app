/**
 * Write a recursive function called capitalizeFirst. Given an array of strings, capitalize the first letter of each string in the array.
 * capitalizeFirst(['car','taco','banana']); // ['Car','Taco','Banana']
 */

function capitalizeFirst (arr) {
    // add whatever parameters you deem necessary - good luck!
    let newArr = [];

    if (arr.length === 0) {
        return newArr;
    }

    let str = arr[0];
    let char = str.charAt(0).toUpperCase() + str.slice(1);
    newArr.push(char);
    arr.shift();

    return newArr.concat(capitalizeFirst(arr));

}

console.log("result : ", capitalizeFirst(['car','taco','banana']));
