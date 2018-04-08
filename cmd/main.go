package main

import (
	"fmt"
	"log"
	// "os"

	"github.com/Gin/node-tracker/server"
)

func main() {
	fmt.Println("hello world!")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	// f, err := os.OpenFile("error.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	// if err != nil {
		// log.Fatal(err)
	// }
	// defer f.Close()
	// log.SetOutput(f)

	server := server.NewServer()
	server.Run()
}
