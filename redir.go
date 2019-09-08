package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

// FAILURE is a string presented to the client if there's problem with the request or configuration
const FAILURE = "Domain not configured or service error"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.Body.Close()
		host := strings.Split(r.Host, ":")
		if len(host) < 1 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, FAILURE)
			return
		}
		rr, err := net.LookupTXT("_redir." + host[0])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, FAILURE)
			fmt.Println(err)
			return
		}
		if len(rr) < 1 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, FAILURE)
			return
		}
		url := rr[0]
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		fmt.Fprintf(w, `<html><head><title>Redirecting...</title></head><body><a href="%s">%s</a></body></html>\n`, url, url)
	})
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		fmt.Println(err)
	}
}
