
// 정규표현식이 charCodeAt 보다 성능이 좋지 않다.

function charCount(str) {

    let result = {};
    for (let char of str) {

        if (char === " " || char === undefined) {
            continue;
        }

        char = char.toLowerCase();

        if (isAlphaNumeric(char)) {
            // result[char] 값이 존재하면 ++, 없으면 1
            result[char] = ++result[char] || 1;
        }
    }

    return result;
}

function isAlphaNumeric(char) {
    let code = char.charCodeAt(0);
    if (!(code > 47 && code < 58) &&
        !(code > 64 && code < 91) &&
        !(code > 96 && code < 123)) {
        return false
    }

    return true;
}

console.log("charCount : ", charCount("123 !!! 232 Hello Cowhello"));