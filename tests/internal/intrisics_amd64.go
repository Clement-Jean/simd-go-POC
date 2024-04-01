package internal

/*
#include <stdint.h>
#include <immintrin.h>
*/
import "C"

// typedef uchar uint8_t;
type Uint8 = C.uint8_t

// typedef char int8_t;
type Int8 = C.int8_t

// typedef float __m128 __attribute__((__vector_size__(16), __aligned__(16)));
type M128 = C.__m128

// typedef double __m128d __attribute__((__vector_size__(16), __aligned__(16)));
type M128D = C.__m128d

// typedef longlong __m128i __attribute__((__vector_size__(16), __aligned__(16)));
type M128I = C.__m128i

// Add packed 8-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname MmAddEpi8 MmAddEpi8
//go:noescape
func MmAddEpi8(r *M128I, v0 *M128I, v1 *M128I)

// Add packed 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
//
//go:linkname MmAddsEpi8 MmAddsEpi8
//go:noescape
func MmAddsEpi8(r *M128I, v0 *M128I, v1 *M128I)

// Add packed unsigned 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
//
//go:linkname MmAddsEpu8 MmAddsEpu8
//go:noescape
func MmAddsEpu8(r *M128I, v0 *M128I, v1 *M128I)

// Subtract packed 8-bit integers in "a" and "b", and store the results in "dst".
//
//go:linkname MmSubEpi8 MmSubEpi8
//go:noescape
func MmSubEpi8(r *M128I, v0 *M128I, v1 *M128I)

// Subtract packed 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
//
//go:linkname MmSubsEpi8 MmSubsEpi8
//go:noescape
func MmSubsEpi8(r *M128I, v0 *M128I, v1 *M128I)

// Subtract packed unsigned 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
//
//go:linkname MmSubsEpu8 MmSubsEpu8
//go:noescape
func MmSubsEpu8(r *M128I, v0 *M128I, v1 *M128I)
