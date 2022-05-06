#include "zksnark.h"

#include <iostream>
#include "gadgetlib1/gadgets/basic_gadgets.hpp"
#include "zk_proof_systems/ppzksnark/r1cs_gg_ppzksnark/r1cs_gg_ppzksnark.hpp"
#include "common/default_types/r1cs_gg_ppzksnark_pp.hpp"

#include "common/serialization.hpp"

using namespace libsnark;
using namespace std;

typedef Fr<default_r1cs_gg_ppzksnark_pp> FieldT;

const int input_length = 32;
const int proof_length = 259;
const int vk_length = 2882;

const int pk_chars_size_v2 = 344733; // 1709933;  3417773; 6843373;
const int vk_chars_size_v2 = 2882;
const int prof_chars_size_v2 = 259;

bool is_init = false;

void initialize(bool ff)
{
    default_r1cs_gg_ppzksnark_pp::init_public_params();
    inhibit_profiling = ff;
    inhibit_profiling_info = true;
    inhibit_profiling_counters = true;
    signal(SIGILL, crash_callback_handler);
    is_init = true;
}

void crash_callback_handler(int sig){
    cout << "Lib zkSnark is not working! Fix it!" << endl;
    exit(1);
}

void keypair_generator_v2 ( void *pk_ptr,
                            void *vk_ptr,
                            uint64_t input,
                            uint64_t constraints
                            ){
    unsigned char *pk_out = reinterpret_cast<unsigned char *>(pk_ptr);
    unsigned char *vk_out = reinterpret_cast<unsigned char *>(vk_ptr);
    size_t num_inputs = input;
    size_t num_constraints = constraints;

    if(!is_init){
       signal(SIGSEGV, crash_callback_handler);
    }
    r1cs_constraint_system<FieldT> cs;
    cs.primary_input_size = num_inputs;
    cs.auxiliary_input_size = num_constraints;

    size_t lastvar = num_inputs-1;
    for (size_t i = 0; i < num_constraints; ++i)
    {
        ++lastvar;
        const size_t u = (i == 0 ? 42 % num_inputs : i % 2);
        const size_t v = (i == 0 ? 42 % num_inputs : i % 2);
        linear_combination<FieldT> A, B, C;
        A.add_term(u+1, 2);
        B.add_term(v+1, 1);
        if (u == v)
        {
            C.add_term(u+1, 2);
        }
        else
        {
            C.add_term(u+1, 1);
            C.add_term(v+1, 1);
        }
        C.add_term(lastvar+1, -FieldT::one());
        cs.add_constraint(r1cs_constraint<FieldT>(A, B, C));
    }
    assert(cs.num_variables() >= num_inputs);
    assert(cs.num_inputs() == num_inputs);
    assert(cs.num_constraints() == num_constraints);
    auto keypair = r1cs_gg_ppzksnark_generator<default_r1cs_gg_ppzksnark_pp>(cs);
    std::stringstream pk_data;
    pk_data << keypair.pk ;
    auto pk_str = pk_data.str();
    if(inhibit_profiling)
    cout << "* PK size in char count: " << pk_str.size() << endl;
    assert(pk_str.size() == pk_chars_size_v2);
    for (int i = 0; i < pk_chars_size_v2; i++) {
        pk_out[i] = pk_str[i];
    }
    std::stringstream vk_data;
    vk_data << keypair.vk ;
    auto vk_str = vk_data.str();
    if(inhibit_profiling)
    cout << "* VK size in char count: " << vk_str.size() << endl;
    assert(vk_str.size() == vk_chars_size_v2);
    for (int i = 0; i < vk_chars_size_v2; i++) {
        vk_out[i] = vk_str[i];
    }
}

void prove_v2(
                void *proof_ptr,
                void *pk_ptr,
                void *hash_ptr,
                uint64_t input,
                uint64_t constraints
){
    unsigned char *proof_out = reinterpret_cast<unsigned char *>(proof_ptr);
    unsigned char *pk_in = reinterpret_cast<unsigned char *>(pk_ptr);
    unsigned char *hash_in = reinterpret_cast<unsigned char *>(hash_ptr);

    size_t num_inputs = input;
    size_t num_constraints = constraints;

    if(!is_init){
        signal(SIGSEGV, crash_callback_handler);
    }
    std::stringstream pk_data;
    for (int i = 0; i < pk_chars_size_v2; i++) {
       pk_data << pk_in[i];
    }
    assert(pk_data.str().size() == pk_chars_size_v2);
    pk_data.rdbuf()->pubseekpos(0, std::ios_base::in);

    r1cs_gg_ppzksnark_proving_key<default_r1cs_gg_ppzksnark_pp> proving_key;
    pk_data >> proving_key;

    //size_t num_inputs = 32;//proving_key.constraint_system.num_inputs();
    //size_t num_constraints = 20000;//proving_key.constraint_system.num_constraints();

    r1cs_variable_assignment<FieldT> full_variable_assignment;
    for (size_t i = 0; i < num_inputs; ++i)
    {
        full_variable_assignment.push_back(FieldT(hash_in[i]));
    }
    r1cs_constraint_system<FieldT> cs;
    cs.primary_input_size = num_inputs;
    cs.auxiliary_input_size = num_constraints;

    size_t lastvar = num_inputs-1;
    for (size_t i = 0; i < num_constraints; ++i)
    {
        ++lastvar;
        const size_t u = (i == 0 ? 42 % num_inputs : i % 2);
        const size_t v = (i == 0 ? 42 % num_inputs : i % 2);
        linear_combination<FieldT> A, B, C;
        A.add_term(u+1, 2);
        B.add_term(v+1, 1);
        if (u == v)
        {
            C.add_term(u+1, 2);
        }
        else
        {
            C.add_term(u+1, 1);
            C.add_term(v+1, 1);
        }
        C.add_term(lastvar+1, -FieldT::one());
        cs.add_constraint(r1cs_constraint<FieldT>(A, B, C));
        full_variable_assignment.push_back(full_variable_assignment[u] + full_variable_assignment[v] - full_variable_assignment[u] * full_variable_assignment[v] - full_variable_assignment[u] * full_variable_assignment[v]);
    }

    /* split variable assignment */
    r1cs_primary_input<FieldT> primary_input(full_variable_assignment.begin(), full_variable_assignment.begin() + num_inputs);
    r1cs_primary_input<FieldT> auxiliary_input(full_variable_assignment.begin() + num_inputs, full_variable_assignment.end());

    /* sanity checks */
    assert(cs.num_variables() == full_variable_assignment.size());
    assert(cs.num_variables() >= num_inputs);
    assert(cs.num_inputs() == num_inputs);
    assert(cs.num_constraints() == num_constraints);
    assert(cs.is_satisfied(primary_input, auxiliary_input));

    auto proof_zksnark = r1cs_gg_ppzksnark_prover<default_r1cs_gg_ppzksnark_pp>(proving_key, primary_input, auxiliary_input);

    std::stringstream proof_data;
    proof_data << proof_zksnark;
    auto proof_str = proof_data.str();
    if(inhibit_profiling)
    cout << "* Proof size in char count: " << proof_str.size() << endl;
    assert(proof_str.size() == prof_chars_size_v2);
    for (int i = 0; i < prof_chars_size_v2; i++) {
        proof_out[i] = proof_str[i];
    }
}

bool verify_v2( void *proof_ptr,
                void *vk_ptr,
                void *hash_ptr,
                uint64_t input){

    unsigned char *proof_in = reinterpret_cast<unsigned char *>(proof_ptr);
    unsigned char *vk_in = reinterpret_cast<unsigned char *>(vk_ptr);
    unsigned char *hash_in = reinterpret_cast<unsigned char *>(hash_ptr);

    size_t num_inputs = input;

    std::vector<unsigned char> proof_v(proof_in, proof_in+prof_chars_size_v2);
    std::stringstream proof_data;
    for (int i = 0; i < prof_chars_size_v2; i++) {
        proof_data << proof_v[i];
    }
    assert(proof_data.str().size() == prof_chars_size_v2);
    proof_data.rdbuf()->pubseekpos(0, std::ios_base::in);
    r1cs_gg_ppzksnark_proof<default_r1cs_gg_ppzksnark_pp> proof;
    proof_data >> proof;

    std::stringstream vk_data;
    for (int i = 0; i < vk_chars_size_v2; i++) {
       vk_data << vk_in[i];
    }
    assert(vk_data.str().size() == vk_chars_size_v2 );
    vk_data.rdbuf()->pubseekpos(0, std::ios_base::in);
    r1cs_gg_ppzksnark_verification_key<default_r1cs_gg_ppzksnark_pp> verification_key;
    vk_data >> verification_key;

    r1cs_variable_assignment<FieldT> full_variable_assignment;
    for (size_t i = 0; i < num_inputs; ++i)
    {
        full_variable_assignment.push_back(FieldT(hash_in[i]));
    }
    r1cs_primary_input<FieldT> primary_input(full_variable_assignment.begin(), full_variable_assignment.begin() + num_inputs);
    const bool ans = r1cs_gg_ppzksnark_verifier_strong_IC<default_r1cs_gg_ppzksnark_pp>(verification_key, primary_input, proof);
    return ans;
}

void prove(     void *proof_ptr,
                void *vk_ptr,
                void *value_ptr,
                uint64_t constraints){
    unsigned char *proof = reinterpret_cast<unsigned char *>(proof_ptr);
    unsigned char *vk = reinterpret_cast<unsigned char *>(vk_ptr);
    unsigned char *value = reinterpret_cast<unsigned char *>(value_ptr);

    if(!is_init){
        signal(SIGSEGV, crash_callback_handler);
    }

    size_t num_inputs = input_length;
    size_t num_constraints = constraints;
    r1cs_constraint_system<FieldT> cs;

    cs.primary_input_size = num_inputs;
    cs.auxiliary_input_size = num_constraints; /* we will add one auxiliary variable per constraint */

    r1cs_variable_assignment<FieldT> full_variable_assignment;
    for (size_t i = 0; i < num_inputs; ++i)
    {
        full_variable_assignment.push_back(FieldT(value[i] % 2));
    }
    size_t lastvar = num_inputs-1;
    for (size_t i = 0; i < num_constraints; ++i)
    {
        ++lastvar;
        const size_t u = (i == 0 ? std::rand() % num_inputs : std::rand() % i);
        const size_t v = (i == 0 ? std::rand() % num_inputs : std::rand() % i);
        linear_combination<FieldT> A, B, C;
        A.add_term(u+1, 2);
        B.add_term(v+1, 1);
        if (u == v)
        {
            C.add_term(u+1, 2);
        }
        else
        {
            C.add_term(u+1, 1);
            C.add_term(v+1, 1);
        }
        C.add_term(lastvar+1, -FieldT::one());

        cs.add_constraint(r1cs_constraint<FieldT>(A, B, C));
        full_variable_assignment.push_back(full_variable_assignment[u] + full_variable_assignment[v] - full_variable_assignment[u] * full_variable_assignment[v] - full_variable_assignment[u] * full_variable_assignment[v]);
    }

     /* split variable assignment */
    r1cs_primary_input<FieldT> primary_input(full_variable_assignment.begin(), full_variable_assignment.begin() + num_inputs);
    r1cs_primary_input<FieldT> auxiliary_input(full_variable_assignment.begin() + num_inputs, full_variable_assignment.end());

    /* sanity checks */
    assert(cs.num_variables() == full_variable_assignment.size());
    assert(cs.num_variables() >= num_inputs);
    assert(cs.num_inputs() == num_inputs);
    assert(cs.num_constraints() == num_constraints);
    assert(cs.is_satisfied(primary_input, auxiliary_input));

    const r1cs_constraint_system<FieldT> constraint_system = cs;
    cout << "Number of R1CS constraints: " << constraint_system.num_constraints() << endl;
    auto keypair = r1cs_gg_ppzksnark_generator<default_r1cs_gg_ppzksnark_pp>(constraint_system);

    auto proof_zksnark = r1cs_gg_ppzksnark_prover<default_r1cs_gg_ppzksnark_pp>(keypair.pk, primary_input, auxiliary_input);
    std::stringstream proof_data;
    proof_data << proof_zksnark;
    auto proof_str = proof_data.str();
    cout << "* Proof size in char count: " << proof_str.size() << endl;
    assert(proof_str.size() == proof_length);
    for (int i = 0; i < proof_length; i++) {
        proof[i] = proof_str[i];
    }
    std::stringstream vk_data;
    vk_data << keypair.vk ;
    auto vk_str = vk_data.str();
    cout << "* VK size in char count: " << vk_str.size() << endl;
    assert(vk_str.size() == vk_length);

    for (int i = 0; i < vk_length; i++) {
        vk[i] = vk_str[i];
    }
}

bool verify(    void *proof_ptr,
                void *vk_ptr,
                void *value_prt){

    unsigned char *proof_v = reinterpret_cast<unsigned char *>(proof_ptr);
    unsigned char *vk_v = reinterpret_cast<unsigned char *>(vk_ptr);
    unsigned char *value = reinterpret_cast<unsigned char *>(value_prt);

    if(!is_init){
            signal(SIGSEGV, crash_callback_handler);
    }

    std::stringstream proof_data;
    for (int i = 0; i < proof_length; i++) {
        proof_data << proof_v[i];
    }
    assert(proof_data.str().size() == proof_length);
    proof_data.rdbuf()->pubseekpos(0, std::ios_base::in);
    r1cs_gg_ppzksnark_proof<default_r1cs_gg_ppzksnark_pp> proof_obj;
    proof_data >> proof_obj;


    std::stringstream vk_data;
    for (int i = 0; i < vk_length; i++) {
       vk_data << vk_v[i];
    }
    assert(vk_data.str().size() == vk_length);
    vk_data.rdbuf()->pubseekpos(0, std::ios_base::in);
    r1cs_gg_ppzksnark_verification_key<default_r1cs_gg_ppzksnark_pp> vk_obj;
    vk_data >> vk_obj;

    size_t num_inputs = input_length;
    r1cs_variable_assignment<FieldT> full_variable_assignment;
    for (size_t i = 0; i < num_inputs; ++i)
    {
        full_variable_assignment.push_back(FieldT(value[i] % 2));
    }
    r1cs_primary_input<FieldT> primary_input(full_variable_assignment.begin(), full_variable_assignment.begin() + num_inputs);

    const bool ans = r1cs_gg_ppzksnark_verifier_strong_IC<default_r1cs_gg_ppzksnark_pp>(vk_obj, primary_input, proof_obj);
    return ans;
}


bool verify_by_hash(void *proof_ptr,
                    void *vk_ptr,
                    void *hash_prt){

    unsigned char *proof_v = reinterpret_cast<unsigned char *>(proof_ptr);
    unsigned char *vk_v = reinterpret_cast<unsigned char *>(vk_ptr);
    unsigned char *hash_v = reinterpret_cast<unsigned char *>(hash_prt);

    std::stringstream proof_data;
    for (int i = 0; i < proof_length; i++) {
        proof_data << proof_v[i];
    }
    assert(proof_data.str().size() == proof_length);
    proof_data.rdbuf()->pubseekpos(0, std::ios_base::in);
    r1cs_gg_ppzksnark_proof<default_r1cs_gg_ppzksnark_pp> proof_obj;
    proof_data >> proof_obj;


    std::stringstream vk_data;
    for (int i = 0; i < vk_length; i++) {
       vk_data << vk_v[i];
    }
    assert(vk_data.str().size() == vk_length);
    vk_data.rdbuf()->pubseekpos(0, std::ios_base::in);
    r1cs_gg_ppzksnark_verification_key<default_r1cs_gg_ppzksnark_pp> vk_obj;
    vk_data >> vk_obj;

    r1cs_primary_input<FieldT> primary_input(hash_v, hash_v + 32);

    const bool ans = r1cs_gg_ppzksnark_verifier_strong_IC<default_r1cs_gg_ppzksnark_pp>(vk_obj, primary_input, proof_obj);
    return ans;
}

void circuit( void *hash_ptr, void *value_ptr){

    unsigned char *hash = reinterpret_cast<unsigned char *>(hash_ptr);
    unsigned char *value = reinterpret_cast<unsigned char *>(value_ptr);

    size_t num_inputs = input_length;
    r1cs_variable_assignment<FieldT> full_variable_assignment;
    for (size_t i = 0; i < num_inputs; ++i)
    {
        hash[i] = (unsigned char)(value[i] % 2);
    }
}



