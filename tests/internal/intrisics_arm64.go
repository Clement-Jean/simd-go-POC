package internal

/*
#include <arm_neon.h>
*/
import "C"

// typedef uchar uint8_t;
type Uint8 = C.uint8_t

// typedef char int8_t;
type Int8 = C.int8_t

// typedef __attribute__((neon_vector_type(16))) int8_t int8x16_t;
type Int8x16 = C.int8x16_t

// typedef __attribute__((neon_vector_type(16))) uint8_t uint8x16_t;
type Uint8x16 = C.uint8x16_t

//go:linkname VaddqS8 VaddqS8
//go:noescape
func VaddqS8(r, v0, v1 *Int8x16)

//go:linkname VaddqU8 VaddqU8
//go:noescape
func VaddqU8(r, v0, v1 *Uint8x16)

//go:linkname VqaddqS8 VqaddqS8
//go:noescape
func VqaddqS8(r, v0, v1 *Int8x16)

//go:linkname VqaddqU8 VqaddqU8
//go:noescape
func VqaddqU8(r, v0, v1 *Uint8x16)

//go:linkname VsubqS8 VsubqS8
//go:noescape
func VsubqS8(r, v0, v1 *Int8x16)

//go:linkname VsubqU8 VsubqU8
//go:noescape
func VsubqU8(r, v0, v1 *Uint8x16)

//go:linkname VqsubqS8 VqsubqS8
//go:noescape
func VqsubqS8(r, v0, v1 *Int8x16)

//go:linkname VqsubqU8 VqsubqU8
//go:noescape
func VqsubqU8(r, v0, v1 *Uint8x16)

//go:linkname VandqS8 VandqS8
//go:noescape
func VandqS8(r, v0, v1 *Int8x16)

//go:linkname VandqU8 VandqU8
//go:noescape
func VandqU8(r, v0, v1 *Uint8x16)

//go:linkname VorrqS8 VorrqS8
//go:noescape
func VorrqS8(r, v0, v1 *Int8x16)

//go:linkname VorrqU8 VorrqU8
//go:noescape
func VorrqU8(r, v0, v1 *Uint8x16)

//go:linkname VeorqS8 VeorqS8
//go:noescape
func VeorqS8(r, v0, v1 *Int8x16)

//go:linkname VeorqU8 VeorqU8
//go:noescape
func VeorqU8(r, v0, v1 *Uint8x16)

//go:linkname VmaxqS8 VmaxqS8
//go:noescape
func VmaxqS8(r, v0, v1 *Int8x16)

//go:linkname VmaxqU8 VmaxqU8
//go:noescape
func VmaxqU8(r, v0, v1 *Uint8x16)

//go:linkname VminqS8 VminqS8
//go:noescape
func VminqS8(r, v0, v1 *Int8x16)

//go:linkname VminqU8 VminqU8
//go:noescape
func VminqU8(r, v0, v1 *Uint8x16)

//go:linkname VmaxvqS8 VmaxvqS8
//go:noescape
func VmaxvqS8(r *Int8, v0 *Int8x16)

//go:linkname VmaxvqU8 VmaxvqU8
//go:noescape
func VmaxvqU8(r *Uint8, v0 *Uint8x16)

//go:linkname VminvqS8 VminvqS8
//go:noescape
func VminvqS8(r *Int8, v0 *Int8x16)

//go:linkname VminvqU8 VminvqU8
//go:noescape
func VminvqU8(r *Uint8, v0 *Uint8x16)

//go:linkname VextqS8 VextqS8
//go:noescape
func VextqS8(r, v0, v1 *Int8x16)

//go:linkname VextqU8 VextqU8
//go:noescape
func VextqU8(r, v0, v1 *Uint8x16)

//go:linkname Vqtbl1qS8 Vqtbl1qS8
//go:noescape
func Vqtbl1qS8(r, v0 *Int8x16, v1 *Uint8x16)

//go:linkname Vqtbl1qU8 Vqtbl1qU8
//go:noescape
func Vqtbl1qU8(r, v0, v1 *Uint8x16)
