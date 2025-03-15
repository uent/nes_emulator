# NES Emulator

A Go project implementing NES (Nintendo Entertainment System) emulation functionality with a graphical debugging UI.

## Project Structure

- `cmd/`: Contains main applications for this project
- `pkg/`: Contains code that's ok to be used by external applications
  - `apu/`: Audio Processing Unit emulation
  - `cpu/`: 6502 CPU emulation
  - `debug/`: Debugging UI tools
  - `memory/`: Memory system emulation
  - `nes/`: NES ROM file handling
  - `ppu/`: Picture Processing Unit emulation
  - `utils/`: Helper utilities
- `internal/`: Contains private application and library code
- `roms/`: Contains NES ROM files

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

## CPU Debugger UI

The CPU Debugger UI provides a real-time view of the NES CPU state, including:

- Register values (A, X, Y, PC, SP)
- Processor status flags
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

# Build
go build -o nesemu main.go
```

## Command Line Options

- `-debug`: Run with debug UI
- `-rom [path]`: Specify ROM file path