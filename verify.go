package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
)

func main() {
	// 逐行verify并打印是否成功
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		splited := strings.Split(line, " ")
		s := splited[0]
		h := splited[1]
		ok := bcrypt.CompareHashAndPassword([]byte(h), []byte(s))
		if ok == nil {
			fmt.Println(s, h, "verify ok")
		} else {
			fmt.Println(s, h, "verify failed")
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
