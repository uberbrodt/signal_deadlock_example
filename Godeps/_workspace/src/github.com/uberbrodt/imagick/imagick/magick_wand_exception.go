// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <wand/magick_wand.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type MagickWandException struct {
	kind        ExceptionType
	description string
}

func (mwe *MagickWandException) Error() string {
	return fmt.Sprintf("%s: %s", mwe.kind.String(), mwe.description)
}

// Returns the kind, reason and description of any error that occurs when using other methods in this API
func (mw *MagickWand) GetLastError() error {
	var et C.ExceptionType
	csdescription := C.MagickGetException(mw.mw, &et)
	defer relinquishMemory(unsafe.Pointer(csdescription))
	if ExceptionType(et) != EXCEPTION_UNDEFINED {
		//mw.clearException()
		return &MagickWandException{ExceptionType(C.int(et)), C.GoString(csdescription)}
	}
	return nil
}
