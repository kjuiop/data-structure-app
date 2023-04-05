function factorial(num) {
    let total = 1;
    for (let i=num; i>0; i--) {
        total *= i;
    }
    return total;
}

let total = factorial(4);

console.log("answer : ", total);