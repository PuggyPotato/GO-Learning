package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Hello From Backend!")
}

func main(){
	http.HandleFunc("/",handler)
	fmt.Println("Server Is running on http://localhost:8000")
	http.ListenAndServe(":8000",nil)
}