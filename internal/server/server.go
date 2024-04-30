package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	srv.Logger.Zap.Info("ROUNTES INITIALIZED")
	srv.Router.HandleFunc("POST /exchange", handler.CreateExchange)
	fmt.Println()
	serv := &http.Server{
		Addr:    srv.Config.ServerHost + srv.Config.ServerPort,
		Handler: srv.Router,
	}

	go func() {
		if err := serv.ListenAndServe(); err != nil {
			srv.Logger.Zap.Error("Error to run server", zapcore.Field{String: err.Error()})
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := serv.Shutdown(ctx); err != nil {
		srv.Logger.Zap.Error("Error to shutdown server")
	} else {
		// here is no connections
		srv.Logger.Zap.Info("All connections gracefully closed")
	}
}
