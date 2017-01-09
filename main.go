package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"fmt"

	"errors"

	"github.com/gorilla/mux"
)

// Task ...
type Task struct {
	Meta struct {
		UUID string `json:"UUID"`
		Task string `json:"task"`
	} `json:"Meta"`
	A        int    `json:"a"`
	B        int    `json:"b"`
	Error    error  `json:"error"`
	Operator string `json:"operator"`
}

// Result ...
type Result struct {
	UUID     string
	A        int
	B        int
	Operator string
	Result   string
	Err      error
}

func main() {
	ServiceAPI()
}

// ServiceAPI ...
func ServiceAPI() {
	addr := ":8081"
	log.Fatal(http.ListenAndServe(addr, handlers()))
}

// Handlers ...
func handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/arithsubscriber", handlerArithSubs)
	return r
}

func handlerArithSubs(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var task Task
	if err := json.Unmarshal(body, &task); err != nil {
		log.Fatal(err)
	}

	result := task.Calculator()
	resp, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("DevOpsGIG", "arithsubscriber")
	w.Write(resp)
}

// Calculator solve the task and returns a result struct
func (t Task) Calculator() Result {
	var r = Result{
		UUID:     t.Meta.UUID,
		A:        t.A,
		B:        t.B,
		Operator: t.Operator,
	}

	switch t.Operator {
	case "+":
		r.Result = strconv.Itoa(r.A + r.B)
	case "-":
		r.Result = strconv.Itoa(r.A - r.B)
	case "*":
		r.Result = strconv.Itoa(r.A * r.B)
	case "/":
		if r.B != 0 {
			r.Result = fmt.Sprintf("%.3f", float64(r.A)/float64(r.B))
		} else {
			r.Err = errors.New("zero division")
		}
	}
	return r
}
