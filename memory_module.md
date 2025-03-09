# NES Memory Module

## Overview
This memory module provides a proper implementation of the NES memory system, which has a complex addressing scheme with memory mirroring, memory-mapped registers, and cartridge space.

## Memory Map of NES
The NES has a 16-bit address space (0x0000 - 0xFFFF) divided into several regions:

- **0x0000 - 0x1FFF**: 2KB internal RAM, mirrored 4 times
- **0x2000 - 0x3FFF**: PPU registers, mirrored every 8 bytes
- **0x4000 - 0x401F**: APU and I/O registers
- **0x4020 - 0xFFFF**: Cartridge space (PRG ROM, PRG RAM, and mapper registers)

## Key Features

### Memory Mirroring
The memory module handles appropriate mirroring of RAM and PPU registers.

### Memory Access Functions
- `Read(address uint16) byte`: Read a byte from any memory address
- `Write(address uint16, value byte)`: Write a byte to any memory address
- `ReadWord(address uint16) uint16`: Read a 16-bit word (considering NES is little-endian)
- `WriteWord(address uint16, value uint16)`: Write a 16-bit word

### ROM Loading
The module provides functionality to load the program ROM data into memory.

## Integration with CPU
The CPU has been updated to use the memory module through an interface, which allows for more proper memory access than the previous direct array access.

## Usage Example
The `example.go` file demonstrates how to use the memory module with the NES system:

1. Create a new NES instance
2. Reset all components
3. Load ROM data
4. Connect the memory module to the CPU
5. Read/write memory using the proper memory-mapped addressing

This implementation provides a solid foundation for accurate NES emulation by properly handling the complex memory mapping of the NES architecture.