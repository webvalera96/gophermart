package user

import "net/http"

type RegisterHandler struct{}

func (h RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("register"))
}
