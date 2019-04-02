package features

import (
	"github.com/julienschmidt/httprouter"
)

// Routes ...
func Routes(router *httprouter.Router) {
	router.HandlerFunc("GET", "/", sampleGETRequest)
}
