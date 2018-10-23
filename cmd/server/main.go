package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/khigor777/timeservice"
)

func main() {

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	ts := timeservice.NewServer("8080")
	go ts.Run()
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	ts.Shutdown(ctx)

}
