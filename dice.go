package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func sum(nums []int) (total int) {
	for _, n := range nums {
		total += n
	}
	return
}

func single(sides int) int {
	return 1 + rand.Intn(sides-1)
}

func bones(num int, sides int) []int {
	rolls := make([]int, num)
	for i := range rolls {
		rolls[i] = single(sides)
	}
	return rolls
}

type Output struct {
	Rolls []int `json:"rolls"`
	Total int   `json:"total"`
}

type Server struct{}

func convertString(s string, defaultValue int) (r int) {
	if s == "" {
		r = defaultValue
	} else {
		parsed, err := strconv.Atoi(s)
		if err != nil {
			r = defaultValue
		} else {
			r = parsed
		}
	}
	return
}

func parseDie(s string) (num, sides int, err error) {
	num = 1
	sides = 6
	hasd := false
	for _, ch := range s {
		if ch == 'd' {
			hasd = true
		}
	}
	if !hasd {
		err = errors.New(fmt.Sprintf("die '%s' has no 'd'", s))
		return
	}
	parts := strings.Split(s, "d")
	num = convertString(parts[0], 1)
	sides = convertString(parts[1], 6)
	return
}

func handleError(err error, code int, w http.ResponseWriter) {
	log.Println(err)
	data, _ := json.Marshal(map[string]string{"error": err.Error()})
	w.WriteHeader(code)
	io.WriteString(w, string(data))
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	num, sides, err := parseDie(r.FormValue("dice"))
	if err != nil {
		handleError(err, 400, w)
		return
	}

	rolls := bones(num, sides)
	output := Output{rolls, sum(rolls)}

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
	http.ListenAndServe("0.0.0.0:8080", s)
}
