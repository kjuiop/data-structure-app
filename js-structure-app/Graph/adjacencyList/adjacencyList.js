
class Graph {
    constructor() {
        this.adjacencyList = {}
    }
    addVertex(vertex) {
        if (!this.adjacencyList[vertex]) {
            this.adjacencyList[vertex] = [];
        }
    }
    addEdge(v1, v2) {
        this.adjacencyList[v1].push(v2);
        this.adjacencyList[v2].push(v1);
    }
}

let g = new Graph();
g.addVertex("Dallas");
g.addVertex("Seoul");
g.addVertex("Tokyo");

g.addEdge("Dallas", "Seoul");
g.addEdge("Seoul", "Tokyo");

console.log("result : ", g);


