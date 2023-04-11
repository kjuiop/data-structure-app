function power(num1, num2) {

    if (num2 <= 0) {
        return 1;
    }

    return num1 * power(num1, num2-1);
}


console.log("result : ", power(2, 0));
console.log("result : ", power(2, 2));
console.log("result : ", power(2, 4));

