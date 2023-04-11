/**
 * Write a recursive function called capitalizeWords.
 * Given an array of words,
 * return a new array containing each word capitalized.
 */

function capitalizeWords(arr) {

    let newArr = [];

    if (arr.length === 0) {
        return newArr;
    }

    let str = arr[0];
    let char = str.toUpperCase();
    newArr.push(char);
    arr.shift();

    return newArr.concat(capitalizeWords(arr));

}

console.log("result : ", capitalizeWords(['car','taco','banana']));

