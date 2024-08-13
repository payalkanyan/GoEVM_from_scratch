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

func (mem *Memory) Store(offset uint64, value []byte) (expansioncost uint64) {

	currentSize := uint64(mem.Len())
	currentCost := CalcMemoryGasCost(currentSize)
	newSize := offset + uint64(len(value))

	if currentSize == 0 {
		mem.data = make([]byte, 32)
		copy(mem.data, value[:])
		return CalcMemoryGasCost(32)
	}
}

func Len(b []byte) {
	panic("unimplemented")
}

func (mem *Memory) Len() int {
	return len(mem.data)
}
