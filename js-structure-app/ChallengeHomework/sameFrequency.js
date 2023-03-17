/**
 * 문제를 이해하라
 * 문제를 해결할 바탕이 되는 계획들을 세워라
 * 작성한 내용들을 다 지우고, 다시 시작해서 계획들을 세워라
 * 해결 가능한 문제들은 풀고, 불가한 문제들은 단순화시켜라
 * 회고하고, 리팩토링하라
 */


/**
 * Frequency Counter - sameFrequency**
 *
 * Write a function called **sameFrequency.
 * ** Given two positive integers,
 * find out if the two numbers have the same frequency of digits.
 *
 *     Your solution MUST have the following complexities:
 *
 *     Time: O(N)
 *
 * Sample Input:
 *
 * 1. sameFrequency(182,281) // true
 * 2. sameFrequency(34,14) // false
 * 3. sameFrequency(3589578, 5879385) // true
 * 4. sameFrequency(22,222) // false`
 */

function sameFrequency(num1, num2) {

    let arr1 = num1.toString().split('');
    let arr2 = num2.toString().split('');

    let frequencyCounter1 = {}
    let frequencyCounter2 = {}

    for (let val of arr1) {
        frequencyCounter1[val] = (frequencyCounter1[val] || 0) + 1;
    }

    for (let val of arr2) {
        frequencyCounter2[val] = (frequencyCounter2[val] || 0) + 1;
    }

    console.log(frequencyCounter1);
    console.log(frequencyCounter2);

    for (let key in frequencyCounter1) {

        if (!frequencyCounter2[key]) {
            return false;
        }

        if (frequencyCounter1[key] !== frequencyCounter2[key]) {
            return false;
        }
    }

    return true;
}


console.log("sameFrequency : " + sameFrequency(182,281));
console.log("sameFrequency : " + sameFrequency(34,14));
console.log("sameFrequency : " + sameFrequency(3589578, 5879385));
console.log("sameFrequency : " + sameFrequency(22,222));


