#include <arm_neon.h>

void VaddqS8(int8x16_t* r, int8x16_t* v0, int8x16_t* v1) { *r = vaddq_s8(*v0, *v1); }
void VaddqU8(uint8x16_t* r, uint8x16_t* v0, uint8x16_t* v1) { *r = vaddq_u8(*v0, *v1); }
void VqaddqS8(int8x16_t* r, int8x16_t* v0, int8x16_t* v1) { *r = vqaddq_s8(*v0, *v1); }
void VqaddqU8(uint8x16_t* r, uint8x16_t* v0, uint8x16_t* v1) { *r = vqaddq_u8(*v0, *v1); }
void VsubqS8(int8x16_t* r, int8x16_t* v0, int8x16_t* v1) { *r = vsubq_s8(*v0, *v1); }
void VsubqU8(uint8x16_t* r, uint8x16_t* v0, uint8x16_t* v1) { *r = vsubq_u8(*v0, *v1); }
void VqsubqS8(int8x16_t* r, int8x16_t* v0, int8x16_t* v1) { *r = vqsubq_s8(*v0, *v1); }
void VqsubqU8(uint8x16_t* r, uint8x16_t* v0, uint8x16_t* v1) { *r = vqsubq_u8(*v0, *v1); }