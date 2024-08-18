package goevm

import (
	"math"

	"github.com/ethereum/go-ethereum/common"
)

func CalcMemoryGasCost(size uint64) uint64 {
	word := (size + 31) / 32
	cost := (uint64(math.Pow(float64(word), 2)))/512 + (3 * word)
	return cost
}

func toWordSize(size uint64) uint64 {
	return (size + 31) / 32
}

func getData(data []byte, start uint64, size uint64) []byte {
	length := uint64(len(data))
	if start > length {
		start = length
	}
	end := start + size
	if end > length {
		end = length
	}
	return common.RightPadBytes(data[start:end], int(size))
}

func CalcLogGasCost(topicCount, size, memExpansionCost uint64) uint64 {
	staticGas := uint64(375)
	return staticGas*topicCount + 8*size + memExpansionCost
}

func CalcSstoreGasCost(evm *EVM, slot int, newValue common.Hash) (gasCost uint64) {
	// Load the current value stored at the specified slot.
	currentValue, isWarm := evm.Storage.Get(slot)

	// Determine the access cost (cold vs warm)
	accessCost := uint64(2100)
	if isWarm {
		accessCost = 100
	}

	// If the current value is the same as the new value, it's a no-op.
	if currentValue == newValue {
		return accessCost
	}

	// Calculate the dynamic gas cost
	var dynamicCost uint64
	if currentValue == (common.Hash{}) && newValue != (common.Hash{}) {
		// Zero to non-zero
		dynamicCost = 20000
	} else if currentValue != (common.Hash{}) && newValue == (common.Hash{}) {
		// Non-zero to zero
		dynamicCost = 2900
		evm.addRefund(4800)
	} else {
		// Non-zero to non-zero
		dynamicCost = 2900
	}

	return accessCost + dynamicCost
}
