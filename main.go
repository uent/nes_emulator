package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/example/my-golang-project/pkg/debug"
	"github.com/example/my-golang-project/pkg/nes"
)

func main() {
	// Command line flags
	debugMode := flag.Bool("debug", false, "Run in debug mode with UI")
	//romPath := flag.String("rom", "roms/test_cpu_exec_space_apu.nes", "Path to ROM file")
	//romPath := flag.String("rom", "roms/official_only.nes", "Path to ROM file")
	romPath := flag.String("rom", "roms/Legend of Zelda, The (USA) (Rev A).nes", "Path to ROM file")
	//romPath := flag.String("rom", "roms/cpu_dummy_reads.nes", "Path to ROM file")

	flag.Parse()

	// Check if the ROM file exists
	if _, err := os.Stat(*romPath); os.IsNotExist(err) {
		fmt.Printf("Error: ROM file not found: %s\n", *romPath)
		return
	}

	// Create a new NES instance
	nesSystem := nes.New()

	// Read the NES ROM file
	header, prgROM, err := nes.ReadNESFile(*romPath)
	if err != nil {
		fmt.Printf("Error reading NES file: %v\n", err)
		return
	}

	// Print ROM information
	fmt.Println("NES ROM Header Information:")
	fmt.Println(header.String())

	fmt.Println("\nPRG ROM Information:")
	fmt.Println(prgROM.String())

	// Load the ROM data
	nesSystem.LoadROM(prgROM.Data)

	// Reset the NES components
	nesSystem.Reset()

	fmt.Println("\nNES system initialized successfully.")

	// Run in debug mode if flag is set
	if *debugMode {
		fmt.Println("Starting debug UI with graphics...")
		if err := debug.StartDebugger(nesSystem); err != nil {
			fmt.Printf("Error starting debug UI: %v\n", err)
		}
	} else {
		// Run in normal mode with game rendering
		fmt.Println("Running NES emulation with graphics...")
		if err := nes.StartGame(nesSystem); err != nil {
			fmt.Printf("Error running game: %v\n", err)
		}
	}
}
