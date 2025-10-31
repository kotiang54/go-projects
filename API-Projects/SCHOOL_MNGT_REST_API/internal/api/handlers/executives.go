package handlers

import "net/http"

func ExecutivesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on Executives Route"))
		return

	case http.MethodPost:
		w.Write([]byte("Hello POST method on Executives Route"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Executives Route"))
		return

	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Executives Route"))
		return

	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Executives Route"))
		return
	}
}
