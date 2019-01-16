package script

import (
	"github.com/profzone/data_structures/algorithm"
	"github.com/sirupsen/logrus"
)

func opFalse(pc *uint64, code []byte, interpreter *Interpreter, stack *Stack) ([]byte, error) {
	integer := interpreter.GetStackElementPool().get()
	stack.push(integer.SetBytes([]byte{0x00}))

	return nil, nil
}

func op1Negate(pc *uint64, code []byte, interpreter *Interpreter, stack *Stack) ([]byte, error) {
	integer := interpreter.GetStackElementPool().get()
	stack.push(integer.SetBytes(scriptNum(-1).Bytes()))

	return nil, nil
}

func makeOpPush(size uint64) executionFunc {
	return func(pc *uint64, code []byte, interpreter *Interpreter, stack *Stack) ([]byte, error) {

		codeLen := len(code)

		// 检查code不要越界
		start := algorithm.IntMin(int64(codeLen), int64(*pc+1))
		end := algorithm.IntMin(int64(codeLen), start+int64(size))

		integer := interpreter.GetStackElementPool().get()
		stack.push(integer.SetBytes(algorithm.RightPadBytes(code[start:end], int(size))))

		*pc += size
		return nil, nil
	}
}

func makeOpPushData(size uint64) executionFunc {
	return func(pc *uint64, code []byte, interpreter *Interpreter, stack *Stack) ([]byte, error) {

		codeLen := len(code)

		// 检查code不要越界
		start := algorithm.IntMin(int64(codeLen), int64(*pc+1))
		end := algorithm.IntMin(int64(codeLen), start+int64(size))

		var length uint
		switch size {
		case 1:
			length = uint(code[end])
		case 2:
			length = uint(code[end])<<8 | uint(code[start])
		case 4:
			length = uint(code[end]<<24) | uint(code[end-1]<<16) | uint(code[start+1]<<8) | uint(code[start])
		default:
			logrus.Panicf("makeOpPushData only support [1, 2, 4], current %d", size)
		}

		// 检查code不要越界
		start = algorithm.IntMin(int64(codeLen), int64(end+1))
		end = algorithm.IntMin(int64(codeLen), start+int64(length))

		integer := interpreter.GetStackElementPool().get()
		stack.push(integer.SetBytes(algorithm.RightPadBytes(code[start:end], int(size))))

		*pc += size + uint64(length)
		return nil, nil
	}
}
