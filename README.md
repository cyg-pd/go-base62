# go-base62

[![tag](https://img.shields.io/github/tag/cyg-pd/go-base62.svg)](https://github.com/cyg-pd/go-base62/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-%23007d9c)
[![GoDoc](https://godoc.org/github.com/cyg-pd/go-base62?status.svg)](https://pkg.go.dev/github.com/cyg-pd/go-base62)
![Build Status](https://github.com/cyg-pd/go-base62/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/cyg-pd/go-base62)](https://goreportcard.com/report/github.com/cyg-pd/go-base62)
[![Coverage](https://img.shields.io/codecov/c/github/cyg-pd/go-base62)](https://codecov.io/gh/cyg-pd/go-base62)
[![Contributors](https://img.shields.io/github/contributors/cyg-pd/go-base62)](https://github.com/cyg-pd/go-base62/graphs/contributors)
[![License](https://img.shields.io/github/license/cyg-pd/go-base62)](./LICENSE)

## 🚀 Install

```sh
go get github.com/cyg-pd/go-base62@v1
```

This library is v1 and follows SemVer strictly.

No breaking changes will be made to exported APIs before v2.0.0.

This library has no dependencies outside the Go standard library.

## 💡 Usage

You can import `base62` using:

```go
import (
    "log/slog"

    "github.com/cyg-pd/go-base62"
)

func main() {
    s := base62.Encode("9e3635adb80a4e138220de9f6d12d1cb", 16)
    slog.Info(s) // 4OxyOCWOdblodfDYb7qMyf
}
```
