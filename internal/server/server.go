package server

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Angstreminus/exchanger/internal/handler"
	"github.com/Angstreminus/exchanger/internal/service"
	"github.com/Angstreminus/exchanger/pkg/config"
	"github.com/Angstreminus/exchanger/pkg/logger"
	"go.uber.org/zap/zapcore"
)

type Server struct {
	Config *config.Config
	Logger *logger.Logger
	Router *http.ServeMux
}

func NewServer(config *config.Config, log *logger.Logger) *Server {
	return &Server{
		Logger: log,
		Config: config,
	}
}

func (srv *Server) MustRunWithGracefullShutdown() {
	srv.Router = http.NewServeMux()

	service := service.NewService(srv.Logger)
	handler := handler.NewHandler(service, srv.Logger)
	srv.Router.HandleFunc("POST /exchange", handler.CreateExchange)

	// Create a listener for graceful shutdown.
	listener, err := net.Listen("tcp", srv.Config.ServerHost+srv.Config.ServerPort)
	if err != nil {
		srv.Logger.Zap.Fatal("ERROR TO START SERVER", zapcore.Field{String: err.Error()})
	}

	// Run the HTTP server, using the listener we created for graceful shutdown.
	go func() {
		if err := http.Serve(listener, srv.Router); err != nil {
			srv.Logger.Zap.Fatal("ERROR TO START SERVER", zapcore.Field{String: err.Error()})
		}
	}()

	// Wait for a SIGHUP or SIGINT signal, which will trigger a graceful shutdown.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGHUP)
	<-signalChan

	// Trigger a graceful shutdown.
	if err := listener.Close(); err != nil {
		srv.Logger.Zap.Error("ERROR TO CLOSE SERVER", zapcore.Field{String: err.Error()})
	}
}
