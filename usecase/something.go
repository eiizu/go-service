package usecase

import (
	"strings"
)

// Something -
type Something struct{}

// NewSomething -
func NewSomething() *Something {
	return &Something{}
}

// DoSomething -
func (op *Something) DoSomething(data string) (map[string]int, error) {
	resp := map[string]int{}

	arr := strings.Split(data, " ")
	for _, value := range arr {
		resp[value]++
	}

	return resp, nil
}
