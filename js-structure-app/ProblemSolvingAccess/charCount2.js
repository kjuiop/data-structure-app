
function charCount(str) {

    let result = {};
    for (let char of str) {

        if (char === " " || char === undefined) {
            continue;
        }

        char = char.toLowerCase();

        if (/[a-z0-9]/.test(char)) {
            // result[char] 값이 존재하면 ++, 없으면 1
            result[char] = ++result[char] || 1;
        }
    }

    return result;
}

console.log("charCount : ", charCount("123 !!! 232 Hello Cowhello"));