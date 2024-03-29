package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	remote := flag.String("remote", "http://localhost:3001", "Server location")
	flag.Parse()
	start := time.Now()
	resp, err := http.Get(*remote)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s said: %s after %s\n", *remote, bytes, time.Since(start))
}
