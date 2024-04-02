#include <immintrin.h>

void MmAddEpi8(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_add_epi8(*v0, *v1); }
void MmAddsEpi8(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_adds_epi8(*v0, *v1); }
void MmAddsEpu8(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_adds_epu8(*v0, *v1); }
void MmSubEpi8(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_sub_epi8(*v0, *v1); }
void MmSubsEpi8(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_subs_epi8(*v0, *v1); }
void MmSubsEpu8(__m128i* r, __m128i* v0, __m128i* v1) { *r = _mm_subs_epu8(*v0, *v1); }
void MmAndSi128(__m128i* r, __m128i* v0, __m128i* v1) { *r =  _mm_and_si128(*v0, *v1); }
void MmOrSi128(__m128i* r, __m128i* v0, __m128i* v1) { *r =  _mm_or_si128(*v0, *v1); }
void MmXorSi128(__m128i* r, __m128i* v0, __m128i* v1) { *r =  _mm_xor_si128(*v0, *v1); }
void MmMaxEpi8(__m128i* r, __m128i* v0, __m128i* v1) { *r =  _mm_max_epi8(*v0, *v1); }
void MmMaxEpu8(__m128i* r, __m128i* v0, __m128i* v1) { *r =  _mm_max_epu88(*v0, *v1); }
void MmMinEpi8(__m128i* r, __m128i* v0, __m128i* v1) { *r =  _mm_min_epi8(*v0, *v1); }
void MmMinEpu8(__m128i* r, __m128i* v0, __m128i* v1) { *r =  _mm_min_epu88(*v0, *v1); }
