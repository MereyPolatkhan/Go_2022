package main

// import (
// 	"log"
// 	"net/http"
// )

// func write(writer http.ResponseWriter, message string) {
// 	_, err := writer.Write([]byte(message))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func engHandler(writer http.ResponseWriter, request *http.Request) {
// 	write(writer, "Hi!")
// }

// func kzHandler(writer http.ResponseWriter, request *http.Request) {
// 	write(writer, "Салем!")
// }

// func ruHandler(writer http.ResponseWriter, request *http.Request) {
// 	write(writer, "Привет!")
// }

// func main() {
// 	http.HandleFunc("/eng", engHandler)
// 	http.HandleFunc("/kz", kzHandler)
// 	http.HandleFunc("/ru", ruHandler)
// 	err := http.ListenAndServe("localhost:8080", nil)
// 	log.Fatal(err)
// }
