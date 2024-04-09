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
__m128i adds_epi8(__m128i a, __m128i b) { return _mm_adds_epi8(a, b); }
__m128i adds_epu8(__m128i a, __m128i b) { return _mm_adds_epu8(a, b); }
__m128i sub_epi8(__m128i a, __m128i b) { return _mm_sub_epi8(a, b); }
__m128i subs_epi8(__m128i a, __m128i b) { return _mm_subs_epi8(a, b); }
__m128i subs_epu8(__m128i a, __m128i b) { return _mm_subs_epu8(a, b); }
__m128i and_si128(__m128i a, __m128i b) { return _mm_and_si128(a, b); }
__m128i or_si128(__m128i a, __m128i b) { return _mm_or_si128(a, b); }
__m128i xor_si128(__m128i a, __m128i b) { return _mm_xor_si128(a, b); }
__m128i max_epi8(__m128i a, __m128i b) { return _mm_max_epu8(a, b); }
__m128i max_epu8(__m128i a, __m128i b) { return _mm_max_epu8(a, b); }
__m128i min_epi8(__m128i a, __m128i b) { return _mm_min_epu8(a, b); }
__m128i min_epu8(__m128i a, __m128i b) { return _mm_min_epu8(a, b); }
*/
import "C"

// typedef uchar uint8_t;
type Uint8 = C.uint8_t

// typedef char int8_t;
type Int8 = C.int8_t

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
func MmAddsEpi8(v0, v1 M128I) M128I { return C.adds_epi8(v0, v1) }

// Add packed unsigned 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
func MmAddsEpu8(v0, v1 M128I) M128I { return C.adds_epu8(v0, v1) }

// Subtract packed 8-bit integers in "a" and "b", and store the results in "dst".
func MmSubEpi8(v0, v1 M128I) M128I { return C.sub_epi8(v0, v1) }

// Subtract packed 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
func MmSubsEpi8(v0, v1 M128I) M128I { return C.subs_epi8(v0, v1) }

// Subtract packed unsigned 8-bit integers in "a" and "b" using saturation, and store the results in "dst".
func MmSubsEpu8(v0, v1 M128I) M128I { return C.subs_epu8(v0, v1) }

// Compute the bitwise AND of 128 bits (representing integer data) in a and b, and store the result in dst.
func MmAndSi128(v0, v1 M128I) M128I { return C.and_si128(v0, v1) }

// Compute the bitwise OR of 128 bits (representing integer data) in a and b, and store the result in dst.
func MmOrSi128(v0, v1 M128I) M128I { return C.or_si128(v0, v1) }

// Compute the bitwise Xor of 128 bits (representing integer data) in a and b, and store the result in dst.
func MmXorSi128(v0, v1 M128I) M128I { return C.xor_si128(v0, v1) }

// Compare packed signed 8-bit integers in a and b, and store packed maximum values in dst.
func MmMaxEpi8(v0, v1 M128I) M128I { return C.max_epi8(v0, v1) }

// Compare packed unsigned 8-bit integers in a and b, and store packed maximum values in dst.
func MmMaxEpu8(v0, v1 M128I) M128I { return C.max_epu8(v0, v1) }

// Compare packed signed 8-bit integers in a and b, and store packed minimum values in dst.
func MmMinEpi8(v0, v1 M128I) M128I { return C.min_epi8(v0, v1) }

// Compare packed unsigned 8-bit integers in a and b, and store packed minimum values in dst.
func MmMinEpu8(v0, v1 M128I) M128I { return C.min_epu8(v0, v1) }
