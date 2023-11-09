<h1 align="center">Vuelto</h1>
<h3 align="center">A small CGo Game Engine that uses a custom Software Renderer.</h3>

<p align="center">
  <a href="https://github.com/dimkauzh/vuelto"><img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/dimkauzh/vuelto"></a>
  <a href="https://github.com/dimkauzh/vuelto"><img alt="GitHub license" src="https://img.shields.io/github/license/dimkauzh/vuelto"></a>
  <a href="https://github.com/dimkauzh/vuelto"><img alt="CI Check" src="https://github.com/dimkauzh/vuelto/actions/workflows/ci_check.yml/badge.svg"></a>
  <a href="https://github.com/dimkauzh/vuelto"><img alt="Lines of code" src="https://tokei.rs/b1/github/dimkauzh/vuelto?category=lines"></a>
  <a href="https://goreportcard.com/report/github.com/dimkauzh/vuelto"><img alt="Report card" src="https://goreportcard.com/badge/github.com/dimkauzh/vuelto"></a>
</p>

> **Please note that vuelto is currently under high maintenance and is not production-ready.** The project is actively being developed and improved, which is why there is only a `main` branch available. Also this project is not intended to be used in production. It's just a hobby project.

## Building
### Prerequisites
You need to have the following tools installed:
 - [Go](https://golang.org/dl/)
 - [C compiler](README.md)
 - [SDL2](https://www.libsdl.org/)

We have a [installation guide here](INSTALLATION.md)

### Getting the Go package
Vertex is an ordinary go package, so you can get it using this command:
```bash
go get github.com/dimkauzh/vuelto@latest
```

## Example
We have a example located at [example/example.go](example/example1/example.go)

### Running the Example
1. Open a terminal.

2. Clone the repo
```bash
git clone https://github.com/dimkauzh/vuelto.git
```

3. Navigate to the root directory of the vuelto project.
```bash
cd vertex
```

4. To run the example, use the following command:
```bash
go run example/example.go
```

## Documentation
The documentation [is available here](https://pkg.go.dev/github.com/dimkauzh/vuelto#section-documentation). It will cover all of the api. We have the main documentation and the one that goes about the src folder which holds all the functions and methods. The main one is only to init and create a window

## License
This project is licensed under the GPLv3 License - see the LICENSE file for details.
