package main

import (
	"github.com/zograf/task/server"
)

func main() {
	srv := server.New()
	srv.Run()
}
