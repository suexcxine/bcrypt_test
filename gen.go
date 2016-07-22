package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	cost = bcrypt.MinCost
)

func main() {
	// 设定生成随机字符串的数量
	n := 20
	if len(os.Args) > 1 {
		n, _ = strconv.Atoi(os.Args[1])
	}

	// 生成指定数量的随机字符串并生成hash, 写入stdout
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var v int64
	for i := 0; i < n; i++ {
		v = 100000000000 + (int64)(r.Intn(100000000000))
		s := strconv.FormatInt(v, 36)
		hashed, err := bcrypt.GenerateFromPassword([]byte(s), cost)
		if err != nil {
			fmt.Println("encrypt error")
		}
		// 写入hash结果
		os.Stdout.WriteString(s)
		os.Stdout.WriteString(" ")
		os.Stdout.Write(hashed)
		os.Stdout.WriteString("\n")
	}
}
