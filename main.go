package main

import (
        "fmt"
        "io/ioutil"
        "net/http"
        "strings"
)

func readResCodeResDelay(r *http.Request) (string, string) {
        var resCode string
        var resDelay string
        keys, ok := r.URL.Query()["resCode"]
        if !ok || len(keys[0]) < 1 {
                resCode = "200"
        } else {
                resCode = keys[0]
                // ex: "200"
        }
        keys, ok = r.URL.Query()["resDelay"]
        if !ok || len(keys[0]) < 1 {
                resDelay = "100"
        } else {
                resDelay = keys[0]
                // ex: "100"
        }
        return resCode, resDelay
}

func readBody(r *http.Request) []string {
        bodyBytes, err := ioutil.ReadAll(r.Body)
        if err != nil {
                return []string{}
        }
        bodyString := string(bodyBytes)
        bodyAryString := strings.Split(bodyString, "\n")
        return bodyAryString
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("Req: %s %s %s\n", r.Host, r.URL.Path, r.URL.Query())

        // read req-url-parameters: resCode resDelay
        resCode, resDelay := readResCodeResDelay(r)
        fmt.Println(resCode, resDelay)

        // read body as []string
        bodyAry := readBody(r)
        fmt.Println(bodyAry[0])
}

func main() {

        // We register our handlers on server routes using the
        // `http.HandleFunc` convenience function. It sets up
        // the *default router* in the `net/http` package and
        // takes a function as an argument.
        http.HandleFunc("/", handleFunc)
        fmt.Println("== Listening on :8080")
        http.ListenAndServe(":8080", nil)
}


