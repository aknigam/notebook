package main

import (
	"fmt"
	"time"
)

type User struct {
	Id        string   `json:"uid"`
	Email     string   `json:"email"`
	Interests []string `json:"interests"`
}

func main() {
	start := time.Now()
	fmt.Println(start)
}
