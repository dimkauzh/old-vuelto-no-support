# Vuelto - Installation Guide

This guide provides step-by-step instructions for setting up a C compiler and SDL2 library on different operating systems.

## Table of Contents

1. [Windows](#windows)
2. [Mac](#mac)
3. [Linux](#linux)

## Windows

### 1. Install a C Compiler (MinGW-w64)

1. Download the MinGW-w64 installer from the official website: [MinGW-w64](https://mingw-w64.org/doku.php).

2. Run the installer and follow the on-screen instructions.

3. During installation, select the architecture (32 or 64-bit) and the threading model as needed.

4. Add the MinGW-w64 `bin` directory to your system PATH.

### 2. Install SDL2 on Windows

1. Download the SDL2 development libraries for Windows from the official website: [SDL2](https://www.libsdl.org/download-2.0.php).

2. Extract the downloaded archive.

3. Copy the contents of the `lib` folder into the MinGW-w64 `lib` directory.

4. Copy the contents of the `include` folder into the MinGW-w64 `include` directory.

5. Copy the SDL2.dll file from the `lib` folder to the directory where your compiled executable will be.

## Mac

### 1. Install Xcode Command Line Tools (if not already installed)

Open a terminal and run the following command:

```bash
xcode-select --install
```

Follow the on-screen instructions to complete the installation.

### 2. Install Homebrew (if not already installed)

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
```

3. Install a C Compiler (GCC)

```bash

brew install gcc

```

4. Install SDL2 on Mac

```bash

brew install sdl2

```

## Linux
### 1. Install a C Compiler (GCC)

On Debian-based systems (e.g., Ubuntu):

```bash

sudo apt-get update
sudo apt-get install build-essential
```

On Red Hat-based systems (e.g., Fedora):

```bash

sudo dnf install gcc
```

On Arch Linux:

```bash

sudo pacman -S gcc
```

### 2. Install SDL2 on Linux

On Debian-based systems:

```bash

sudo apt-get install libsdl2-dev
```

On Red Hat-based systems:

```bash

sudo dnf install SDL2-devel
```

On Arch Linux:

```bash

sudo pacman -S sdl2
```

## Conclusion

After following the steps for your respective operating system, you should have a working C compiler and SDL2 library installed. You can now proceed with your development and use SDL2 in your C projects.