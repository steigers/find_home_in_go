// For a stand-alone Go program this code would be in package main.
// The Go App Engine Runtime provides a special main package, so you should put
// HTTP handler code in a package of your choice.
package find_home

import (
  "fmt"
  "net/http"
)

func init() {
  http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello, world!")
}