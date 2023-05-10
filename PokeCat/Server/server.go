package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func StartServer() {
	http.HandleFunc("/", test)
	http.ListenAndServe(":8080", nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	test_page, err := os.ReadFile("./PokeCat/HTML/test.html")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Fprintf(w, string(test_page))
}
