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
 * Write a function called averagePair.
 * Given a sorted array of integers and a target average,
 * determine if there is a pair of values in the array where the average of the pair equals the target average.
 * There may be more than one pair that matches the average target.
 *
 * Examples:
 *
 * averagePair([1,2,3],2.5) // true
 * averagePair([1,3,3,5,6,7,10,12,19],8) // true
 * averagePair([-1,0,3,4,5,6], 4.1) // false
 * averagePair([],4) // false
 *
 */

function averagePair(arr, average) {

    if (arr.length === 0) {
        return false;
    }

    let left = 0;
    let right = arr.length - 1;

    while(left < right) {

        let val = (arr[left] + arr[right]) / 2;
        if (val === average) {
            return true;
        } else if (left < right) {
            left++;
        } else {
            right--;
        }
    }


    return false;

}



console.log("averagePair : " + averagePair([1,2,3],2.5));
console.log("averagePair : " + averagePair([1,3,3,5,6,7,10,12,19],8));
console.log("averagePair : " + averagePair([-1,0,3,4,5,6], 4.1));
console.log("averagePair : " + averagePair([],4));