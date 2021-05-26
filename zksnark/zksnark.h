#ifndef _ZKSNARK_H_
#define _ZKSNARK_H_

#include <stdlib.h>
#include <inttypes.h>
#include <stdbool.h>
#include <signal.h>

#ifdef __cplusplus
extern "C" {
#endif

void crash_callback_handler(int sig);
void initialize(bool ff);
void keypair_generator_v2 ( void *pk_ptr, void *vk_ptr, uint64_t input, uint64_t constraints);
void prove_v2(void *proof_ptr, void *pk_ptr, void *hash_ptr, uint64_t input, uint64_t constraints);
bool verify_v2(void *proof_ptr, void *vk_ptr, void *value_prt, uint64_t input);
void prove(void *proof_ptr, void *vk_ptr, void *value_prt, uint64_t constraints);
bool verify(void *proof_ptr, void *vk_ptr, void *value_prt);
bool verify_by_hash(void *proof_ptr, void *vk_ptr, void *hash_ptr);
void circuit(void *hash_ptr, void *value_ptr);

#ifdef __cplusplus
}
#endif

#endif