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
list.push("FRIEND")
list.push("SEE")
list.push("YOU")


list.traverse();
