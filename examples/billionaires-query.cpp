#include "../include/secrecy.h"

using namespace secrecy::service::mpi_service::spdz_npc;

int main(int argc, char **argv){
    secrecy_init(argc, argv);
    auto pID = secrecy::service::runTime.getPartyID();
    int test_size = 128;
    if(argc >= 5){
        test_size = atoi(argv[4]);
    }

    // Prologue to read in the inputs
    const int size = test_size;
    std::vector<std::string> schema = {"[BILLIONAIRE]", "[WEALTH]"};
    EncodedTable<int> billionaires("billionaires", schema, size);
    billionaires.input("[BILLIONAIRE]", "input-billionaires");
    billionaires.input("[WEALTH]", "input-wealth");

    // The logic
    billionaires.sort({"[WEALTH]"}, {false});

    // Epilogue to return the outputs
    billionaires.output("[BILLIONAIRE]", "output-res");

    // Print the result
    // auto c_open = billionaires.open();
    // secrecy::debug::print_table(c_open, secrecy::service::runTime.getPartyID());

#if defined(MPC_USE_MPI_COMMUNICATOR)
    MPI_Finalize();
#endif

    return 0;
}
