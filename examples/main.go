package main

import (
	"code.google.com/p/gopages/pkg"
	"net/http"

	//replace project with "code.google.com/p/gopages/examples"
	//generated package, required for initialization
	_ "code.google.com/p/gopages/examples/pages"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("images")))
	
	http.HandleFunc("/hello", gopages.Handler("src/hello.ghtml"))
	http.HandleFunc("/echo", gopages.Handler("src/echo.ghtml"))
	
	println("navigate to http://localhost:9999/hello")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}
