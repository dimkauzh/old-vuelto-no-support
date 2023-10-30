<h1 align="center">Vuelto</h1>
<h3 align="center">A small CGo Game Engine that uses a custom Software Renderer.</h3>

<p align="center">
  <a href="https://github.com/dimkauzh/vuelto"><img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/dimkauzh/vuelto"></a>
  <a href="https://github.com/dimkauzh/vuelto"><img alt="GitHub license" src="https://img.shields.io/github/license/dimkauzh/vuelto"></a>
  <a href="https://github.com/dimkauzh/vuelto"><img alt="CI Check" src="https://github.com/dimkauzh/vuelto/actions/workflows/ci_check.yml/badge.svg"></a>
  <a href="https://github.com/dimkauzh/vuelto"><img alt="Lines of code" src="https://tokei.rs/b1/github/dimkauzh/vuelto?category=lines"></a>
  <a href="https://goreportcard.com/report/github.com/dimkauzh/vuelto"><img alt="Report card" src="https://goreportcard.com/badge/github.com/dimkauzh/vuelto"></a>
</p>

> **Please note that vuelto is currently under high maintenance and is not production-ready.** The project is actively being developed and improved, which is why there is only a `main` branch available. Also this project is not intended to be used in production. It is just a hobby project.

## Building
### Prerequisites
You need to have the following tools installed:
 - [Go](https://golang.org/dl/)
 - [C compiler](https://developer.fyne.io/started/)
 - [SDL2](https://www.libsdl.org/)

Without these tools installed you can't use this engine.


> **Warning: Our tool only has been tested on MacOS and one environment. If you have any issues, please a Github Issue**

### Getting the Go package
Vertex is an ordinary go package, so you can get it using this command:
```bash
go get github.com/dimkauzh/vuelto@latest
```
## Example
We have a example located at [example/example.go](https://github.com/dimkauzh/vuelto/blob/main/example/example.go)

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

## License
This project is licensed under the GPLv3 License - see the LICENSE file for details.
