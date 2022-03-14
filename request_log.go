package nod

import "net/http"

func RequestLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Log("%s %s", r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}
