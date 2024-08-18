package goevm

type Memory struct {
	data []byte
}

func (mem *Memory) Access(offset, size uint64) (cpy []byte) {
	if size == 0 {
		return nil
	}
	cpy = make([]byte, offset+size)
	copy(cpy, mem.data[offset:offset+size])
	return cpy
}

func (mem *Memory) Load(offset uint64) []byte {
	return mem.Access(offset, 32)
}

func (mem *Memory) Store(offset uint64, value []byte) (expansionCost uint64) {

	currentSize := uint64(mem.Len())
	currentCost := CalcMemoryGasCost(currentSize)
	newSize := offset + uint64(len(value))

	if currentSize == 0 {
		mem.data = make([]byte, 32)
		copy(mem.data, value[:])
		return CalcMemoryGasCost(32)
	}

	if currentSize < newSize {
		expansionSize := newSize - currentSize
		if expansionSize > 0 {
			mem.data = append(mem.data, make([]byte, expansionSize)...)
		}
		newCost := CalcMemoryGasCost(uint64(mem.Len()))
		expansionCost = newCost - currentCost
	}

	copy(mem.data[offset:newSize], value)
	return expansionCost
}

func (mem *Memory) Store32(offset uint64, value []byte) (expansionCost uint64) {
	// Current memory size and cost
	currentMemSize := uint64(mem.Len())
	currentCost := CalcMemoryGasCost(currentMemSize)
	newMemSize := offset + 32

	// Handle initial allocation separately
	if currentMemSize == 0 {
		mem.data = make([]byte, 32)
		copy(mem.data, value[:])
		return CalcMemoryGasCost(32)
	}

	if currentMemSize < newMemSize {
		expansionSize := newMemSize - currentMemSize
		if expansionSize > 0 {
			mem.data = append(mem.data, make([]byte, expansionSize)...)
		}
		newCost := CalcMemoryGasCost(uint64(mem.Len()))
		expansionCost = newCost - currentCost
	}

	copy(mem.data[offset:offset+32], value)
	return expansionCost
}

func (mem *Memory) Len() int {
	return len(mem.data)
}

func NewMemory() *Memory {
	return &Memory{data: make([]byte, 0)}
}
