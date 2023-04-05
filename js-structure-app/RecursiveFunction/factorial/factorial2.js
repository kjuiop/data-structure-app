
function factorial2(num) {
    if (num === 1) {
        return 1;
    }
    return num * factorial2(num-1);
}



let total = factorial2(4);

console.log("total : ", total);