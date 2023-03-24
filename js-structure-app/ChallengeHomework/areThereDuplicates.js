/**
 * 문제를 이해하라
 * 문제를 해결할 바탕이 되는 계획들을 세워라
 * 작성한 내용들을 다 지우고, 다시 시작해서 계획들을 세워라
 * 해결 가능한 문제들은 풀고, 불가한 문제들은 단순화시켜라
 * 회고하고, 리팩토링하라
 */


/**
 * **Frequency Counter / Multiple Pointers - areThereDuplicates**
 *
 * Implement a function called, areThereDuplicates which accepts a variable number of arguments,
 * and checks whether there are any duplicates among the arguments passed in.
 * You can solve this using the frequency counter pattern OR the multiple pointers pattern.
 *
 * Examples:
 *
 * `1. areThereDuplicates(1, 2, 3) // false
 * 2. areThereDuplicates(1, 2, 2) // true
 * 3. areThereDuplicates('a', 'b', 'c', 'a') // true`
 *
 * **Restrictions:**
 * Time - O(n)
 * Space - O(n)
 *
 * **Bonus:**
 * Time - O(n log n)
 * Space - O(1)
 */

function areThereDuplicates(arr) {

    if (arr.length === 0) {
        return 0;
    }

    let frequency = {}
    for (let val in arr) {
        frequency[arr[val]] = (frequency[arr[val]] || 0) + 1
    }
    for (let val in frequency) {
        if (frequency[val] > 1) {
            return true;
        }
    }

    return false;
}

let array1 = [1,2,3];
let array2 = [1,2,2];
let array3 = [2,3,3,4];
let array4 = ["a", "b", "c", "a"];

console.log("areThereDuplicates : " + areThereDuplicates(array1));
console.log("areThereDuplicates : " + areThereDuplicates(array2));
console.log("areThereDuplicates : " + areThereDuplicates(array3));
console.log("areThereDuplicates : " + areThereDuplicates(array4));