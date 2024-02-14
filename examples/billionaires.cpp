#include "../include/secrecy.h"

using namespace secrecy::service::mpi_service::spdz_npc;

int main(int argc, char **argv){
    secrecy_init(argc, argv);

    // Prologue to read in the inputs
    BSharedVector<int32_t> a(1, "input-jeff"), b(1, "input-elon");

    // The logic
    BSharedVector<int32_t> c = a < b;

    // Epilogue to return the outputs
    c.output("output-res");

    // Print the result
    auto c_open = c.open();
    secrecy::debug::print(c_open, secrecy::service::runTime.getPartyID());

#if defined(MPC_USE_MPI_COMMUNICATOR)
    MPI_Finalize();
#endif

    return 0;
}
