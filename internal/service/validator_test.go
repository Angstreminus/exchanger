package service

import (
	"testing"

	"github.com/Angstreminus/exchanger/internal/apperrors"
	"github.com/Angstreminus/exchanger/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestValidateData(t *testing.T) {
	tests := []struct {
		name    string
		dataset *dto.Request
		want    apperrors.Apperror
	}{
		{
			name: "Zero amount test",
			dataset: &dto.Request{
				Amount:    0,
				Banknotes: []int{100, 22},
			},
			want: apperrors.InvalidData{
				Message: "Amount must be greater than 0",
			},
		},
		{
			name: "Negative amount test",
			dataset: &dto.Request{
				Amount:    -333,
				Banknotes: []int{100, 22},
			},
			want: apperrors.InvalidData{
				Message: "Amount must be greater than 0",
			},
		},
		{
			name: "Zero banknote",
			dataset: &dto.Request{
				Amount:    33,
				Banknotes: []int{0, 0},
			},
			want: apperrors.InvalidData{
				Message: "Each of banknotes must be greater than 0",
			},
		},
		{
			name: "Negative banknote",
			dataset: &dto.Request{
				Amount:    33,
				Banknotes: []int{-92, 0},
			},
			want: apperrors.InvalidData{
				Message: "Each of banknotes must be greater than 0",
			},
		},
		{
			name: "Valid data",
			dataset: &dto.Request{
				Amount:    33,
				Banknotes: []int{55, 182},
			},
			want: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := ValidateData(test.dataset)
			assert.Equal(t, test.want, res)
		})
	}

}
