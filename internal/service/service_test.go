package service

import (
	"testing"

	"github.com/Angstreminus/exchanger/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestCreateExchange(t *testing.T) {
	tests := []struct {
		name    string
		dataset *dto.Request
		want    [][]int
	}{
		{
			name: "Emptyness of banknotes",
			dataset: &dto.Request{
				Amount:    10,
				Banknotes: []int{},
			},
			want: [][]int{{}},
		},
		{
			name: "Negative amount",
			dataset: &dto.Request{
				Amount:    -10,
				Banknotes: []int{5000, 2000, 1000, 500, 200, 100, 50},
			},
			want: [][]int{{}},
		},
		{
			name: "Zero amount",
			dataset: &dto.Request{
				Amount:    0,
				Banknotes: []int{5000, 2000, 1000, 500, 200, 100, 50},
			},
			want: [][]int{{}},
		},
		{
			name: "Valid data",
			dataset: &dto.Request{
				Amount:    400,
				Banknotes: []int{5000, 2000, 1000, 500, 200, 100, 50},
			},
			want: [][]int{
				{200, 200},
				{200, 100, 100},
				{200, 100, 50, 50},
				{200, 50, 50, 50, 50},
				{100, 100, 100, 100},
				{100, 100, 100, 50, 50},
				{100, 100, 50, 50, 50, 50},
				{100, 50, 50, 50, 50, 50, 50},
				{50, 50, 50, 50, 50, 50, 50, 50},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := ChangeBanknotes(test.dataset)
			assert.Equal(t, test.want, res)
		})
	}

}
