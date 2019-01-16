package script

import (
	"fmt"
)

type Interpreter struct {
	vm               *VM
	stackElementPool *stackElementPool
	returnData       []byte

	instructionSet [256]step
}

func (i *Interpreter) GetStackElementPool() *stackElementPool {
	return i.stackElementPool
}

func NewInterpreter(vm *VM, instructionSet [256]step) *Interpreter {
	// 检查指令集是否初始化，否则用默认代替
	if !instructionSet[OP_CODE__0].valid {
		instructionSet = newBtcInstructionSet()
	}

	return &Interpreter{
		vm:             vm,
		instructionSet: instructionSet,
	}
}

func (i *Interpreter) Run(code []byte, input []byte) ([]byte, error) {
	if i.stackElementPool == nil {
		i.stackElementPool = stackElementPools.get()
		defer func() {
			stackElementPools.put(i.stackElementPool)
			i.stackElementPool = nil
		}()
	}

	if len(code) == 0 {
		return nil, nil
	}

	var (
		opCode OpCode
		stack         = newStack()
		pc     uint64 = 0
	)

	defer func() {
		i.stackElementPool.put(stack.data...)
	}()

	for {
		if pc < uint64(len(code)) {
			opCode = OpCode(code[pc])
		} else {
			opCode = OP_CODE__UNKNOWN246
		}

		step := i.instructionSet[opCode]
		if !step.valid {
			return nil, fmt.Errorf("invalid opcode 0x%x", int(opCode))
		}
		if err := step.stackValidation(stack); err != nil {
			return nil, err
		}

		result, err := step.exec(&pc, code, i, stack)
		if step.returns {
			i.returnData = result
		}

		switch {
		case err != nil:
			return nil, err
		case step.halt:
			return result, nil
		case !step.jumps:
			pc++
		}
	}

	return nil, nil
}
