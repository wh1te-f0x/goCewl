package main

import (
	_"io"
	"log"
	_"net/http"
	_"os"
	_"time"
	"runner"
)

func main() {

	options := runner.ParseOptions()
	runner, err := runner.New(options)
	if err != nil {
		log.Fatal(err)
	}
	/* client := &http.Client{
		Timeout: 30 * time.Second,
	}
	response, err := http.NewRequest(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "Not Firefox")
	// Use defer to wait for the end of the execution
	defer response.Body.Close()
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n)*/
}