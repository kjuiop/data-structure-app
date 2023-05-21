
class Student {
    constructor(firstName, lastName, year) {
        this.firstName = firstName;
        this.lastName = lastName;
        this.year = year;
    }
    fullName() {
        return `Your full name is ${this.firstName} ${this.lastName}`;
    }
}

let firstStudent = new Student("Colt", "Steele", 23);
let secondStudent = new Student("Blue", "Steele", 25);

console.log('result : ', firstStudent)
console.log('result : ', secondStudent)
console.log('result : ', secondStudent.fullName())
