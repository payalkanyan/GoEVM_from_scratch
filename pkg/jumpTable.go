package goevm

// opcodes -> instruction functions
type JumpTable map[Opcode]func(*EVM)

func NewJumpTable() JumpTable {
	jumpTable := map[Opcode]func(*EVM){
		STOP:           stop,
		ADD:            add,
		MUL:            mul,
		SUB:            sub,
		DIV:            div,
		SDIV:           sdiv,
		MOD:            mod,
		SMOD:           smod,
		ADDMOD:         addmod,
		MULMOD:         mulmod,
		EXP:            exp,
		SIGNEXTEND:     signextend,
		LT:             lt,
		GT:             gt,
		SLT:            slt,
		SGT:            sgt,
		EQ:             eq,
		ISZERO:         iszero,
		AND:            and,
		OR:             or,
		XOR:            xor,
		NOT:            not,
		BYTE:           _byte,
		SHL:            shl,
		SHR:            shr,
		SAR:            sar,
		SHA3:           sha3,
		ADDRESS:        address,
		BALANCE:        balance,
		ORIGIN:         origin,
		CALLER:         caller,
		CALLVALUE:      callvalue,
		CALLDATALOAD:   calldataload,
		CALLDATASIZE:   calldatasize,
		CALLDATACOPY:   calldatacopy,
		CODESIZE:       codesize,
		CODECOPY:       codecopy,
		GASPRICE:       gasprice,
		GAS:            gas,
		EXTCODESIZE:    extcodesize,
		EXTCODECOPY:    extcodecopy,
		RETURNDATASIZE: returndatasize,
		RETURNDATACOPY: returndatacopy,
		BLOCKHASH:      blockhash,
		COINBASE:       coinbase,
		TIMESTAMP:      timestamp,
		NUMBER:         number,
		BASEFEE:        basefee,
		GASLIMIT:       gaslimit,
		CHAINID:        chainid,
		POP:            pop,
		PUSH0:          push0,
		MLOAD:          mload,
		MSTORE:         mstore,
		MSTORE8:        mstore8,
		MSIZE:          msize,
		MCOPY:          mcopy,
		SLOAD:          sload,
		SSTORE:         sstore,
		TLOAD:          tload,
		TSTORE:         tstore,
		JUMP:           jump,
		JUMPI:          jumpi,
		PC:             pc,
		JUMPDEST:       jumpdest,
		INVALID:        invalid,
		REVERT:         revert,
		RETURN:         _return,
		LOG0:           log0,
		LOG1:           log1,
		LOG2:           log2,
		LOG3:           log3,
		LOG4:           log4,
	}

	for i := 0; i <= 31; i++ {
		i := uint8(i)
		opcode := PUSH1 + Opcode(i)
		jumpTable[opcode] = generatePushFunc(i + 1)
	}

	for i := 0; i <= 15; i++ {
		i := uint8(i)
		opcode := DUP1 + Opcode(i)
		jumpTable[opcode] = generateDupNFunc(i + 1)
	}

	for i := 0; i <= 15; i++ {
		i := uint8(i)
		opcode := SWAP1 + Opcode(i)
		jumpTable[opcode] = generateSwapNFunc(i + 1)
	}

	return jumpTable
}
