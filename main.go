package main

import (
	"fmt"
	"os"

	"github.com/example/my-golang-project/pkg/nes"
)

func main() {

	filePath := "roms/official_only.nes"

	header, prgROM, err := nes.ReadNESFile(filePath)
	if err != nil {
		fmt.Printf("Error reading NES file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("NES ROM Header Information:")
	fmt.Println(header.String())
	
	fmt.Println("\nPRG ROM Information:")
	fmt.Println(prgROM.String())

}
