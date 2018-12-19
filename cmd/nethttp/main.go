package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/jhjaggars/dice/pkg/dice"
)

type Server struct{}

func handleError(err error, code int, w http.ResponseWriter) {
	log.Println(err)
	data, _ := json.Marshal(map[string]string{"error": err.Error()})
	w.WriteHeader(code)
	io.WriteString(w, string(data))
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d, err := dice.ParseDie(r.FormValue("dice"))
	if err != nil {
		handleError(err, 400, w)
		return
	}

	output := d.Roll()

	data, err := json.Marshal(output)
	if err != nil {
		handleError(err, 500, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(data))
}

func main() {
	var s Server
	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", s)
}
