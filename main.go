package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := 8049
	http.Handle("/", http.FileServer(http.Dir("polymer")))
	fmt.Printf("server running on port %d ", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))

}
