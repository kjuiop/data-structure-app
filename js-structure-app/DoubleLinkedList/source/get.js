
class Node {
    constructor(val) {
        this.val = val;
        this.next = null;
        this.prev = null;
    }
}

class DoubleLinkedList {
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
        node.prev = this.tail;
        this.tail = node;
        this.length++;
    }

    pop() {
        if (!this.head) {
            return undefined;
        }

        if (this.length === 1) {
            this.head = null;
            this.tail = null;
            this.length = 0;
            return;
        }

        let poppedNode = this.tail;
        this.tail = poppedNode.prev;
        this.tail.next = null;
        this.length--;
    }

    shift() {
        if (!this.head) {
            return undefined;
        }

        if (this.length === 1) {
            this.head = null;
            this.tail = null;
            this.length = 0;
            return;
        }

        let current = this.head.next;
        current.prev = null;
        this.head = current;
        this.length --;
    }

    unShift(val) {
        let node = new Node(val);
        if (this.length === 0) {
            this.head = node;
            this.tail = node;
            this.length++;
            return;
        }

        this.head.prev = node;
        node.next = this.head;
        this.head = node;
        this.length++;
    }

    get(idx) {

        if (idx < 0 || idx > list.length || list.length === 0) {
            return null;
        }

        let standard = list.length / 2;
        let current;

        if (idx < standard) {
            current = this.head;
            let count = 0;
            while(current) {

                if (idx === count) {
                    return current;
                }

                current = current.next;
                count++;
            }
            return;
        }

        current = this.tail;
        let count = this.length -1;
        while(current) {

            if (count === idx) {
                return current;
            }

            current = current.prev;
            count--;
        }

    }


    traverse() {
        let current = this.head;
        while(current) {
            console.log(current.val);
            current = current.next;
        }
    }
}

let list = new DoubleLinkedList();
list.push("Hello");
list.push("Friend");
list.push("BoB");
console.log("return : ", list.get(0));

