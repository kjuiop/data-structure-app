

function isPalindrome(str) {

    if (str.length <= 1) {
        return true;
    }

    let arr = Array.from(str);
    let result = arr[0] === arr[arr.length-1];

    if (result) {
        arr.shift();
        arr.pop();
        return isPalindrome(arr.join(""));
    }

    return false;
}

console.log("result : ", isPalindrome('awesome'))
console.log("result : ", isPalindrome('foobar'))
console.log("result : ", isPalindrome('tacocat'))
console.log("result : ", isPalindrome('amanaplanacanalpanama'))
console.log("result : ", isPalindrome('amanaplanacanalpandemonium'))

