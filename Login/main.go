package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

type User struct {
	RealName string `json:"realname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var ctx = context.Background()

func sha1Hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func login(rdb *redis.Client, username, password string) bool {
	key := "login_" + username
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("User not found")
		return false
	}

	var user User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		fmt.Println("Data error")
		return false
	}

	if user.Password == sha1Hash(password) {
		fmt.Printf("Welcome, %s!\n", user.RealName)
		return true
	}
	fmt.Println("Wrong password")
	return false
}

func register(rdb *redis.Client, username, realname, email, password string) error {
	key := "login_" + username
	// Cek apakah user sudah ada
	_, err := rdb.Get(ctx, key).Result()
	if err == nil {
		return fmt.Errorf("Username sudah terdaftar")
	}
	user := User{
		RealName: realname,
		Email:    email,
		Password: sha1Hash(password),
	}
	data, _ := json.Marshal(user)
	return rdb.Set(ctx, key, data, 0).Err()
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	for {
		var menu int
		fmt.Println("\n1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&menu)

		if menu == 1 {
			var username, realname, email, password string
			fmt.Print("Username: ")
			fmt.Scanln(&username)
			fmt.Print("Real Name: ")
			fmt.Scanln(&realname)
			fmt.Print("Email: ")
			fmt.Scanln(&email)
			fmt.Print("Password: ")
			fmt.Scanln(&password)
			err := register(rdb, username, realname, email, password)
			if err != nil {
				fmt.Println("Gagal register:", err)
			} else {
				fmt.Println("Register berhasil!")
			}
		} else if menu == 2 {
			var username, password string
			fmt.Print("Username: ")
			fmt.Scanln(&username)
			fmt.Print("Password: ")
			fmt.Scanln(&password)
			if !login(rdb, username, password) {
				// Tidak exit, hanya info gagal
			}
		} else if menu == 3 {
			fmt.Println("Keluar aplikasi.")
			os.Exit(0)
		} else {
			fmt.Println("Menu tidak valid")
		}
	}
}
