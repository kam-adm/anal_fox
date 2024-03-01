package main

import (
	"encoding/json"
	"fmt"
	"fox_anal/models"
	simnet "github.com/webkarlon/simnet"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("start")
	myServer := simnet.NewServer(&simnet.Server{
		PortHTTP:        8080,
		ListenAddress:   "localhost",
		ShutdownTimeout: 2,
	})

	myServer.AddRouter(http.MethodPost, "/aggregation", Auth, Aggregation)

	go myServer.Start()

	osSignalsCh := make(chan os.Signal, 1)
	signal.Notify(osSignalsCh, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-osSignalsCh

	err := myServer.Stop()
	fmt.Println("close")
	if err != nil {
		panic(err)
	}
}

func Aggregation(w http.ResponseWriter, r *http.Request) {
	var req models.Request
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(req.Seller.Rv)

	for i, sclad := range req.Compositions {
		fmt.Println(i, sclad)
	}

	w.Write([]byte("Aggregation"))
}

func Auth(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Auth\n"))
}
