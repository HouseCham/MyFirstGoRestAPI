package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(writer http.ResponseWriter, reader *http.Request){
	writer.Write([]byte("Hello World!"))
}

func main(){
	router := mux.NewRouter();

	router.HandleFunc("/", homeHandler)

	http.ListenAndServe(":3000", router);
}