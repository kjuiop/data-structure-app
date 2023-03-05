function addUpTo1(n) {
    let total = 0;
    for (let i=1; i<=n; i++) {
        total += i;
    }
    return total;
}

let t1 = performance.now();
addUpTo1(1000000000);
let t2 = performance.now();

console.log(`addUpTo1 Time Elapsed: ${(t2 - t1) / 1000} seconds.`);


