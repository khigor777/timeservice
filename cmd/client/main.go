package main

import (
	"fmt"
	"log"
	"os"

	"github.com/khigor777/timeservice"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("need host and port parameter, example: time-client time.nist.gov 37")
	}
	host := os.Args[1]
	port := os.Args[2]
	res, err := timeservice.TcpClientRun(host, port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
