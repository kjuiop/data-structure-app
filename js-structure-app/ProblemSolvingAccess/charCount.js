/**
 *
 * 결과 값 객체를 반환
 * String Loop 시작
 * 숫자나, 문자열이며, 객체에 값이 있다면 count 을 1 추가
 * 숫자나, 문자열이며, 객체에 값이 없다면 count 을 1로 키값 추가
 * 숫자나 문자가 아니면 Skip
 * 결과 Object 반환
 */

function charCount(str) {

    let result = {};
    for (let i=0; i<str.length; i++) {

        let char = str[i];
        if (char === " " || char === undefined) {
            continue;
        }

        char = char.toLowerCase();
        if (!/[a-z0-9]/.test(char)) {
            continue;
        }

        if (result[char] > 0) {
            result[char]++;
            continue;
        }

        result[char] = 1;
    }

    return result;
}

console.log("charCount : ", charCount("123 !!! 232 Hello Cowhello"));