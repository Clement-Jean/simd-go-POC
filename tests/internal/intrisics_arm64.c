#include <arm_neon.h>

void VaddqS8(int8x16_t* r, int8x16_t* v0, int8x16_t* v1) { *r = vaddq_s8(*v0, *v1); }
void VaddqU8(uint8x16_t* r, uint8x16_t* v0, uint8x16_t* v1) { *r = vaddq_u8(*v0, *v1); }
void VqaddqS8(int8x16_t* r, int8x16_t* v0, int8x16_t* v1) { *r = vqaddq_s8(*v0, *v1); }
void VqaddqU8(uint8x16_t* r, uint8x16_t* v0, uint8x16_t* v1) { *r = vqaddq_u8(*v0, *v1); }
void VsubqS8(int8x16_t* r, int8x16_t* v0, int8x16_t* v1) { *r = vsubq_s8(*v0, *v1); }
void VsubqU8(uint8x16_t* r, uint8x16_t* v0, uint8x16_t* v1) { *r = vsubq_u8(*v0, *v1); }
void VqsubqS8(int8x16_t* r, int8x16_t* v0, int8x16_t* v1) { *r = vqsubq_s8(*v0, *v1); }
void VqsubqU8(uint8x16_t* r, uint8x16_t* v0, uint8x16_t* v1) { *r = vqsubq_u8(*v0, *v1); }
void VandqS8(int8x16_t* r, int8x16_t* v0, int8x16_t* v1) { *r = vandq_s8(*v0, *v1); }
void VandqU8(uint8x16_t* r, uint8x16_t* v0, uint8x16_t* v1) { *r = vandq_u8(*v0, *v1); }
void VorrqS8(int8x16_t* r, int8x16_t* v0, int8x16_t* v1) { *r = vorrq_s8(*v0, *v1); }
void VorrqU8(uint8x16_t* r, uint8x16_t* v0, uint8x16_t* v1) { *r = vorrq_u8(*v0, *v1); }
void VeorqS8(int8x16_t* r, int8x16_t* v0, int8x16_t* v1) { *r = veorq_s8(*v0, *v1); }
void VeorqU8(uint8x16_t* r, uint8x16_t* v0, uint8x16_t* v1) { *r = veorq_u8(*v0, *v1); }
void VmaxqS8(int8x16_t* r, int8x16_t* v0, int8x16_t* v1) { *r = vmaxq_s8(*v0, *v1); }
void VmaxqU8(uint8x16_t* r, uint8x16_t* v0, uint8x16_t* v1) { *r = vmaxq_u8(*v0, *v1); }
void VminqS8(int8x16_t* r, int8x16_t* v0, int8x16_t* v1) { *r = vminq_s8(*v0, *v1); }
void VminqU8(uint8x16_t* r, uint8x16_t* v0, uint8x16_t* v1) { *r = vminq_u8(*v0, *v1); }
void VmaxvqS8(int8_t* r, int8x16_t* v0) { *r = vmaxvq_s8(*v0); }
void VmaxvqU8(uint8_t* r, uint8x16_t* v0) { *r = vmaxvq_u8(*v0); }
void VminvqS8(int8_t* r, int8x16_t* v0) { *r = vminvq_s8(*v0); }
void VminvqU8(uint8_t* r, uint8x16_t* v0) { *r = vminvq_u8(*v0); }
