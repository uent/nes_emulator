package main

import (
	"github.com/example/my-golang-project/pkg/nes"
)

func main() {
	filePath := "roms/official_only.nes"
	
	// Run the example that demonstrates the memory system
	nes.RunExample(filePath)
}
