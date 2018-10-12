package main

import (
	"fmt"
	"net/http"
  "os"
)

func main() {

  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	   fmt.Fprint(w, os.Args[1])
   })


   http.ListenAndServe(":80", nil)

}
