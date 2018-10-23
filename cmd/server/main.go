package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/khigor777/timeservice"
)

var port string

func init() {
	flag.StringVar(&port, "p", "37", "-p")
}

func main() {
	flag.Parse()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	ts := timeservice.NewServer(port)
	go ts.Run()
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	ts.Shutdown(ctx)
}
