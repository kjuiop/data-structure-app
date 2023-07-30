# 해시 테이블이란?

---

- 해시 알고리즘
- 해시 테이블에서 충돌이 일어나는 경우와 충돌이 무슨 의미인지
- 충돌을 해결하는 방법
    - 개별 체이닝
    - 선형 조사법
- 해시테이블
    - 키-값 쌍을 저장하는데 사용
    - 해시 테이블의 키는 순서를 가지지 않음
    - 모든 언어에는 내장 해시 테이블 코드가 존재한다.
- 좋은 해시함수의 조건
    - 속도
    - 무언가를 찾거나, 바뀌거나, 제거하기 위해 접근할 때마다 해시 함수를 다시 실행해야 한다.
    - 기본적으로 일관된 방식으로 분배를 해서 다른 것들과 겹치지 않아야 함
    - 보관장소가 고르게 퍼져야함
- 해시테이블의 특징
  - 해시테이블은 key, value 의 쌍으로 이루어져있다.
  - 해시테이블은 key 를 빠르게 찾을 수 있다.
  - key 를 빠르게 추가할 수 있다.
  - 해시테이블은 가능한 넓게 퍼뜨려서 데이터를 저장해야 한다.
  - 같은 입력값을 넣으면 언제나 같은 출력값이 나와야 한다.
  - 요소들 간의 충돌은 가능한 적게 일어나야 한다.

# 해시함수 작성하기

---

```java
function hash(key, arrayLen) {

    let total = 0;
    for (let char of key) {
        let value = char.charCodeAt(0) - 96;
        total = (total + value) % arrayLen;
    }
    return total;
}

console.log('result : ', hash("asdf", 11))
```

- 이 함수는 string 에 대해서만 해시 처리를 진행할 수 있다.
- BigO 의 상수값의 시간 값을 가지지 않는다.
- 무작위성이 떨어진다. 데이터가 뭉치기 쉽다.

### 소수를 사용하여 순회의 최대값을 적용

```java
function hashBetter1(key, arrayLen) {

    let total = 0;
    let WEIRD_PRIME = 31;
    for (let i=0; i<Math.min(key.length, 100); i++) {
        let char = key[i];
        let value = char.charCodeAt(0) - 96;
        total = (total * WEIRD_PRIME + value) % arrayLen;
    }
    return total;
}

console.log('result : ', hashBetter1("asdf",
```

- 소수
- 배열의 길이가 소수인 경우 충돌의 가능성이 줄어든다.
- 맨 앞의 백글자를 가져와서 값을 만듬
- 상수 값의 성능을 가짐

# 충돌 처리

---

- 해시 함수를 사용할 때 데이터가 아주 많이 있는 경우라면 충돌이 어느정도 일어날 수 있음
- 충돌을 해결하는 전략
  - 개별 체이닝
    - 같은 장소에 여러 데이터를 저장할 때 배열이나 연결 리스트 등과 같은 것을 활용하여 이중 데이터 구조를 사용하는 것
  - 직선 탐색법
    - 각 위치에 하나의 데이터만 저장한다는 규칙을 적용
    - 충돌이 발생하면 다음 빈 칸이 어디인지 확인함

## SET/GET

---

- SET
  - key 와 value 값을 입력한다.
  - key-value 쌍을 해시테이블에 개별 체이닝을 통해 저장한다.

- GET
  - key 값을 받는다.
  - Key 를 해쉬처리한다.
  - 루프를 돌면서 해당 값을 찾아낸다.
  - 키를 찾지 못하면 undefinded 를 반환한다.

### Keys / Values

---

- keys
  - 루프를 통해 모든 해쉬 테이블의 키배열을 반환

- values
  - 루프를 통해 모든 해쉬 테이블의 value 배열을 반환
  - 같은 값을 가진 요소들도 존재할 수 있음

# BigO HashTables

---

- Insert : O(1)
- Deletion : O(1)
- Access : O(1)




