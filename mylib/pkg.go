package mylib

// #cgo CFLAGS: -DPNG_DEBUG=1
// #cgo amd64 386 CFLAGS: -DX86=1
// #cgo LDFLAGS: -lpng
// #cgo pkg-config: png
// #include <png.h>
import "C"
