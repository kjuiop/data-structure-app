function addUpTo2(n) {
    return n * (n+1) / 2;
}


let time1 = performance.now();
addUpTo2(1000000000);
let time2 = performance.now();

console.log(`addUpTo2 Time Elapsed: ${(time2 - time1) / 1000} seconds.`);


