package main

import (
	"fmt"
	"net/http"

	"github.com/yeasin2002/url-shortener/routers"
)

func main() {

	http.HandleFunc("/", routers.GetUrl)
	http.HandleFunc("/redirect", routers.RedirectToOriginalUrl)

	server := &http.Server{Addr: ":8080"}
	fmt.Println("server Started  on :8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server error:", err)
	}
}
