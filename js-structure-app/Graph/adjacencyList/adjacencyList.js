
class Graph {
    constructor() {
        this.adjacencyList = {}
    }
    addVertex(vertex) {
        if (!this.adjacencyList[vertex]) {
            this.adjacencyList[vertex] = [];
        }

    }
}

let g = new Graph();
g.addVertex("sample");
console.log("result : ", g);


