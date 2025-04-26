# NES Emulator

A Go project implementing NES (Nintendo Entertainment System) emulation functionality with a graphical debugging UI.

## Project Structure

- `cmd/`: Contains main applications for this project
- `pkg/`: Contains code that's ok to be used by external applications
  - `apu/`: Audio Processing Unit emulation
  - `cpu/`: 6502 CPU emulation with full instruction set implementation
  - `debug/`: Debugging UI tools with real-time CPU state visualization
  - `memory/`: Memory system emulation with NES memory map implementation
  - `nes/`: NES ROM file handling and system integration
  - `ppu/`: Picture Processing Unit emulation
  - `utils/`: Helper utilities
- `internal/`: Contains private application and library code
- `roms/`: Contains NES ROM files for testing and running the emulator

## Features

- Complete 6502 CPU emulation with all instructions
- Memory system with proper NES memory mapping
- Basic PPU (Picture Processing Unit) implementation
- Basic APU (Audio Processing Unit) implementation
- ROM loading and parsing
- Interactive debugging UI with real-time CPU state visualization
- Support for various test ROMs

## How to Run

### Normal Execution

```bash
go run main.go
```

### Debug Mode

Run with the debug UI to monitor CPU state in real-time:

```bash
go run main.go -debug
```

Or use the provided script:

```bash
./run_debug.sh
```

### Specify ROM

You can specify a ROM file to load:

```bash
go run main.go -rom path/to/rom.nes
```

## CPU Debugger UI

The CPU Debugger UI provides a real-time view of the NES CPU state, including:

- Register values (A, X, Y, PC, SP)
- Processor status flags (NV-BDIZC)
- Instruction disassembly
- Memory view (Zero Page)
- Stack view

### Controls

- **Space**: Toggle pause/resume emulation
- **S**: Step forward one instruction (when in pause mode)

### UI Sections

1. **CPU Registers**: Shows current values of all CPU registers
2. **Flags**: Displays processor status flags (NV-BDIZC)
3. **Disassembly**: Shows recently executed instructions
4. **Memory**: Displays memory contents in the zero page
5. **Stack**: Shows the current state of the stack

## Building from Source

```bash
# Install dependencies
go get github.com/hajimehoshi/ebiten/v2
go get golang.org/x/image/font
go get golang.org/x/image/font/basicfont

# Build
go build -o nesemu main.go
```

You can also use the provided Makefile:

```bash
# Install dependencies and build
make all

# Just build
make build

# Run the emulator
make run
```

## Command Line Options

- `-debug`: Run with debug UI
- `-rom [path]`: Specify ROM file path (defaults to a test ROM)

## Included ROMs

The project includes several test ROMs:
- `cpu_dummy_reads.nes`: Tests CPU dummy read behavior
- `test_cpu_exec_space_apu.nes`: Tests CPU execution with APU
- `official_only.nes`: Official test ROM
- `Legend of Zelda, The (USA) (Rev A).nes`: The Legend of Zelda game ROM for testing

## Technical Implementation

The emulator implements the core components of the NES:

- **CPU**: MOS 6502 processor emulation with complete instruction set
- **Memory**: 2KB internal RAM with proper memory mapping
- **PPU**: Picture Processing Unit for graphics rendering
- **APU**: Audio Processing Unit for sound generation
- **ROM Loading**: Support for iNES format ROM files

## Development Status

This is a work-in-progress emulator with focus on accuracy and debugging capabilities. The CPU implementation is mostly complete, while PPU and APU implementations are still in early stages.
