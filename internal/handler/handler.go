package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Angstreminus/exchanger/internal/apperrors"
	"github.com/Angstreminus/exchanger/internal/dto"
	"github.com/Angstreminus/exchanger/pkg/logger"
)

type ExchangerService interface {
	CreateExchange(req *dto.Request) ([][]int, apperrors.Apperror)
}

func NewHandler(service ExchangerService, log *logger.Logger) *Hanlder {
	return &Hanlder{
		Service: service,
		Logger:  log,
	}
}

type Hanlder struct {
	Service ExchangerService
	Logger  *logger.Logger
}

func (exH *Hanlder) CreateExchange(w http.ResponseWriter, r *http.Request) {
	var ExReq dto.Request
	err := json.NewDecoder(r.Body).Decode(&ExReq)
	defer r.Body.Close()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	exCh, err := exH.Service.CreateExchange(&ExReq)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(exCh); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}