package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server //Абстрация над структурой сервер стандартного пакета http, и имеет метод запуска и остановки
}

func (s *Server) Run(port string) error {
	s.httpServer = &http.Server{Addr: ":" + port,
		MaxHeaderBytes: 1 << 28, //1MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second, //10 секунд
	}
	return s.httpServer.ListenAndServe() //метод под капотом запускает бесконечный цикл for,и слушает все входящие запросы для обработки
}

/*Метод, необходимый для остановки сервера*/
func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
