package usecase

// Operation -
type Operation struct{}

// New -
func New() *Operation {
	return &Operation{}
}

// Compute -
func (op *Operation) Compute() (string, error) {
	return "Everything is ok", nil
}
