package main

import (
	"net/http"
)

func main() {

	http.Handle("/", http.StripPrefix("/",http.FileServer(http.Dir("b"))))
	http.Handle("/anime.html/furthernotes/", http.StripPrefix("/anime.html/furthernotes/",http.FileServer(http.Dir("a"))))
	http.ListenAndServe(":8000", nil)

}




