function factorialFunction(n) {
    if (n < 0) {
        throw new Error("Bilangan bulat tidak boleh negatif");
    }
    let faktorial = 1;
    for (let i = 2; i <= n; i++) {
        faktorial *= i;
    }
    let pembagi = Math.pow(2, n);
    let hasil = Math.ceil(faktorial / pembagi);
    return hasil;
}

let n = 5;
console.log(factorialFunction(n));