package main

import (
	"github.com/example/my-golang-project/pkg/nes"
)

func main() {
	//filePath := "roms/Legend of Zelda, The (USA) (Rev A).nes"

	filePath := "roms/test_cpu_exec_space_apu.nes"
	//filePath := "roms/official_only.nes"

	// Run the example that demonstrates the memory system
	nes.RunExample(filePath)
}
