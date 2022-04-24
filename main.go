package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	//fmt.Println("hello wolrd")
	http.HandleFunc("/",HelloHandler)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func HelloHandler(w http.ResponseWriter, _*http.Request){
	fmt.Fprintf(w,"Hello,there\n")
}