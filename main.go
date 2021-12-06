package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		var remoteAddress string
		forwardedFor := r.Header.Get("X-Forwarded-For")
		if forwardedFor != "" {
			remoteAddress = forwardedFor
		} else {
			remoteAddress = r.RemoteAddr
		}
		rw.Write([]byte(remoteAddress))
	})
	http.ListenAndServe(":80", nil)
}
