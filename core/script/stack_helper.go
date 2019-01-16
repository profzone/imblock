package script

import (
	"fmt"
)

const stackLimit = 1024

func makeStackValidationFunc(pop, push int) stackValidationFunc {
	return func(stack *Stack) error {
		if err := stack.require(pop); err != nil {
			return err
		}

		if stack.len()+push-pop > stackLimit {
			return fmt.Errorf("stack limit reached %d (%d)", stack.len(), stackLimit)
		}
		return nil
	}
}
