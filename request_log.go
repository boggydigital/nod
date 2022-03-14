package nod

import "net/http"

//RequestLog is a trivial logging middleware that logs the URL of the request
func RequestLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Log("%s %s", r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}
