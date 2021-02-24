package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func readBody(r *http.Request) []string {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []string{}
	}
	bodyString := string(bodyBytes)
	bodyAryString := strings.Split(bodyString, "\n")
	return bodyAryString
}

func readQueryParam(r *http.Request, queryParamName string, queryParamDefaultValue string) string {
	var qVal string
	keys, ok := r.URL.Query()[queryParamName]
	if !ok || len(keys[0]) < 1 {
		qVal = queryParamDefaultValue
	} else {
		qVal = keys[0]
	}
	return qVal
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Req: %s %s %s\n", r.Host, r.URL.Path, r.URL.Query())

	resCode := readQueryParam(r, "resCode", "200")
	resCodeInt, err := strconv.Atoi(resCode)
	if err != nil {
		fmt.Printf("Error converting resCode '%s' to int", resCode)
	}
	w.WriteHeader(resCodeInt)

	resDelay := readQueryParam(r, "resDelay", "100")
	resDelayInt, err := strconv.Atoi(resDelay)
	if err != nil {
		fmt.Printf("Error converting resDelay '%s' to int", resDelay)
	}
	time.Sleep(time.Duration(resDelayInt) * time.Millisecond)

	// read body as []string  (maybe empty)
	bodySlice := readBody(r)

	fmt.Println(bodySlice[0])
}

func main() {

	/*
		http://localhost:8080?resCode=201&resDelay=101

		http://D1:8080?resCode=201&resDelay=101
		http://D2:8080?resCode=201&resDelay=101
		http://D3:8080?resCode=201&resDelay=101
		http://www.tsf.pt
	*/

	// We register our handlers on server routes using the
	// `http.HandleFunc` convenience function. It sets up
	// the *default router* in the `net/http` package and
	// takes a function as an argument.
	http.HandleFunc("/", handleFunc)
	fmt.Println("== Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
