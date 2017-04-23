package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/haya14busa/go-openbrowser"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 0, "port to listen on")
	flag.Parse()

	html, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	if port == 0 {
		port = ln.Addr().(*net.TCPAddr).Port
	}
	url := fmt.Sprintf("http://localhost:%v", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(html)
	})

	go openbrowser.WaitAndStart(url)

	fmt.Printf("Open %v", url)
	log.Fatal(http.Serve(ln, mux))
}
