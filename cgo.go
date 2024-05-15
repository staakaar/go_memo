package main

/*
#include <stdlib.h>
#include <mylib.h>

// cgo: LDFLAGS -lmylib
// cgo: CFLAGS -l/usr/local/include
*/

import "C"
import "unsafe"

func main() {
	s := "index.dat"

	ptr := C.CString(s)
	defer C.free(unsafe.Pointer(ptr))
	C.mylib_update_data(ptr)
}
