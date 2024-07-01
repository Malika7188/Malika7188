package errorhandling

import "net/http"

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 not found", http.StatusNotFound)
}
