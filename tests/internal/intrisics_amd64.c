#include <immintrin.h>

void MmAddEpi8(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_add_epi8(*v0, *v1); }
