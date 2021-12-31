package main

import (
	"github.com/pescox/learning-go/proj/chat/api"
)

func main() {
	s := api.NewServer("127.0.0.1", 8080)
	s.Start()
}
