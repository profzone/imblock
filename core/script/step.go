package script

type executionFunc func(pc *uint64, code []byte, interpreter *Interpreter, stack *Stack) ([]byte, error)
type stackValidationFunc func(stack *Stack) error

type step struct {
	exec            executionFunc
	stackValidation stackValidationFunc

	halt    bool
	valid   bool
	returns bool
	jumps   bool
}
