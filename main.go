package main

import (
	"net/http"

	"github.com/BearWithPy/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
