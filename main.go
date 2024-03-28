package main

import (
	"net/http"
	
	authcontroller "github.com/zihunn/belajar-go/controllers"
)


func main(){

	http.HandleFunc("/", authcontroller.Index)
}