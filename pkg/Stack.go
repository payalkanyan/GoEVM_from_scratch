package goevm

import (
	"fmt"
	"strings"

	"github.com/holiman/uint256"
)

const MAX_STACK_SIZE = 1024

type Stack struct {
	data []uint256.Int
	// uint256.Int is an array of four uint64 elements, arranged in little-endian order. This means the least significant 64 bits are stored in the first element, and the most significant 64 bits are stored in the last element
}

func (st *Stack) push(value *uint256.Int) {
	if len(st.data) == MAX_STACK_SIZE {
		panic("Stack overflow")
	}
	st.data = append(st.data, *value)
}

func (st *Stack) pop() uint256.Int {
	if len(st.data) == 0 {
		panic("Stack underflow")
	}
	last := st.data[len(st.data)-1]
	st.data = st.data[:len(st.data)-1] //creates a new list where it includes all other elements except last.
	return last
}

func (st *Stack) peek() uint256.Int {
	if len(st.data) == 0 {
		panic("Stack underflow")
	}
	return st.data[len(st.data)-1]
}

func (st Stack) ToString() string {
	var d string
	if len(st.data) == 0 {
		d = "[]"
		return d
	}
	for i := len(st.data) - 1; i >= 0; i-- {
		if i == len(st.data)-1 {
			d = "["
		}
		d += fmt.Sprintf("%v, ", st.data[i].Hex())
		if i == 0 {
			d = strings.TrimRight(d, ", ")
			d += "]"
		}
	}
	return d
}

func NewStack() *Stack {
	return &Stack{data: make([]uint256.Int, 0)}
}
