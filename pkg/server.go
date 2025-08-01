package todo

import (
	"context"
	"net/http"
	"time"
)

// Структура для запуска http-сервера

type Server struct{
	httpServer *http.Server // указатель на структуру
}

// Запуск сервера
func (s *Server) Run(port string, handler http.Handler) error{
	s.httpServer = &http.Server{
		Addr: ":" + port,
// Handler присваивается 1 раз!
// gin.Engine.ServeHTTP(w, req) смотрит, что это POST /sign-up, и вызывает уже h.signUp()
		Handler: handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout: 10 + time.Second,
		WriteTimeout: 10 + time.Second,
	}

	// Создает подключение по порты Addr
	return s.httpServer.ListenAndServe()
}

// Остановка сервера
func (s *Server) Shutdown(ctx context.Context) error{
	return s.httpServer.Shutdown(ctx)
}