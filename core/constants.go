package core

import "fmt"

// CompileConfig is a struct that contains the flags and constants
// that are used to compile the MPC file written in CPP.
type CompileConfig struct {
	Flags     map[string]bool
	Constants map[string]int
}

// DefaultCompileConfig is a CompileConfig struct that contains the default flags and constants
// These flags are the one specified in the debug/debug.h header file in the secrecy directory.
var DefaultCompileConfig = CompileConfig{
	Flags: map[string]bool{
		"MPC_USE_RANDOM_GENERATOR_TRIPLES": true,
		"MPC_USE_MPI_COMMUNICATOR":         true, // Turns on/off communication between parties
		"MPC_USE_GRPC_COMMUNICATOR":        false,

		"MPC_GENERATE_DATA":             true,
		"MPC_USE_RANDOM_GENERATOR_DATA": true,
		"MPC_EVALUATE_CORRECT_OUTPUT":   true,
		"MPC_CHECK_CORRECTNESS":         true,

		"MPC_PRINT_RESULT":            false,
		"MPC_COMMUNICATOR_PRINT_DATA": false,

		"USE_PARALLEL_PREFIX_ADDER": true, // Turns on the parallel prefix adder. Otherwise, the ripple carry adder is used
		"USE_DIVISION_CORRECTION":   true, // Turns on division correction for the off by 1 error in the public divison.

		"RECYCLE_THREAD_MEMORY": false, // Turns on/off the recycling of the memory inside each thread (experimental)

		"USE_FANTASTIC_FOUR": false, // Turn on to use fantastic four protocol in the benchmarking experiments.

		"random_generation": false,
	},
	Constants: map[string]int{
		"MPC_RANDOM_DATA_RANGE": 100,
	},
}

// GenerateCMakeDefintionsString is a function that generates the arguments for the CMake command using the given compile configuration.
func (c CompileConfig) GenerateCMakeDefintionsString(compileConfig CompileConfig) string {
	var mergedCompileConfig = c
	for k, v := range compileConfig.Flags {
		mergedCompileConfig.Flags[k] = v
	}
	for k, v := range compileConfig.Constants {
		mergedCompileConfig.Constants[k] = v
	}

	// -DPROTOCOL_VAR='-DUSE_FANTASTIC_FOUR -DMPC_RANDOM_DATA_RANGE=100'
	var cmakeDefinitinosString = "-DPROTOCOL_VAR='"

	var cmakeCompilationFlags = ""
	for k, v := range mergedCompileConfig.Flags {
		if v {
			cmakeCompilationFlags += "-D" + k + " "
		}
	}

	var cmakeConstants = ""
	for k, v := range mergedCompileConfig.Constants {
		cmakeConstants += "-D" + k + "=" + fmt.Sprint(v) + " "
	}

	cmakeDefinitinosString += cmakeCompilationFlags + cmakeConstants + "' "

	return cmakeDefinitinosString
}

// NewComputingParty is a function that creates a new computing party with the given ID and hostname.
func NewComputingParty(partyID int, hostname string) ComputingParty {
	return ComputingParty{
		PartyID:  partyID,
		Hostname: hostname,
	}
}

// ComputingParty is a struct that contains the ID and the hostname of each computing party.
type ComputingParty struct {
	PartyID  int
	Hostname string
}

// NewExecuteConfig is a function that creates a new execution configuration with the given computing parties, file name, threads count, and batch size.
func NewExecuteConfig(computingParties []ComputingParty, fileName string, threadsCount int, batchSize int) ExecuteConfig {
	return ExecuteConfig{
		ComputingParties: computingParties,
		FileName:         fileName,
		ThreadsCount:     threadsCount,
		BatchSize:        batchSize,
	}
}

// ExecuteConfig is a struct that contains the configuration for the
// execution of the MPC file on the specified computing parties cluster
// in addition to runtime paramters (threads - batch size).
type ExecuteConfig struct {
	ComputingParties []ComputingParty

	FileName string

	ThreadsCount int
	BatchSize    int
}

// GetComputingPartiesHostnames is a function that returns the hostnames of
// the computing parties in the execution configuration.
func (e *ExecuteConfig) GetComputingPartiesHostnames() string {
	var hostnames = ""
	for _, computingParty := range e.ComputingParties {
		hostnames += computingParty.Hostname + ","
	}
	return hostnames
}
