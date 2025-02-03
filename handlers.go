package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homeHandler (w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w,`
<html>
<head><title>Главная</title></head>
<body>
	<h1>Shalom,Mykita,Verny 3Hungreat Bucks$ !!!</h1>
</body>
</html>
`)
}

func aboutHandler (w http.ResponseWriter,r *http.Request)  {
	fmt.Println(w, "About Us")
}

func MainHandler()  {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/about", aboutHandler)

	log.Println("Server run on: 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r ))
}