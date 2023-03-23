package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const charset = "123456789!@#$%^&*()_=-{}<>abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var n int
	fmt.Print("Masukkan Panjang Karakter : ")
	fmt.Scanln(&n)
	fmt.Println(randomString(n))
}
