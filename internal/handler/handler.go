package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Angstreminus/exchanger/internal/apperrors"
	"github.com/Angstreminus/exchanger/internal/dto"
	"github.com/Angstreminus/exchanger/pkg/logger"
	"go.uber.org/zap/zapcore"
)

type ExchangerService interface {
	CreateExchange(req *dto.Request) ([][]int, apperrors.Apperror)
}

func NewHandler(service ExchangerService, log *logger.Logger) *Handler {
	return &Handler{
		Service: service,
		Logger:  log,
	}
}

type Handler struct {
	Service ExchangerService
	Logger  *logger.Logger
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  create_exchange
// @Tags         exchanges
// @Accept       json
// @Produce      json
// @Success      201  {object}  dto.Response
// @Failure      400 {object} httputil.HTTPError
// @Router       /http://localhost:8080/exchanges [post]

func (exH *Handler) CreateExchange(w http.ResponseWriter, r *http.Request) {
	var exReq dto.Request
	err := json.NewDecoder(r.Body).Decode(&exReq)
	defer r.Body.Close()
	if err != nil {
		exH.Logger.Zap.Error("ERROR TO DECODE REQUEST", zapcore.Field{String: err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}
	exCh, err := exH.Service.CreateExchange(&exReq)
	if err != nil {
		exH.Logger.Zap.Error(err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}
	exH.Logger.Zap.Info("EXCHANGE CREATED")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(&dto.Response{
		Exchanges: exCh,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
