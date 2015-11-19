// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <wand/magick_wand.h>
*/
import "C"
import "unsafe"

// Convert a boolean to a integer
func b2i(boolean bool) C.uint {
	if boolean {
		return C.uint(1)
	}
	return C.uint(0)
}

func cStringArrayToStringSlice(p **C.char) []string {
	var strings []string
	q := uintptr(unsafe.Pointer(p))
	for {
		p = (**C.char)(unsafe.Pointer(q))
		if *p == nil {
			break
		}
		strings = append(strings, C.GoString(*p))
		q += unsafe.Sizeof(q)
	}
	return strings
}

func sizedCStringArrayToStringSlice(p **C.char, num C.size_t) []string {
	var strings []string
	q := uintptr(unsafe.Pointer(p))
	for i := 0; i < int(num); i++ {
		p = (**C.char)(unsafe.Pointer(q))
		if *p == nil {
			break
		}
		strings = append(strings, C.GoString(*p))
		q += unsafe.Sizeof(q)
	}
	return strings
}

func sizedCStringArrayToStringSliceLong(p **C.char, num C.ulong) []string {
	var strings []string
	q := uintptr(unsafe.Pointer(p))
	for i := 0; i < int(num); i++ {
		p = (**C.char)(unsafe.Pointer(q))
		if *p == nil {
			break
		}
		strings = append(strings, C.GoString(*p))
		q += unsafe.Sizeof(q)
	}
	return strings
}

func sizedDoubleArrayToFloat64Slice(p *C.double, num C.size_t) []float64 {
	var nums []float64
	q := uintptr(unsafe.Pointer(p))
	for i := 0; i < int(num); i++ {
		p = (*C.double)(unsafe.Pointer(q))
		nums = append(nums, float64(*p))
		q += unsafe.Sizeof(q)
	}
	return nums
}
