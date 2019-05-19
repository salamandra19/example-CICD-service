package main

import (
	"fmt"
	"net/http"

	"github.com/powerman/structlog"
)

type countHandler struct{}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	var log = *structlog.New()
	http.Handle("/count", new(countHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
