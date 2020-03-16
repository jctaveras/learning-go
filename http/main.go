package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logger struct{}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	customLogger := logger{}

	io.Copy(customLogger, resp.Body)
}

func (logger) Write(data []byte) (int, error) {
	fmt.Println(string(data))
	fmt.Println("Written Bytes: ", len(data))

	return len(data), nil
}
