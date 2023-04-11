
function fib(num) {

    if (num <= 2 ) {
        return 1;
    }

    return fib(num-1) + fib(num-2);
}

console.log("result : ", fib(4));
console.log("result : ", fib(10));
console.log("result : ", fib(28));
console.log("result : ", fib(35));