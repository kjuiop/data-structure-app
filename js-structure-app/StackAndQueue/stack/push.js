class Node {
    constructor(value) {
        this.value = value;
        this.next = null;
    }

}

class Stack {
    constructor() {
        this.first = null;
        this.last = null;
        this.size = 0;
    }

    push(val) {
        let newNode = new Node(val);

        if (!this.first) {
            this.first = newNode;
            this.last = newNode;
            this.size++;
            return;
        }

        let temp = this.first;
        this.first = newNode;
        this.first.next = temp;
        this.size++;
    }
}

let stack = new Stack();
stack.push("HELLO")
stack.push("Friend")

console.log("stack : ", stack)