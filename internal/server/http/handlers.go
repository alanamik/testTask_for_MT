package server

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrInternalServerError500 = errors.New("internal server error")
)

// A handler for test.
func (s *Server) HelloHandler(w http.ResponseWriter, _ *http.Request) {
	res := ResponseSuccess{
		StatusCode: 200,
		Cypher:     "Server is running",
	}
	respondWithSuccess(w, res)
}

func (s *Server) GetEncryptedStrHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, 500, "The method should be GET")
		return
	}
	var req GetEncryptedStrIn
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithError(w, 500, ErrInternalServerError500.Error())
		return
	}
	cypher, err := s.Encryption.Encrypt(req.Str, req.Algorithm)
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}
	s.Storage.AddInCache(r.Context(), req.Str, cypher)
	res := ResponseSuccess{
		StatusCode: 200,
		Cypher:     cypher,
	}
	respondWithSuccess(w, res)
}
