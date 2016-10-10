package main

import (
	"salad/store"
	"salad/command"
	"time"
	"fmt"
	"math/rand"
	"os"
	"net/http"
	"log"
	"strings"
	"encoding/json"
)

func startServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	address := os.Getenv("ADDRESS")

	http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {
		cmd := req.PostFormValue("text")
		w.Header()["Content-Type"] = []string{"application/json"}
		jsonData := make(map[string]string)
		jsonData["response_type"] = "in_channel"
		jsonData["text"] = command.ExecuteCommandLine(cmd)
		jsonBytes, _ := json.Marshal(jsonData)
		w.Write(jsonBytes)
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
