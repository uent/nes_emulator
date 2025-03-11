// Package cpu provides 6502 CPU emulation functionality
package cpu

// AddressingMode represents different addressing modes for CPU instructions
type AddressingMode int

const (
	// Addressing modes
	Implied AddressingMode = iota
	Accumulator
	Immediate
	ZeroPage
	ZeroPageX
	ZeroPageY
	Relative
	Absolute
	AbsoluteX
	AbsoluteY
	Indirect
	IndirectX
	IndirectY
)

// Instruction represents a 6502 CPU instruction
type Instruction struct {
	Opcode      byte           // The instruction's opcode
	Mnemonic    string         // Instruction mnemonic (e.g., "LDA", "STA")
	Addressing  AddressingMode // Addressing mode
	Bytes       int            // Number of bytes the instruction takes
	Cycles      int            // Base number of cycles the instruction takes
	PageCycles  int            // Additional cycles if page boundary is crossed
	IsIllegal   bool           // Whether it's an illegal/unofficial opcode
	ExecuteFunc func(*CPU)     // Function to execute the instruction
}

// InstructionTable maps opcodes to their respective instructions
var InstructionTable = map[byte]Instruction{
	// BRK
	0x00: {0x00, "BRK", Implied, 1, 7, 0, false, BRK},

	// ORA
	0x01: {0x01, "ORA", IndirectX, 2, 6, 0, false, nil},
	0x05: {0x05, "ORA", ZeroPage, 2, 3, 0, false, nil},
	0x09: {0x09, "ORA", Immediate, 2, 2, 0, false, nil},
	0x0D: {0x0D, "ORA", Absolute, 3, 4, 0, false, nil},
	0x11: {0x11, "ORA", IndirectY, 2, 5, 1, false, nil},
	0x15: {0x15, "ORA", ZeroPageX, 2, 4, 0, false, nil},
	0x19: {0x19, "ORA", AbsoluteY, 3, 4, 1, false, nil},
	0x1D: {0x1D, "ORA", AbsoluteX, 3, 4, 1, false, nil},

	// ASL
	0x06: {0x06, "ASL", ZeroPage, 2, 5, 0, false, nil},
	0x0A: {0x0A, "ASL", Accumulator, 1, 2, 0, false, nil},
	0x0E: {0x0E, "ASL", Absolute, 3, 6, 0, false, nil},
	0x16: {0x16, "ASL", ZeroPageX, 2, 6, 0, false, nil},
	0x1E: {0x1E, "ASL", AbsoluteX, 3, 7, 0, false, nil},

	// PHP
	0x08: {0x08, "PHP", Implied, 1, 3, 0, false, nil},

	// BPL
	0x10: {0x10, "BPL", Relative, 2, 2, 2, false, nil},

	// CLC
	0x18: {0x18, "CLC", Implied, 1, 2, 0, false, nil},

	// JSR
	0x20: {0x20, "JSR", Absolute, 3, 6, 0, false, nil},

	// AND
	0x21: {0x21, "AND", IndirectX, 2, 6, 0, false, nil},
	0x25: {0x25, "AND", ZeroPage, 2, 3, 0, false, nil},
	0x29: {0x29, "AND", Immediate, 2, 2, 0, false, nil},
	0x2D: {0x2D, "AND", Absolute, 3, 4, 0, false, nil},
	0x31: {0x31, "AND", IndirectY, 2, 5, 1, false, nil},
	0x35: {0x35, "AND", ZeroPageX, 2, 4, 0, false, nil},
	0x39: {0x39, "AND", AbsoluteY, 3, 4, 1, false, nil},
	0x3D: {0x3D, "AND", AbsoluteX, 3, 4, 1, false, nil},

	// BIT
	0x24: {0x24, "BIT", ZeroPage, 2, 3, 0, false, nil},
	0x2C: {0x2C, "BIT", Absolute, 3, 4, 0, false, nil},

	// ROL
	0x26: {0x26, "ROL", ZeroPage, 2, 5, 0, false, nil},
	0x2A: {0x2A, "ROL", Accumulator, 1, 2, 0, false, nil},
	0x2E: {0x2E, "ROL", Absolute, 3, 6, 0, false, nil},
	0x36: {0x36, "ROL", ZeroPageX, 2, 6, 0, false, nil},
	0x3E: {0x3E, "ROL", AbsoluteX, 3, 7, 0, false, nil},

	// PLP
	0x28: {0x28, "PLP", Implied, 1, 4, 0, false, nil},

	// BMI
	0x30: {0x30, "BMI", Relative, 2, 2, 2, false, nil},

	// SEC
	0x38: {0x38, "SEC", Implied, 1, 2, 0, false, nil},

	// RTI
	0x40: {0x40, "RTI", Implied, 1, 6, 0, false, nil},

	// EOR
	0x41: {0x41, "EOR", IndirectX, 2, 6, 0, false, nil},
	0x45: {0x45, "EOR", ZeroPage, 2, 3, 0, false, nil},
	0x49: {0x49, "EOR", Immediate, 2, 2, 0, false, nil},
	0x4D: {0x4D, "EOR", Absolute, 3, 4, 0, false, nil},
	0x51: {0x51, "EOR", IndirectY, 2, 5, 1, false, nil},
	0x55: {0x55, "EOR", ZeroPageX, 2, 4, 0, false, nil},
	0x59: {0x59, "EOR", AbsoluteY, 3, 4, 1, false, nil},
	0x5D: {0x5D, "EOR", AbsoluteX, 3, 4, 1, false, nil},

	// LSR
	0x46: {0x46, "LSR", ZeroPage, 2, 5, 0, false, nil},
	0x4A: {0x4A, "LSR", Accumulator, 1, 2, 0, false, nil},
	0x4E: {0x4E, "LSR", Absolute, 3, 6, 0, false, nil},
	0x56: {0x56, "LSR", ZeroPageX, 2, 6, 0, false, nil},
	0x5E: {0x5E, "LSR", AbsoluteX, 3, 7, 0, false, nil},

	// PHA
	0x48: {0x48, "PHA", Implied, 1, 3, 0, false, nil},

	// JMP
	0x4C: {0x4C, "JMP", Absolute, 3, 3, 0, false, nil},
	0x6C: {0x6C, "JMP", Indirect, 3, 5, 0, false, nil},

	// BVC
	0x50: {0x50, "BVC", Relative, 2, 2, 2, false, nil},

	// CLI
	0x58: {0x58, "CLI", Implied, 1, 2, 0, false, nil},

	// RTS
	0x60: {0x60, "RTS", Implied, 1, 6, 0, false, nil},

	// ADC
	0x61: {0x61, "ADC", IndirectX, 2, 6, 0, false, nil},
	0x65: {0x65, "ADC", ZeroPage, 2, 3, 0, false, nil},
	0x69: {0x69, "ADC", Immediate, 2, 2, 0, false, nil},
	0x6D: {0x6D, "ADC", Absolute, 3, 4, 0, false, nil},
	0x71: {0x71, "ADC", IndirectY, 2, 5, 1, false, nil},
	0x75: {0x75, "ADC", ZeroPageX, 2, 4, 0, false, nil},
	0x79: {0x79, "ADC", AbsoluteY, 3, 4, 1, false, nil},
	0x7D: {0x7D, "ADC", AbsoluteX, 3, 4, 1, false, nil},

	// ROR
	0x66: {0x66, "ROR", ZeroPage, 2, 5, 0, false, nil},
	0x6A: {0x6A, "ROR", Accumulator, 1, 2, 0, false, nil},
	0x6E: {0x6E, "ROR", Absolute, 3, 6, 0, false, nil},
	0x76: {0x76, "ROR", ZeroPageX, 2, 6, 0, false, nil},
	0x7E: {0x7E, "ROR", AbsoluteX, 3, 7, 0, false, nil},

	// PLA
	0x68: {0x68, "PLA", Implied, 1, 4, 0, false, nil},

	// BVS
	0x70: {0x70, "BVS", Relative, 2, 2, 2, false, nil},

	// SEI
	0x78: {0x78, "SEI", Implied, 1, 2, 0, false, nil},

	// STA
	0x81: {0x81, "STA", IndirectX, 2, 6, 0, false, nil},
	0x85: {0x85, "STA", ZeroPage, 2, 3, 0, false, nil},
	0x8D: {0x8D, "STA", Absolute, 3, 4, 0, false, nil},
	0x91: {0x91, "STA", IndirectY, 2, 6, 0, false, nil},
	0x95: {0x95, "STA", ZeroPageX, 2, 4, 0, false, nil},
	0x99: {0x99, "STA", AbsoluteY, 3, 5, 0, false, nil},
	0x9D: {0x9D, "STA", AbsoluteX, 3, 5, 0, false, nil},

	// STY
	0x84: {0x84, "STY", ZeroPage, 2, 3, 0, false, nil},
	0x8C: {0x8C, "STY", Absolute, 3, 4, 0, false, nil},
	0x94: {0x94, "STY", ZeroPageX, 2, 4, 0, false, nil},

	// STX
	0x86: {0x86, "STX", ZeroPage, 2, 3, 0, false, nil},
	0x8E: {0x8E, "STX", Absolute, 3, 4, 0, false, nil},
	0x96: {0x96, "STX", ZeroPageY, 2, 4, 0, false, nil},

	// DEY
	0x88: {0x88, "DEY", Implied, 1, 2, 0, false, nil},

	// TXA
	0x8A: {0x8A, "TXA", Implied, 1, 2, 0, false, nil},

	// BCC
	0x90: {0x90, "BCC", Relative, 2, 2, 2, false, nil},

	// TYA
	0x98: {0x98, "TYA", Implied, 1, 2, 0, false, nil},

	// TXS
	0x9A: {0x9A, "TXS", Implied, 1, 2, 0, false, nil},

	// LDY
	0xA0: {0xA0, "LDY", Immediate, 2, 2, 0, false, nil},
	0xA4: {0xA4, "LDY", ZeroPage, 2, 3, 0, false, nil},
	0xAC: {0xAC, "LDY", Absolute, 3, 4, 0, false, nil},
	0xB4: {0xB4, "LDY", ZeroPageX, 2, 4, 0, false, nil},
	0xBC: {0xBC, "LDY", AbsoluteX, 3, 4, 1, false, nil},

	// LDA
	0xA1: {0xA1, "LDA", IndirectX, 2, 6, 0, false, nil},
	0xA5: {0xA5, "LDA", ZeroPage, 2, 3, 0, false, nil},
	0xA9: {0xA9, "LDA", Immediate, 2, 2, 0, false, nil},
	0xAD: {0xAD, "LDA", Absolute, 3, 4, 0, false, nil},
	0xB1: {0xB1, "LDA", IndirectY, 2, 5, 1, false, nil},
	0xB5: {0xB5, "LDA", ZeroPageX, 2, 4, 0, false, nil},
	0xB9: {0xB9, "LDA", AbsoluteY, 3, 4, 1, false, nil},
	0xBD: {0xBD, "LDA", AbsoluteX, 3, 4, 1, false, nil},

	// LDX
	0xA2: {0xA2, "LDX", Immediate, 2, 2, 0, false, nil},
	0xA6: {0xA6, "LDX", ZeroPage, 2, 3, 0, false, nil},
	0xAE: {0xAE, "LDX", Absolute, 3, 4, 0, false, nil},
	0xB6: {0xB6, "LDX", ZeroPageY, 2, 4, 0, false, nil},
	0xBE: {0xBE, "LDX", AbsoluteY, 3, 4, 1, false, nil},

	// TAY
	0xA8: {0xA8, "TAY", Implied, 1, 2, 0, false, nil},

	// TAX
	0xAA: {0xAA, "TAX", Implied, 1, 2, 0, false, nil},

	// BCS
	0xB0: {0xB0, "BCS", Relative, 2, 2, 2, false, nil},

	// CLV
	0xB8: {0xB8, "CLV", Implied, 1, 2, 0, false, nil},

	// TSX
	0xBA: {0xBA, "TSX", Implied, 1, 2, 0, false, nil},

	// CPY
	0xC0: {0xC0, "CPY", Immediate, 2, 2, 0, false, nil},
	0xC4: {0xC4, "CPY", ZeroPage, 2, 3, 0, false, nil},
	0xCC: {0xCC, "CPY", Absolute, 3, 4, 0, false, nil},

	// CMP
	0xC1: {0xC1, "CMP", IndirectX, 2, 6, 0, false, nil},
	0xC5: {0xC5, "CMP", ZeroPage, 2, 3, 0, false, nil},
	0xC9: {0xC9, "CMP", Immediate, 2, 2, 0, false, nil},
	0xCD: {0xCD, "CMP", Absolute, 3, 4, 0, false, nil},
	0xD1: {0xD1, "CMP", IndirectY, 2, 5, 1, false, nil},
	0xD5: {0xD5, "CMP", ZeroPageX, 2, 4, 0, false, nil},
	0xD9: {0xD9, "CMP", AbsoluteY, 3, 4, 1, false, nil},
	0xDD: {0xDD, "CMP", AbsoluteX, 3, 4, 1, false, nil},

	// DEC
	0xC6: {0xC6, "DEC", ZeroPage, 2, 5, 0, false, nil},
	0xCE: {0xCE, "DEC", Absolute, 3, 6, 0, false, nil},
	0xD6: {0xD6, "DEC", ZeroPageX, 2, 6, 0, false, nil},
	0xDE: {0xDE, "DEC", AbsoluteX, 3, 7, 0, false, nil},

	// INY
	0xC8: {0xC8, "INY", Implied, 1, 2, 0, false, nil},

	// DEX
	0xCA: {0xCA, "DEX", Implied, 1, 2, 0, false, nil},

	// BNE
	0xD0: {0xD0, "BNE", Relative, 2, 2, 2, false, nil},

	// CLD
	0xD8: {0xD8, "CLD", Implied, 1, 2, 0, false, nil},

	// CPX
	0xE0: {0xE0, "CPX", Immediate, 2, 2, 0, false, nil},
	0xE4: {0xE4, "CPX", ZeroPage, 2, 3, 0, false, nil},
	0xEC: {0xEC, "CPX", Absolute, 3, 4, 0, false, nil},

	// SBC
	0xE1: {0xE1, "SBC", IndirectX, 2, 6, 0, false, nil},
	0xE5: {0xE5, "SBC", ZeroPage, 2, 3, 0, false, nil},
	0xE9: {0xE9, "SBC", Immediate, 2, 2, 0, false, nil},
	0xED: {0xED, "SBC", Absolute, 3, 4, 0, false, nil},
	0xF1: {0xF1, "SBC", IndirectY, 2, 5, 1, false, nil},
	0xF5: {0xF5, "SBC", ZeroPageX, 2, 4, 0, false, nil},
	0xF9: {0xF9, "SBC", AbsoluteY, 3, 4, 1, false, nil},
	0xFD: {0xFD, "SBC", AbsoluteX, 3, 4, 1, false, nil},

	// INC
	0xE6: {0xE6, "INC", ZeroPage, 2, 5, 0, false, nil},
	0xEE: {0xEE, "INC", Absolute, 3, 6, 0, false, nil},
	0xF6: {0xF6, "INC", ZeroPageX, 2, 6, 0, false, nil},
	0xFE: {0xFE, "INC", AbsoluteX, 3, 7, 0, false, nil},

	// INX
	0xE8: {0xE8, "INX", Implied, 1, 2, 0, false, nil},

	// NOP
	0xEA: {0xEA, "NOP", Implied, 1, 2, 0, false, nil},

	// BEQ
	0xF0: {0xF0, "BEQ", Relative, 2, 2, 2, false, nil},

	// SED
	0xF8: {0xF8, "SED", Implied, 1, 2, 0, false, nil},

	// Mark illegal opcodes
	0x02: {0x02, "ILLEGAL", Implied, 0, 0, 0, true, nil},
	0x03: {0x03, "ILLEGAL", Implied, 0, 0, 0, true, nil},
	0x04: {0x04, "ILLEGAL", Implied, 0, 0, 0, true, nil},
	0x07: {0x07, "ILLEGAL", Implied, 0, 0, 0, true, nil},
	0x0B: {0x0B, "ILLEGAL", Implied, 0, 0, 0, true, nil},
	0x0C: {0x0C, "ILLEGAL", Implied, 0, 0, 0, true, nil},
	0x0F: {0x0F, "ILLEGAL", Implied, 0, 0, 0, true, nil},
	0x12: {0x12, "ILLEGAL", Implied, 0, 0, 0, true, nil},
	0x13: {0x13, "ILLEGAL", Implied, 0, 0, 0, true, nil},
	0x14: {0x14, "ILLEGAL", Implied, 0, 0, 0, true, nil},
	0x17: {0x17, "ILLEGAL", Implied, 0, 0, 0, true, nil},
	0x1A: {0x1A, "ILLEGAL", Implied, 0, 0, 0, true, nil},
	0x1B: {0x1B, "ILLEGAL", Implied, 0, 0, 0, true, nil},
	0x1C: {0x1C, "ILLEGAL", Implied, 0, 0, 0, true, nil},
	0x1F: {0x1F, "ILLEGAL", Implied, 0, 0, 0, true, nil},
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

// AddressingModeToString converts an addressing mode to its string representation
func AddressingModeToString(mode AddressingMode) string {
	switch mode {
	case Implied:
		return "Implied"
	case Accumulator:
		return "Accumulator"
	case Immediate:
		return "Immediate"
	case ZeroPage:
		return "Zero Page"
	case ZeroPageX:
		return "Zero Page,X"
	case ZeroPageY:
		return "Zero Page,Y"
	case Relative:
		return "Relative"
	case Absolute:
		return "Absolute"
	case AbsoluteX:
		return "Absolute,X"
	case AbsoluteY:
		return "Absolute,Y"
	case Indirect:
		return "Indirect"
	case IndirectX:
		return "(Indirect,X)"
	case IndirectY:
		return "(Indirect),Y"
	default:
		return "Unknown"
	}
}
