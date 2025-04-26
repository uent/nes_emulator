package cpu

import (
	"fmt"
)

// GetInstructionFunc returns the execution function for the given opcode
func GetInstructionFunc(opcode byte) CPUOperation {
	instruction, exists := InstructionTable[opcode]
	if !exists {
		//time.Sleep(1000 * time.Second)
		panic(fmt.Sprintf("Invalid opcode: %02X", opcode))
		// If the opcode is not found, return nil
		//return nil
	}
	return instruction.ExecuteFunc
}
