package main

import (
"log"
"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
message := []byte("Hello, web! YES, I was able do this!!!")
_, err := writer.Write(message)
if err != nil {
 log.Fatal(err)
}
}
func main() {
http.HandleFunc("/hello", viewHandler)
err := http.ListenAndServe("localhost:8087", nil)
log.Fatal(err)
}