package main

/**
This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
If a copy of the MPL was not distributed with this file,
You can obtain one at http://mozilla.org/MPL/2.0/.
© 2014 Hiram J. Pérez
**/

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = flag.String(`p`, `8080`, `the port on wich we're serving`)
var path = flag.String(`path`, `.`, `path that will be served`)

func main() {
	flag.Parse()
	envPort := os.Getenv("PORT")
	servePort := (map[bool]string{true: envPort, false: *port})[envPort != ``]
	//ternary[ish] operator babe!

	servePort = fmt.Sprintf(":%v", servePort)

	fmt.Println("Serving files in", *path, "on port ", servePort)
	http.Handle("/", http.FileServer(http.Dir(*path)))
	err := http.ListenAndServe(servePort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
