function compareByLen(str1, str2) {
    return str1.length - str2.length;
}

console.log("result : ", ["colt", "data structure", "golang", "algorithms"].sort(compareByLen));
