# 스택과 큐

---

- Stack
    - 후입선출 원칙을 다르는 데이터들의 모음
    - 가장 마지막으로 들어온게 가장 먼저 나감
    - 배열도 스택으로 활용할 수 있다.
        - 배열을 pop() 과 push() 를 활용하면 후입선출을 할 수 있기 때문이다.

- Queue
    - 큐는 선입선출
    - 배열로 큐를 활용할 수 있다.
        - 배열의 push() 와 shift() 를 활용하면 선입선출을 할 수 있다.
        - 그러나 O(n) 임으로, 직접 연결리스트를 만드는 것이 좋다.

# 스택 - Push

---

- 새로운 노드와 value 를 만든다.
- 만일, 노드가 스택에 없다면 first, last 에 새로운 노드를 추가한다.
- 만일 최소 하나의 노드가 잇다면, 변수를 만들고, 현재 first 스택을 만들어라
- first 를 새로운 노드로 리셋하라
- 만들어진 변수의 이전 노드의 next 를 설정하라
- size 를 1 만큼 더해라

# 스택 - Pop

---

- 만일 스택에 노드가 없다면 nil 을 리턴한다.
- temp 변수에 first 를 저장한다.
- 만일 하나의 노드만 있다면 first, last 를 모두 null 로 만든다.
- 만일 하나 이상의 노드가 있다면, first 의 next 를 현재 first 로 만들어준다.
- size 를 1 감소한다.
- 삭제된 node 의 value 를 리턴한다.

# 스택의 BigO

---

- Insertion : O(1)
- Removal : O(1)
- Searching : O(N)
- Access : O(N)

# 큐의 Enqueue

---

- 새로운 노드를 만들고 pass 해라
- 만일 queue 에 노드가 없으면 first 와 last 를 queue 에 적용해라
- queue 에 노드가 있다면 next 에 현재 노드를 적용하고, last 도 현재 노드를 적용해라

# 큐의 Dequeue

---

- first 가 없다면 null 을 반환해라
- first 를 저장해라
- 만일 first 와 last 가 같다면 null 로 설정해라
- 만일 하나 이상의 노드라면 first 의 next 를 설정해라
- 사이즈를 1 감소해라
- node 를 dequeue 해라

# 큐의 BigO

---

- Insertion : O(1)
- Removal : O(1)
- Searching : O(N)
- Access : O(N)