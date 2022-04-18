package main

import (
	"fmt"
	"log"
	"net/http"
)

func contactUsHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "Request does not have data")
		return
	}

	firtName := request.FormValue("firstName")
	lastName := request.FormValue("lastName")
	email := request.FormValue("email")
	phoneNumber := request.FormValue("phoneNumber")

	fmt.Fprintf(writer, "Dear %s %s, your application successfully has been accepted! and sent message to your email %s and phone number %s", firtName, lastName, email, phoneNumber)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/contact-us", contactUsHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server couldn't start on port 8080")
	}
}
