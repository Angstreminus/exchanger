package service

import (
	"reflect"
	"sort"

	"github.com/Angstreminus/exchanger/internal/apperrors"
	"github.com/Angstreminus/exchanger/internal/dto"
	"github.com/Angstreminus/exchanger/pkg/logger"
)

type Service struct {
	Log *logger.Logger
}

func NewService(log *logger.Logger) *Service {
	return &Service{
		Log: log,
	}
}

// GetExchange implements the exchange creation usecase
func (serv Service) CreateExchange(req *dto.Request) ([][]int, apperrors.Apperror) {
	if err := ValidateData(req); err != nil {
		return nil, err
	}
	res := ChangeBanknotes(req)
	return res, nil
}

// ChangeBanknotes implements calculations and unification
func ChangeBanknotes(data *dto.Request) [][]int {
	result := change(data.Amount, data.Banknotes)
	uniqRes := make([][]int, 0)
	for _, v := range result {
		sort.Sort(sort.Reverse(sort.IntSlice(v)))
		exist := false
		for _, w := range uniqRes {
			if reflect.DeepEqual(w, v) {
				exist = true
			}
		}
		if !exist {
			uniqRes = append(uniqRes, v)
		}
	}
	return uniqRes
}

// change simple recursive func that implements forming the matrix of currency exchange
func change(amount int, banknotes []int) [][]int {
	if amount <= 0 || len(banknotes) == 0 {
		return [][]int{{}}
	}
	result := [][]int{}
	for i := range banknotes {
		if amount >= banknotes[i] {
			for _, change := range change(amount-banknotes[i], banknotes) {
				result = append(result, append([]int{banknotes[i]}, change...))
			}
		}
	}
	return result
}
