package main

import (
	"fmt"
	"net/http"
)


func main(){

	http.HandleFunc("/", authController.Index)
}