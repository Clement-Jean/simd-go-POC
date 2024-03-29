package internal

/*
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
