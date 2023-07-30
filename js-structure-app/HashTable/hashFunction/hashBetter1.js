
function hashBetter1(key, arrayLen) {

    let total = 0;
    // 소수
    // 배열의 길이가 소수인 경우 충돌의 가능성이 줄어든다.
    // 맨 앞의 백글자를 가져와서 값을 만듬
    let WEIRD_PRIME = 31;
    for (let i=0; i<Math.min(key.length, 100); i++) {
        let char = key[i];
        let value = char.charCodeAt(0) - 96;
        total = (total * WEIRD_PRIME + value) % arrayLen;
    }
    return total;
}

console.log('result : ', hashBetter1("asdf", 13))