// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <wand/magick_wand.h>
*/
import "C"
import "os"
import "unsafe"

var (
	initialized bool
	dumb        *C.char
)

// Inicializes the MagickWand environment
func Initialize() {
	if initialized {
		return
	}
	dumb = C.CString(os.Args[0])
	//defer C.free(unsafe.Pointer(dumb))
	C.InitializeMagick(dumb)
	initialized = true
}

// Terminates the MagickWand environment
func Terminate() {
	if initialized {
		C.free(unsafe.Pointer(dumb))
		C.DestroyMagick()
		initialized = false
	}
}
