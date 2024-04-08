package internal

/*
#cgo CFLAGS: -msse2
#include <stdint.h>
#include <immintrin.h>

void storeu_si128(uint8_t* m, __m128i v0) { _mm_storeu_si128((__m128i*)m, v0); }

__m128i setr_epi8(uint8_t a, uint8_t b, uint8_t c, uint8_t d, uint8_t e, uint8_t f, uint8_t g, uint8_t h, uint8_t i, uint8_t j, uint8_t k, uint8_t l, uint8_t m, uint8_t n, uint8_t o, uint8_t p) {
   return _mm_setr_epi8(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p);
}

__m128i add_epi8(__m128i a, __m128i b) { return _mm_add_epi8(a, b); }
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

// Store 128-bits (composed of 16 packed 8-bit integers) from a into memory. mem_addr does not need to be aligned on any particular boundary.
func MmStoreuSi128(m *Uint8, v0 M128I) { C.storeu_si128(m, v0) }

// Set packed 8-bit integers in dst with the supplied values in reverse order.
func MmSetrEpi8[T uint8 | int8](a [16]T) M128I {
	return C.setr_epi8(
		(Uint8)(a[0]), (Uint8)(a[1]), (Uint8)(a[2]), (Uint8)(a[3]), (Uint8)(a[4]), (Uint8)(a[5]),
		(Uint8)(a[6]), (Uint8)(a[7]), (Uint8)(a[8]), (Uint8)(a[9]), (Uint8)(a[10]), (Uint8)(a[11]),
		(Uint8)(a[12]), (Uint8)(a[13]), (Uint8)(a[14]), (Uint8)(a[15]),
	)
}

// Add packed 8-bit integers in "a" and "b", and store the results in "dst".
func MmAddEpi8(v0, v1 M128I) M128I { return C.add_epi8(v0, v1) }

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

// Compute the bitwise AND of 128 bits (representing integer data) in a and b, and store the result in dst.
//
//go:linkname MmAndSi128 MmAndSi128
//go:noescape
func MmAndSi128(r *M128I, v0 *M128I, v1 *M128I)

// Compute the bitwise OR of 128 bits (representing integer data) in a and b, and store the result in dst.
//
//go:linkname MmOrSi128 MmOrSi128
//go:noescape
func MmOrSi128(r *M128I, v0 *M128I, v1 *M128I)

// Compute the bitwise Xor of 128 bits (representing integer data) in a and b, and store the result in dst.
//
//go:linkname MmXorSi128 MmXorSi128
//go:noescape
func MmXorSi128(r *M128I, v0 *M128I, v1 *M128I)

// Compare packed signed 8-bit integers in a and b, and store packed maximum values in dst.
//
//go:linkname MmMaxEpi8 MmMaxEpi8
//go:noescape
func MmMaxEpi8(r *M128I, v0 *M128I, v1 *M128I)

// Compare packed unsigned 8-bit integers in a and b, and store packed maximum values in dst.
//
//go:linkname MmMaxEpu8 MmMaxEpu8
//go:noescape
func MmMaxEpu8(r *M128I, v0 *M128I, v1 *M128I)

// Compare packed signed 8-bit integers in a and b, and store packed minimum values in dst.
//
//go:linkname MmMinEpi8 MmMinEpi8
//go:noescape
func MmMinEpi8(r *M128I, v0 *M128I, v1 *M128I)

// Compare packed unsigned 8-bit integers in a and b, and store packed minimum values in dst.
//
//go:linkname MmMinEpu8 MmMinEpu8
//go:noescape
func MmMinEpu8(r *M128I, v0 *M128I, v1 *M128I)
