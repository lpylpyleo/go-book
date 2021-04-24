package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		prefix := "http://"
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}

		resp, err := http.Get(url)
		checkErr(err, "failed to get")

		// b, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("===============\nStatus: %s\n===============\n", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)

		resp.Body.Close()
		checkErr(err, "failed to read")
		// fmt.Printf("%s", b)
	}

}

func checkErr(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %s %v\n", msg, err)
		os.Exit(1)
	}
}
