package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getPeopleSimple(url string) ([]Person, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var people []Person
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &people)
	if err != nil {
		panic(err)
	}
	return people, nil
}

func getPeople(url string) ([]Person, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var people []Person
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&people)
	if err != nil {
		panic(err)
	}
	return people, nil
}

func call(url string) (response string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", bytes), nil
}

func main() {
	remote := flag.String("remote", "http://localhost:3001", "Server location")
	calls := flag.Int("calls", 20, "Number of calls to remote server")
	flag.Parse()
	start := time.Now()
	for i := 0; i < *calls; i++ {
		resp, err := call(*remote)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp)
	}
	fmt.Printf("Finished after %s\n", time.Since(start))
}
