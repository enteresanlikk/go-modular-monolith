package hello

import (
	"net/http"

	"github.com/enteresanlikk/go-modular-monolith/internal/common"
)

var Route = common.Route{Prefix: "/hello"}

func AddModule(mux *http.ServeMux) {
	mux.HandleFunc(Route.Create(""), HelloHandler)
}
