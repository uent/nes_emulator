# My Golang Project

A simple Go project implementing NES (Nintendo Entertainment System) emulation functionality, initialized with the standard Go project layout.

## Project Structure

- `cmd/`: Contains main applications for this project
- `pkg/`: Contains code that's ok to be used by external applications
  - `apu/`: Audio Processing Unit emulation
  - `cpu/`: 6502 CPU emulation
  - `nes/`: NES ROM file handling
  - `ppu/`: Picture Processing Unit emulation
  - `utils/`: Helper utilities
- `internal/`: Contains private application and library code
- `roms/`: Contains NES ROM files