package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/hello", http.HandlerFunc(s.HelloHandler))
	mux.Handle("/encrypt", http.HandlerFunc(s.GetEncryptedStrHandler))

	return mux
}

func respondWithSuccess(w http.ResponseWriter, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	res := ResponseError{
		StatusCode: code,
		Error:      msg,
	}
	response, err := json.Marshal(res)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
