package router

import (
	"net/http"
)

// MainRouter combines all sub-routers into a single main router
func MainRouter() *http.ServeMux {

	tRouter := teachersRouter()
	sRouter := studentsRouter()
	exRouter := execsRouter()

	sRouter.Handle("/", exRouter)
	tRouter.Handle("/", sRouter)

	return tRouter
}
