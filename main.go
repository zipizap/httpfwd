package main

import (
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	_ "net/url"
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
	log.Infof("---------------------")
	bodySlice := readBody(r)
	log.Infof("Req: %s %s \t  %s\n%s\n", r.Host, r.URL.Path, r.URL.Query(), strings.Join(bodySlice[:], "\n"))

	resCode := readQueryParam(r, "resCode", "200")
	resCodeInt, err := strconv.Atoi(resCode)
	if err != nil {
		log.Errorf("Error converting resCode '%s' to int: %v", resCode, err)
	}
	w.WriteHeader(resCodeInt)

	resDelay := readQueryParam(r, "resDelay", "100")
	resDelayInt, err := strconv.Atoi(resDelay)
	if err != nil {
		log.Errorf("Error converting resDelay '%s' to int: %v", resDelay, err)
	} else {
		time.Sleep(time.Duration(resDelayInt) * time.Millisecond)
	}

	freqBody_multilinestring := strings.Join(bodySlice[1:], "\n")
	//freqBody_multilinestring_urlencoded := url.QueryEscape(freqBody_multilinestring)
	freqUrl := bodySlice[0]
	freqContentType := "application/x-www-form-urlencoded"
	log.Infof(">> FwReq: %s \n%s\n", freqUrl, freqBody_multilinestring)
	fres, err := http.Post(
		freqUrl,
		freqContentType,
		strings.NewReader(freqBody_multilinestring),
	)
	if err != nil {
		log.Errorf("Error with forward-request to '%s': %v", freqUrl, err)
		return
	}
	fresBody, err := ioutil.ReadAll(fres.Body)
	if err != nil {
		log.Errorf("Error with forward-response body: %v", err)
		return
	}
	log.Infof("Sending back response\n")
	w.Write(fresBody)

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
	log.Info("== Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
