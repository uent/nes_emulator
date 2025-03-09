// Example of how to use the NES emulator with memory system
package nes

import "fmt"

// RunExample shows how to initialize and use the NES system
func RunExample(filePath string) {
	// Read the NES ROM file
	header, prgROM, err := ReadNESFile(filePath)
	if err != nil {
		fmt.Printf("Error reading NES file: %v\n", err)
		return
	}

	// Print ROM information
	fmt.Println("NES ROM Header Information:")
	fmt.Println(header.String())

	fmt.Println("\nPRG ROM Information:")
	fmt.Println(prgROM.String())

	// Create a new NES instance
	nes := New()

	// Reset the NES components
	nes.Reset()

	// Load the ROM data
	nes.LoadROM(prgROM.Data)

	// Example of reading a value from memory
	address := uint16(0x8000) // Typical starting address for PRG ROM
	value := nes.Memory.Read(address)
	fmt.Printf("\nValue at address 0x%04X: 0x%02X\n", address, value)

	// Set up the CPU to use our memory system
	nes.CPU.SetMemory(nes.Memory)

	// Now the CPU can read from and write to memory through the proper memory interface
	// The memory system handles the appropriate memory mapping and mirroring

	fmt.Println("\nNES system initialized successfully with memory system.")

	nes.Run()
}
