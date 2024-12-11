package libErrors

import "net/http"

func ReturnHttpError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
}
func ReturnHttpSucces(w http.ResponseWriter) {
}
