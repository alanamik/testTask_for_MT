package server

import (
	"context"
	"fmt"
	"mt/internal/config"
	"net/http"
	"time"
)

type Encryption interface {
	Encrypt(str string, alg string) (string, error)
	EncryptSHA256(str string) (string, error)
	EncryptMD5(str string) (string, error)
}

type Storage interface {
	AddInCache(ctx context.Context, str string, cypher string) error
	GetFromCache(ctx context.Context, str string) (string, error)
}

type Server struct {
	Serv       *http.Server
	Encryption Encryption
	Storage    Storage
}

func New(e Encryption, s Storage, conf *config.Config) *Server {
	server := Server{
		Encryption: e,
		Storage:    s,
	}
	server.Serv = &http.Server{
		Addr:              conf.Service.Host + ":" + fmt.Sprint(conf.Service.Port),
		Handler:           server.routes(),
		ReadHeaderTimeout: 1 * time.Second,
		ReadTimeout:       2 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	return &server
}

func (s *Server) Start() error {
	if err := s.Serv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.Serv.Shutdown(ctx)
	if err != nil {
		fmt.Println("server shutdown error: " + err.Error())
		return err
	}
	fmt.Println("Server has stopped")

	return nil
}
