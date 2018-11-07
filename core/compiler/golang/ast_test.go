package golang

import "testing"

func TestParse(t *testing.T) {
	err := Parse("./example.ast")
	if err != nil {
		t.Error(err)
	}
}
