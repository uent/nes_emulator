package cpu

import "fmt"

// GetInstructionFunc returns the execution function for the given opcode
func GetInstructionFunc(opcode byte) CPUOperation {
	instruction, exists := InstructionTable[opcode]
	if !exists {
		panic(fmt.Sprintf("Invalid opcode: %02X", opcode))
		// If the opcode is not found, return nil
		//return nil
	}
	return instruction.ExecuteFunc
}
