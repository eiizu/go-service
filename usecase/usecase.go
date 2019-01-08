package usecase

import (
	"strings"
)

// Operation -
type Operation struct{}

// New -
func New() *Operation {
	return &Operation{}
}

// Compute -
func (op *Operation) Compute(data string) (map[string]int, error) {
	resp := map[string]int{}

	arr := strings.Split(data, " ")
	for _, value := range arr {
		resp[value]++
	}

	return resp, nil
}
