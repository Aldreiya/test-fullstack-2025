# Test Fullstack 2025

Repository ini berisi dua program berbeda sebagai latihan fullstack:

## 📂 Struktur Repository

```
.
├── factorial.js          # Program sederhana JavaScript (faktorial)
└── Login/                # Aplikasi login sederhana menggunakan Go + Redis
├── go.mod
├── main.go
└── ...

````

## 1. Program Faktorial (JavaScript)

File: `factorial.js`

### Deskripsi
Program ini menghitung faktorial dari suatu bilangan bulat **n**, kemudian membaginya dengan `2^n` dan membulatkan hasilnya ke atas (*ceil*).  
Fungsi utama:
```js
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
```

## 2. Sistem Login (Go + Redis)

Folder: `Login/`

### Teknologi
- **Go Fiber** (framework web)
- **Redis** (untuk penyimpanan user)

### Struktur Data di Redis
Key disimpan dengan format:
```

login\_<username>

````

Value berupa JSON string:
```json
{
  "realname": "Aberto Doni Sianturi",
  "email": "adss@gmail.com",
  "password": "f7c3bc1d808e0...441" 
}
````

> Password disimpan dalam bentuk **SHA1 hash**.

### Cara Menjalankan

1. Pastikan Redis sudah jalan (default: `localhost:6379`).

   ```bash
   redis-cli ping
   # output: PONG
   ```

2. Masuk ke folder `Login/`:

   ```bash
   cd Login
   ```

3. Jalankan aplikasi Go:

   ```bash
   go run main.go
   ```

4. Aplikasi berjalan di:

   ```
   http://localhost:3000
   ```

### Endpoint API

* **POST /login**

  * Request Body:

    ```json
    {
      "username": "aldre",
      "password": "123456"
    }
    ```
  * Response jika berhasil:

    ```json
    {
      "status": "success",
      "message": "Welcome, Aberto Doni Sianturi"
    }
    ```
  * Response jika gagal:

    ```json
    {
      "status": "error",
      "message": "Invalid username or password"
    }
    ```
