---
description: Repository Information Overview
alwaysApply: true
---

# Android PIN Bruteforcer Information

## Summary
The Android PIN Bruteforcer is a Go-based tool designed to demonstrate the use of the Android Open Accessory (AOA) Protocol to emulate HID events (key presses, mouse, stylus) via USB. Its primary purpose is to automate PIN entry on Android devices for security research and bruteforce demonstrations, specifically targeting the Sony Xperia Z1 Compact as a reference device.

## Structure
- **cmd/**: Contains the main application entry points.
    - **sony-xperia-z1-compact/**: Implementation specific to the Sony Xperia Z1 Compact.
- **pkg/**: Shared library packages.
    - **hid/**: HID report descriptors and touchscreen emulation logic.
    - **utils/**: Bitwise operations, file handling, and data structures.
- **pins/**: Contains lists of PIN combinations used for the bruteforce process.
- **media/**: Documentation assets (demo images).
- **vendor/**: Vendored Go dependencies.

## Language & Runtime
**Language**: Go  
**Version**: 1.23  
**Build System**: Makefile / Shell Scripts  
**Package Manager**: Go Modules

## Dependencies
**Main Dependencies**:
- `github.com/Tryanks/go-accessoryhid`: For Android Accessory HID communication.
- `github.com/google/gousb`: Go bindings for libusb.
- `libusb-1.0-0-dev`: System-level dependency for USB communication.

**Development Dependencies**:
- `golangci/golangci-lint`: Used via Docker for linting.

## Build & Installation
### Prerequisites
Install `libusb` development headers:
```bash
make install-libusb
# or
sudo apt-get install libusb-1.0-0-dev
```

### Build
To build the project using the Makefile:
```bash
make build
```
Alternatively, use the platform-specific scripts:
- **Windows**: `build.bat`
- **Linux/macOS**: `build.sh`

## Usage & Operations
The application requires root privileges to access USB devices directly.

**Run the tool**:
```bash
make run-sony
# or
sudo go run cmd/sony-xperia-z1-compact/main.go
```

**Key Configuration**:
- **pins/Pin 4 leangth.txt**: Default list of 4-digit PINs.
- **settings.txt**: Local configuration (if applicable).

## Testing & Validation
**Framework**: Go Test (standard library)  
**Status**: No test files (`*_test.go`) were found in the current repository, although a test target exists in the Makefile.

**Run Command**:
```bash
make test
```

**Linting**:
Requires Docker to run `golangci-lint`:
```bash
make lint
```
