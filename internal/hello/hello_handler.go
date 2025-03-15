package hello

import (
	"encoding/json"
	"net/http"

	"github.com/enteresanlikk/go-modular-monolith/internal/common"
)

type HelloResponse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func JsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	JsonResponse(w,
		common.SuccessDataResult(
			"Hello, World!",
			[]HelloResponse{
				{
					Name: "John Doe",
					Age:  20,
				},
				{
					Name: "Jane Doe",
					Age:  21,
				},
			},
		),
	)
}
