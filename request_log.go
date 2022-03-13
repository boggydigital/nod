package nod

import "net/http"

func RequestLog(h func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Log("%s %s", r.Method, r.URL)
		h(w, r)
	}
}
