# 문제 해결 패턴

---

### 문제 해결 패턴

---

- 문제를 해결하기 위한 계획을 수립한다.
- 일반적인 문제 해결 패턴을 습득한다.


### 일반적인 문제 해결  패턴

- Frequency Counter
- Multiple Pointers
- Sliding Window
- Divide and Conquer
- Dynamic Programming
- Greedy Algorithms
- Backtracking


## Frequency Counter

---

- 자바스크립트의 객체를 사용해서 다양한 값과 빈도를 수집하는 것
    - 알고리즘과 과제에 있는 여러 데이터와 입력값이 서로 비슷한 값으로 구성되어 있는지
    - 서로간의 아나그램인지
    - 값이 다른 값에 포함되는지 여부를 비교하거나
    - 데이터를 입력값이나 두 개 이상의 빈도 혹은 특정하게 발생하는 빈도와 비교할 때 유용


- 기존 O(n**2)

```java
function same(a, b) {

    let result = true;

    if (a.length !== b.length) {
        return false;
    }

    for (let e of a) {

        if (isNaN(e) && !Number.isInteger(e)) {
            result = false;
            return false;
        }

        let index = b.indexOf(e*e);
        if (index === -1) {
            result = false;
            return false;
        }

        b.splice(index, 1);
    }

    return result;
}
```

- Frequency Counter  O(n)


```java
function same(arr1, arr2) {

    if (arr1.length !== arr2.length) {
        return false;
    }

    let frequencyCounter1 = {};
    let frequencyCounter2 = {};
    // 값이 존재한다면 +1 , 없다면 1로 초기화
    for (let val of arr1) {
        frequencyCounter1[val] = (frequencyCounter1[val] || 0) + 1
    }
    for (let val of arr2) {
        frequencyCounter2[val] = (frequencyCounter2[val] || 0) + 1
    }
    console.log(frequencyCounter1);
    console.log(frequencyCounter2);
    for (let key in frequencyCounter1) {
        // 제곱에 대한 키값이 존재하는지를 체크
        if (!(key ** 2 in frequencyCounter2)) {
            return false;
        }

        // 빈도수가 다르냐를 체크
        if (frequencyCounter2[key ** 2] !== frequencyCounter1[key]) {
            return false;
        }
    }

    return true;
}
```


## Multiple Pointer

---

- 인덱스나 위치에 해당하는 포인터나 값을 만든 다음 특정 조건에 따라 중간 지점에서부터 시작지점이나 끝 지점이나 양쪽 지점을 향해 이동시키는 것
- 배열이나 문자열과 같은 일종의 선형구조나 이중 연결 리스트, 단일 연결 리스트를 만드는 것
- 한 쌍의 값이나 조건을 충족시키는 무언가를 찾는다는 것
- 포인터 변수는 배열이나 문자열의 특정 위치를 가리키는 것
- 서로를 향해 이동하거나 같은 방향으로 이동하든 상관은 없으나 포인터를 두개 사용하는 것

- 기존 O(n**2)

```java
function sumZero(array) {

    for (let i=0; i<array.length; i++) {
        for (let j= i+1; j<array.length; j++) {
            if (array[i] + array[j] === 0) {
                return [array[i], array[j]]
            }
        }
    }
}
```


- 다중 포인터 적용 O(n)

```java
function sumZero(array) {

    let left = 0;
    let right = array.length -1;

    while(left < right) {
        let sum = array[left] + array[right];
        if (sum === 0) {
            return [array[left], array[right]];
        } else if (sum > 0) {
            right--;
        } else {
            left++;
        }
    }
}
```



### 다중 포인터 문제 예제

---

- 정렬된 배열을 취하는 countUniqueValues 라는 함수를 작성해라
- 분류(assorted) 가 아닌 정렬(assorted) 된 배열이어야 한다.
- 다만 오름차순이어야 한다.
- 조건에 따라 두 포인터가 특정 방향으로 움직이도록 하고, countUniqueValues 함수를 구현하여
  정렬된 배열을 전달하면 해당 배열의 고유한 값의 개수를 반환하도록 할 것
- 음수가 포함될 수 있으나 항상 정렬된 상태로 있을 것

[1,1,1,1,1,2] -> 2
[1,2,3,4,4,4,7,7,12,12,13] -> 7
[] -> 0
[-2,-1,-1,0,1] -> 4

### Me

```java
function CountUniqueValues(arr) {

    let i = 0;
    let j = i+1;

    let result = 0;

    if (arr.length > 0) {
        result++;
    }

    while (j < arr.length) {

        if (arr[i] !== arr[j]) {
            result++;
        }
        i++;
        j++;

    }

    return result;
}
```

### 답안

```java
function CountUniqueValues(arr) {

    if (arr.length === 0) {
        return 0;
    }

    let i = 0;

    for (let j=1; j<arr.length; j++) {

        if (arr[i] !== arr[j]) {
            i++;
            arr[i] = arr[j]
        }
    }

    return i+1;
}
```


## Sliding Window

---

- 배열이나 문자열과 같은 일련의 데이터를 입력하거나 특정 방식으로 연속적인 해당 데이터의 하위 집합을 찾는 경우에 유용
- 창문을 하나 만들어야 함
- 이 창문은 단일 변수, 하위 배열 또는 필요한 경우 다른 문자열도 될 수 있음
- 조건에 따라 창문을 이동시키며, 시작 위치에서 시작하면 보통 왼쪽에서 오른쪽으로 이동함
    - 오른쪽 → 왼쪽, 가운데 위치에서 시작할 수도 있으나, 보통 왼쪽 → 오른쪽으로 이동

### 슬라이딩 윈도우 - MaxSubarraySum

---

- 배열과 숫자 하나를 전달하고, 서로 마주한 두 숫자의 가장 큰 합계를 찾아내는 것

```java
maxSubarraySum(1,2,5,2,8,1,5],2)  // 10
maxSubarraySum(1,2,5,2,8,1,5],4)  // 17
maxSubarraySum(4,2,1,6],1) // 6
maxSubarraySum(4,2,1,6,2],4) // 13
maxSubarraySum(],4) // null
```


- 중철 for 문 (On**2)

```java
function maxSubarraySum(arr, num) {
    if (num > arr.length) {
        return null;
    }

    let max = -Infinity;
    for (let i=0; i<arr.length - num + 1; i++) {
        let temp = 0;
        for (let j=0; j<num; j++) {
            temp += arr[i+j];
        }
        if (temp > max) {
            max = temp;
        }
    }
    return max;
}
```


- 슬라이딩 윈도우 O(n)

```java
function maxSubarraySum(arr, num) {
    let maxSum = 0;
    let tempSum = 0;
    if (arr.length < num) {
        return null;
    }
    for (let i=0; i<num; i++) {
        maxSum += arr[i];
    }
    tempSum = maxSum;
    for (let i=num; i<arr.length; i++) {
        tempSum = tempSum - arr[i-num] + arr[i];
        maxSum = Math.max(maxSum, tempSum);
    }
    return maxSum;
}
```


### Divide and Conquer

---

- 분할 정복
- 이 알고리즘은 주로 배열이나 문자열 같은 큰 규모의 데이터 셋을 처리한다.
- 연결 리스트나 트리가 될 수 있다.
- 값을 찾기 위해 왼쪽에서 오른쪽 끝까지 이동하는 것보다는 배열을 작은 조각으로 세분화하여 각 조각들을 어디로 이동시킬지 결정하는 작업
    - 큰 덩어리를 작은 조각으로 나누는 것



### Question : search

---

- 오름차순 정렬된 integer 배열이 들어왔을 때, number 로 들어온 parameter 가 배열의 어느 요소에 있는지 구하라

- 선형탐색 (풀이: me)  → O(n)

```java
function search(arr, num) {

    if (arr.length <= 0) {
        return -1;
    }

    for (let i=0; i<arr.length; i++) {
        if (arr[i] === num) {
            return i+1;
        }
    }

    return -1;
}
```

- 이진탐색 (답안)  → O(log n)

```java
function search(arr, num) {

    let min = 0;
    let max = arr.length -1;

    while(min <= max) {
        let middle = Math.floor((min + max) / 2);
        let currentElement = arr[middle];

        if (arr[middle] < num) {
            min = middle + 1;
        } else if (arr[middle] > num) {
            max = middle -1;
        } else {
            return middle+1;
        }
    }

    return -1;
}
```