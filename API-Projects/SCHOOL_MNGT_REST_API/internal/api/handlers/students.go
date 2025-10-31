package handlers

import "net/http"

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on Students Route"))
		return

	case http.MethodPost:
		w.Write([]byte("Hello POST method on Students Route"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Students Route"))
		return

	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Students Route"))
		return

	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Students Route"))
		return
	}
}
