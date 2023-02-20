package main

import (
"log"
"net/http"
)

func write(writer http.ResponseWriter, massage string){
	_, err := writer.Write ([]byte(massage))
	if err != nil{
		log.Fatal(err)
	}
}
func englHandler(writer http.ResponseWriter, request *http.Request){
	write (writer, "Hello, Web! ICANN!!!")
}
func espHandler(writer http.ResponseWriter, request *http.Request){
	write (writer, "Hola, Web!!!")
}
func hindiHandler(writer http.ResponseWriter, request *http.Request){
	write (writer, "Namaste, Web!!!")
}
func main(){
	http.HandleFunc("/hello", englHandler)
	http.HandleFunc("/Hola", espHandler)
	http.HandleFunc("/namaste", hindiHandler)
	err := http.ListenAndServe("localhost:8087", nil)
	log.Fatal(err)
}