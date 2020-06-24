package main

import (
	"log"
	"net/http"

	"github.com/BearWithPy/todoapp/app"
	"github.com/urfave/negroni"
	//"github.com/BearWithPy/myapp"
)

func main() {
	//http.ListenAndServe(":3000", myapp.NewHttpHandler())
	m := app.MakeNewHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
