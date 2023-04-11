

function factorial(num) {
    if (num <= 1) {
        return 1;
    }

    return num * factorial(num-1);
}

console.log("result : ", factorial(1));
console.log("result : ", factorial(2));
console.log("result : ", factorial(4));
console.log("result : ", factorial(7));
