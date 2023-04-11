/**
 * Write a recursive function called nestedEvenSum.
 * Return the sum of all even numbers in an object which may contain nested objects.
 */

function nestedEvenSum (obj) {
    // add whatever parameters you deem necessary - good luck!

    let sum = 0;
    for (let key in obj) {
        if (typeof obj[key] === "object") {
            let subSum = nestedEvenSum(obj[key]);
            sum += subSum;
            continue;
        }

        if (obj[key] % 2 === 0) {
            sum += obj[key];
        }

    }

    return sum;
}


let obj1 = {
    outer: 2,
    obj: {
        inner: 2,
        otherObj: {
            superInner: 2,
            notANumber: true,
            alsoNotANumber: "yup"
        }
    }
}

let obj2 = {
    a: 2,
    b: {b: 2, bb: {b: 3, bb: {b: 2}}},
    c: {c: {c: 2}, cc: 'ball', ccc: 5},
    d: 1,
    e: {e: {e: 2}, ee: 'car'}
};

console.log("result : ", nestedEvenSum(obj1)); // 6
console.log("result : ", nestedEvenSum(obj2)); // 10