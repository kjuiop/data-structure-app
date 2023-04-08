function reverse(str) {

    if (str.length <= 1) {
        return str;
    }

    let arr = Array.from(str);
    let lastChar = arr[arr.length-1];

    arr.pop();

    return lastChar + reverse(arr.join(""));
}

console.log("result : ", reverse('awesome'));
console.log("result : ", reverse('rithmschool'));