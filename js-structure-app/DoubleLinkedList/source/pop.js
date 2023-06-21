
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
list.push("BB");
list.pop();
list.pop();
list.pop();
list.traverse()
