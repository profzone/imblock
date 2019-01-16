package script

import (
	"testing"
	"math/big"
)

var s *Stack

func init() {
	s = newStack()
}

func TestPush(t *testing.T) {
	s.push(big.NewInt(123))
	s.push(big.NewInt(456))
	s.push(big.NewInt(789))

	if s.len() != 3 {
		t.Errorf("expected length=3, got %d", s.len())
	}

	s.Print()
}

func TestPushN(t *testing.T) {
	s.pushN(big.NewInt(1), big.NewInt(2), big.NewInt(3))

	if s.len() != 6 {
		t.Errorf("expected length=6, got %d", s.len())
	}

	s.Print()
}

func TestPop(t *testing.T) {
	ret := s.pop()
	v := ret.Int64()

	if v != 3 {
		t.Errorf("expected value=3, got %d", v)
	}

	if s.len() != 5 {
		t.Errorf("expected length=5, got %d", s.len())
	}

	s.Print()
}

// swap the n th(from top of stack) element with top element
func TestSwap(t *testing.T) {
	s.swap(s.len()-1)

	ret := s.peek()
	v := ret.Int64()

	if v != 123 {
		t.Errorf("expected value=123, got %d", v)
	}
	s.Print()
}

func TestBack(t *testing.T) {
	ret := s.Back(s.len()-1)
	v := ret.Int64()

	if v != 2 {
		t.Errorf("expected value=2, got %d", v)
	}
	s.Print()
}

func TestRequire(t *testing.T) {
	err := s.require(5)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	err = s.require(6)
	if err == nil {
		t.Errorf("expected err, got nil")
	}
	t.Log(err)
}
