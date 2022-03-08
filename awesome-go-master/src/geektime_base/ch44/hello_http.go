package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello,world!")
	})

	http.HandleFunc("/time", func(writer http.ResponseWriter, request *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		writer.Write([]byte(timeStr))
	})
	http.ListenAndServe(":8080", nil)
}
