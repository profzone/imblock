package script

import (
	"testing"
	"math/big"
)

var interpreter = NewInterpreter(&VM{}, newBtcInstructionSet())
var code = []byte{
	byte(OP_CODE__DATA_1),
	byte(OP_CODE__DATA_1),
	byte(OP_CODE__ADD),
}

func TestRun(t *testing.T) {
	result, err := interpreter.Run(code, []byte{})
	if err != nil {
		t.Error(err)
	}
	t.Log(big.NewInt(0).SetBytes(result).Uint64())
}
