# Installation

## MacOS
We have prebuild binaries for intel macs, so you only need to install a c compiler.
Install xcode command line tools:
```bash
xcode-select install
```

## MacOS (M1)
We don't have prebuild binaries for M1 macs, so you need to install sdl2 using brew. But first a c compiler.
Install xcode command line tools:
```bash
xcode-select install
```

Install brew to install sdl2:
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

Install sdl2 with brew:
```bash
brew install sdl2
```

## Debian/Ubuntu (WSL and Linux)
We have prebuild binaries for linux, so you only need to install a c compiler.
Install gcc using apt:
```bash
sudo apt install gcc
```

Install sdl2 using apt:
```bash
sudo apt install sdl2
```

## Windows
We have prebuild binaries for windows, so you only need to install a c compiler.
We recomment [TDM-GCC](https://jmeubank.github.io/tdm-gcc/download/)

## Other and issues with sdl2
If you have any issues with sdl2 or you are using a different os, you can install sdl2 yourself, it will be linked automatically.