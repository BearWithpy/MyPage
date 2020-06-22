package main

import (
	"net/http"

	"github.com/BearWithPy/Myapp"
)

func main() {
	http.ListenAndServe(":3000", Myapp.NewHttpHandler())
}
