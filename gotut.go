package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func randomize() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Whoa, Go is Neat")
	fmt.Fprintf(w, `<h1>Hey there</h1><p>Go is Fast!</p><p>... and Simple!</p>`)
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is my Http server app")
}

func main() {
	fmt.Println("Welcome to go")
	//randomize()

	resp, _ := http.Get("https://www.washingtonpost.com")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println(string_body)
	resp.Body.Close()
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
