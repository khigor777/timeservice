package main

import "github.com/khigor777/timeservice"

func main() {
	ts := timeservice.NewServer("8080")
	ts.Run()

}
