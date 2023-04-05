function sumRange(num) {
    if (num === 1) {
        return 1;
    }

    return num + sumRange(num-1);
}

let num = sumRange(4);
console.log(num);
