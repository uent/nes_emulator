// Package cpu provides 6502 CPU emulation functionality
package cpu

// AddressingMode represents different addressing modes for CPU instructions
type AddressingMode int

// Instruction represents a 6502 CPU instruction
type Instruction struct {
	Opcode      byte         // The instruction's opcode
	Mnemonic    string       // Instruction mnemonic (e.g., "LDA", "STA")
	IsIllegal   bool         // Whether it's an illegal/unofficial opcode
	ExecuteFunc CPUOperation // Function to execute the instruction
}

// InstructionTable maps opcodes to their respective instructions
var InstructionTable = map[byte]Instruction{
	// BRK
	0x00: {0x00, "BRK", false, BRK},

	// ORA
	0x01: {0x01, "ORA", false, nil},
	0x05: {0x05, "ORA", false, nil},
	0x09: {0x09, "ORA", false, nil},
	0x0D: {0x0D, "ORA", false, nil},
	0x11: {0x11, "ORA", false, nil},
	0x15: {0x15, "ORA", false, nil},
	0x19: {0x19, "ORA", false, nil},
	0x1D: {0x1D, "ORA", false, nil},

	// ASL
	0x06: {0x06, "ASL", false, nil},
	0x0A: {0x0A, "ASL", false, nil},
	0x0E: {0x0E, "ASL", false, nil},
	0x16: {0x16, "ASL", false, nil},
	0x1E: {0x1E, "ASL", false, nil},

	// PHP
	0x08: {0x08, "PHP", false, nil},

	// BPL
	0x10: {0x10, "BPL", false, BPLRelative},

	// CLC
	0x18: {0x18, "CLC", false, nil},

	// JSR
	0x20: {0x20, "JSR", false, JSRAbsolute},

	// AND
	0x21: {0x21, "AND", false, nil},
	0x25: {0x25, "AND", false, nil},
	0x29: {0x29, "AND", false, ANDImmediate},
	0x2D: {0x2D, "AND", false, nil},
	0x31: {0x31, "AND", false, nil},
	0x35: {0x35, "AND", false, nil},
	0x39: {0x39, "AND", false, nil},
	0x3D: {0x3D, "AND", false, nil},

	// BIT
	0x24: {0x24, "BIT", false, BITZero},
	0x2C: {0x2C, "BIT", false, BITAbsolute},

	// ROL
	0x26: {0x26, "ROL", false, nil},
	0x2A: {0x2A, "ROL", false, nil},
	0x2E: {0x2E, "ROL", false, nil},
	0x36: {0x36, "ROL", false, nil},
	0x3E: {0x3E, "ROL", false, nil},

	// PLP
	0x28: {0x28, "PLP", false, PLPImplied},

	// BMI
	0x30: {0x30, "BMI", false, nil},

	// SEC
	0x38: {0x38, "SEC", false, nil},

	// RTI
	0x40: {0x40, "RTI", false, nil},

	// EOR
	0x41: {0x41, "EOR", false, nil},
	0x45: {0x45, "EOR", false, nil},
	0x49: {0x49, "EOR", false, nil},
	0x4D: {0x4D, "EOR", false, nil},
	0x51: {0x51, "EOR", false, nil},
	0x55: {0x55, "EOR", false, nil},
	0x59: {0x59, "EOR", false, nil},
	0x5D: {0x5D, "EOR", false, nil},

	// LSR
	0x46: {0x46, "LSR", false, nil},
	0x4A: {0x4A, "LSR", false, LSRAccumulator},
	0x4E: {0x4E, "LSR", false, nil},
	0x56: {0x56, "LSR", false, nil},
	0x5E: {0x5E, "LSR", false, nil},

	// PHA
	0x48: {0x48, "PHA", false, PHAImplied},

	// JMP
	0x4C: {0x4C, "JMP", false, JMPAbsolute},
	0x6C: {0x6C, "JMP", false, JMPIndirect},

	// BVC
	0x50: {0x50, "BVC", false, nil},

	// CLI
	0x58: {0x58, "CLI", false, nil},

	// RTS
	0x60: {0x60, "RTS", false, RTSImplied},

	// ADC
	0x61: {0x61, "ADC", false, nil},
	0x65: {0x65, "ADC", false, nil},
	0x69: {0x69, "ADC", false, nil},
	0x6D: {0x6D, "ADC", false, nil},
	0x71: {0x71, "ADC", false, nil},
	0x75: {0x75, "ADC", false, ADCZeroPageX},
	0x79: {0x79, "ADC", false, nil},
	0x7D: {0x7D, "ADC", false, nil},

	// ROR
	0x66: {0x66, "ROR", false, nil},
	0x6A: {0x6A, "ROR", false, nil},
	0x6E: {0x6E, "ROR", false, nil},
	0x76: {0x76, "ROR", false, nil},
	0x7E: {0x7E, "ROR", false, nil},

	// PLA
	0x68: {0x68, "PLA", false, PLAImplied},

	// BVS
	0x70: {0x70, "BVS", false, nil},

	// SEI
	0x78: {0x78, "SEI", false, SEIImplied},

	// STA
	0x81: {0x81, "STA", false, nil},
	0x85: {0x85, "STA", false, STAZeroPage},
	0x8D: {0x8D, "STA", false, STAAbsolute},
	0x91: {0x91, "STA", false, STAIndirectY},
	0x95: {0x95, "STA", false, STAZeroPageX},
	0x99: {0x99, "STA", false, nil},
	0x9D: {0x9D, "STA", false, STAAbsoluteX},

	// STY
	0x84: {0x84, "STY", false, nil},
	0x8C: {0x8C, "STY", false, nil},
	0x94: {0x94, "STY", false, nil},

	// STX
	0x86: {0x86, "STX", false, STXZeroPage},
	0x8E: {0x8E, "STX", false, nil},
	0x96: {0x96, "STX", false, nil},

	// DEY
	0x88: {0x88, "DEY", false, DEYImplied},

	// TXA
	0x8A: {0x8A, "TXA", false, nil},

	// BCC
	0x90: {0x90, "BCC", false, nil},

	// TYA
	0x98: {0x98, "TYA", false, nil},

	// TXS
	0x9A: {0x9A, "TXS", false, TXSImplied},

	// LDY
	0xA0: {0xA0, "LDY", false, LDYImmediate},
	0xA4: {0xA4, "LDY", false, nil},
	0xAC: {0xAC, "LDY", false, nil},
	0xB4: {0xB4, "LDY", false, nil},
	0xBC: {0xBC, "LDY", false, nil},

	// LDA
	0xA1: {0xA1, "LDA", false, nil},
	0xA5: {0xA5, "LDA", false, nil},
	0xA9: {0xA9, "LDA", false, LDAImmediate},
	0xAD: {0xAD, "LDA", false, LDAAbsolute},
	0xB1: {0xB1, "LDA", false, nil},
	0xB5: {0xB5, "LDA", false, nil},
	0xB9: {0xB9, "LDA", false, nil},
	0xBD: {0xBD, "LDA", false, LDAAbsoluteX},

	// LDX
	0xA2: {0xA2, "LDX", false, LDXImmediate},
	0xA6: {0xA6, "LDX", false, nil},
	0xAE: {0xAE, "LDX", false, nil},
	0xB6: {0xB6, "LDX", false, nil},
	0xBE: {0xBE, "LDX", false, nil},

	// TAY
	0xA8: {0xA8, "TAY", false, nil},

	// TAX
	0xAA: {0xAA, "TAX", false, TAXImpplied},

	// BCS
	0xB0: {0xB0, "BCS", false, nil},

	// CLV
	0xB8: {0xB8, "CLV", false, nil},

	// TSX
	0xBA: {0xBA, "TSX", false, TSXImplied},

	// CPY
	0xC0: {0xC0, "CPY", false, nil},
	0xC4: {0xC4, "CPY", false, nil},
	0xCC: {0xCC, "CPY", false, nil},

	// CMP
	0xC1: {0xC1, "CMP", false, nil},
	0xC5: {0xC5, "CMP", false, nil},
	0xC9: {0xC9, "CMP", false, CMPImmediate},
	0xCD: {0xCD, "CMP", false, nil},
	0xD1: {0xD1, "CMP", false, CMPIndirectY},
	0xD5: {0xD5, "CMP", false, nil},
	0xD9: {0xD9, "CMP", false, nil},
	0xDD: {0xDD, "CMP", false, nil},

	// DEC
	0xC6: {0xC6, "DEC", false, DECZeroPage},
	0xCE: {0xCE, "DEC", false, nil},
	0xD6: {0xD6, "DEC", false, nil},
	0xDE: {0xDE, "DEC", false, nil},

	// INY
	0xC8: {0xC8, "INY", false, nil},

	// DEX
	0xCA: {0xCA, "DEX", false, DEXImplied},

	// BNE
	0xD0: {0xD0, "BNE", false, BNERelative},

	// CLD
	0xD8: {0xD8, "CLD", false, CLDImplied},

	// CPX
	0xE0: {0xE0, "CPX", false, nil},
	0xE4: {0xE4, "CPX", false, CPXZeroPage},
	0xEC: {0xEC, "CPX", false, nil},

	// SBC
	0xE1: {0xE1, "SBC", false, nil},
	0xE5: {0xE5, "SBC", false, nil},
	0xE9: {0xE9, "SBC", false, nil},
	0xED: {0xED, "SBC", false, nil},
	0xF1: {0xF1, "SBC", false, nil},
	0xF5: {0xF5, "SBC", false, nil},
	0xF9: {0xF9, "SBC", false, nil},
	0xFD: {0xFD, "SBC", false, nil},

	// INC
	0xE6: {0xE6, "INC", false, INCZeroPage},
	0xEE: {0xEE, "INC", false, nil},
	0xF6: {0xF6, "INC", false, nil},
	0xFE: {0xFE, "INC", false, nil},

	// INX
	0xE8: {0xE8, "INX", false, INXImplied},

	// NOP
	0xEA: {0xEA, "NOP", false, NOPImplied},

	// BEQ
	0xF0: {0xF0, "BEQ", false, BEQRelative},

	// SED
	0xF8: {0xF8, "SED", false, nil},

	// Mark illegal opcodes
	0x02: {0x02, "ILLEGAL", true, nil},
	0x03: {0x03, "ILLEGAL", true, nil},
	0x04: {0x04, "ILLEGAL", true, nil},
	0x07: {0x07, "ILLEGAL", true, nil},
	0x0B: {0x0B, "ILLEGAL", true, nil},
	0x0C: {0x0C, "ILLEGAL", true, nil},
	0x0F: {0x0F, "ILLEGAL", true, nil},
	0x12: {0x12, "ILLEGAL", true, nil},
	0x13: {0x13, "ILLEGAL", true, nil},
	0x14: {0x14, "ILLEGAL", true, nil},
	0x17: {0x17, "ILLEGAL", true, nil},
	0x1A: {0x1A, "ILLEGAL", true, nil},
	0x1B: {0x1B, "ILLEGAL", true, nil},
	0x1C: {0x1C, "ILLEGAL", true, nil},
	0x1F: {0x1F, "ILLEGAL", true, nil},
	// Continue with the rest of illegal opcodes...
	// Add all illegal opcodes up to 0xFF
}

// GetInstruction returns the instruction information for the given opcode
func GetInstruction(opcode byte) Instruction {
	instruction, exists := InstructionTable[opcode]
	if !exists {
		// If the opcode is not found, it's an illegal opcode
		return Instruction{
			Opcode:    opcode,
			Mnemonic:  "ILLEGAL",
			IsIllegal: true,
		}
	}
	return instruction
}
