// piece of data - val
// reference to next node - next

class Node {
    constructor(val) {
        this.val = val;
        this.next = null;
    }
}

class SingleLinkedList {
    constructor() {
        this.head = null;
        this.tail = null;
        this.length = 0;
    }

    push(val) {
        let node = new Node(val);
        if (this.length === 0) {
            this.head = node;
            this.tail = node;
            this.length++;
            return;
        }

        this.tail.next = node;
        this.tail = node;
        this.length++;
    }

    pop() {

        if (!this.head) {
            return undefined;
        }

        let current = this.head
        let newTail = current;

        // 루프를 다 돌게하여 마지막 tail을 newTail 로 가져오게 함
        while(current.next) {
            // console.log("next : ", current.next)
            // goodbye
            // !

            // 현재 tail 을 바꿔주기 위한 변수
            newTail = current;
            // console.log("newTail : ", newTail)
            // hello
            // goodbye

            // 반환하기 위한 변수 + while 다음 변수로 가기위한 저장
            current = current.next;
            // console.log("current : ", current)
            // goodbye
            // !

        }

        this.tail = newTail
        this.tail.next = null;
        this.length--;

        if (this.length === 0) {
            this.head = null;
            this.tail = null;
        }
        return current;
    }

    shift() {
        if (!this.head) {
            return undefined;
        }

        let current = this.head;
        this.head = current.next;
        this.length--;

        if (this.length === 0) {
            this.tail = null;
        }
        return current;
    }

    unShift(val) {
        let node = new Node(val);
        if (this.length === 0) {
            this.head = node;
            this.tail = node;
            this.length++;
            return;
        }
        node.next = this.head;
        this.head = node;
        this.length++;
    }


    traverse() {
     let current = this.head;
     while(current) {
         console.log(current.val);
         current = current.next;
     }
    }
}

let list = new SingleLinkedList();
list.push("HELLO")
list.push("GOODBYE")
list.push("!")

list.pop()
list.shift()

list.unShift("FRIEND")
list.unShift("ONE")

list.pop();
list.pop();
list.pop();

list.push("END")
list.unShift("First")

list.traverse()