#!/bin/bash

# Make sure required dependencies are installed
go get github.com/hajimehoshi/ebiten/v2
go get golang.org/x/image/font
go get golang.org/x/image/font/basicfont

# Build the application
echo "Building NES emulator..."
go build -o nesdbg main.go

# Run the application with debug UI enabled
echo "Starting NES debugger UI..."
./nesdbg -debug -rom roms/test_cpu_exec_space_apu.nes