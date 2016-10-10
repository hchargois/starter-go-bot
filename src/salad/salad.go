package main

import (
	"salad/store"
	"salad/command"
	"time"
	"fmt"
	"math/rand"
	"os"
	"io"
	"net/http"
	"log"
	"strings"
)

func startServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	address := os.Getenv("ADDRESS")

	http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {
		cmd := req.PostFormValue("text")
		io.WriteString(w, command.ExecuteCommandLine(cmd))
	})
	addressPort := address + ":" + port

	fmt.Printf("Starting HTTP server on %v...\n", addressPort)
	log.Fatal(http.ListenAndServe(addressPort, nil))
}

func usage() {
	fmt.Println("Usage:\n  mongoo --cli <command> <args>\nOR\n  mongoo --daemon")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	store.Load()

	if len(os.Args) == 1 {
		usage()
		return
	}

	switch os.Args[1] {
	case "--cli":
		cmd := strings.Join(os.Args[2:], " ")
		fmt.Println(command.ExecuteCommandLine(cmd))
	case "--daemon":
		startServer()
	default:
		usage()
	}
}
