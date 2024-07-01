package errorhandling

import "net/http"

func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "400 Bad Request", http.StatusBadRequest)
}
