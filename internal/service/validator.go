package service

import (
	"github.com/Angstreminus/exchanger/internal/apperrors"
	"github.com/Angstreminus/exchanger/internal/dto"
)

// ValidateData validates request data
func ValidateData(req *dto.Request) apperrors.Apperror {
	if req.Amount <= 0 {
		return apperrors.InvalidData{
			Message: "Amount must be greater than 0",
		}
	}
	for _, curr := range req.Banknotes {
		if curr <= 0 {
			return apperrors.InvalidData{
				Message: "Each of banknotes must be greater than 0",
			}
		}
	}
	return nil
}
