package router

import (
	"apitest/pages"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", pages.HomePage)

	log.Fatal(http.ListenAndServe(":8081", nil))
}


//Swc-9HK-uMH-nnp