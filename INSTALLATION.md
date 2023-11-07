# Installation
## Windows
There is no current support for windows, but you can use WSL to run the program. You can try to install SDL2 by yourself, it will be then linked automatically using the `-lSDL2` flag.

## MacOS
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
Install gcc using apt:
```bash
sudo apt install gcc
```

Install sdl2 using apt:
```bash
sudo apt install sdl2
```