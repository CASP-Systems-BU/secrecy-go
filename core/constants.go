package core

import "fmt"

type CompileConfig struct {
	Flags     map[string]bool
	Constants map[string]int
}

var DefaultCompileConfig = CompileConfig{
	Flags: map[string]bool{
		"MPC_USE_RANDOM_GENERATOR_TRIPLES": true,
		"MPC_USE_MPI_COMMUNICATOR":         true,
		"MPC_USE_GRPC_COMMUNICATOR":        false,

		"MPC_GENERATE_DATA":             true,
		"MPC_USE_RANDOM_GENERATOR_DATA": true,
		"MPC_EVALUATE_CORRECT_OUTPUT":   true,
		"MPC_CHECK_CORRECTNESS":         true,

		"MPC_PRINT_RESULT":            false,
		"MPC_COMMUNICATOR_PRINT_DATA": false,

		"USE_PARALLEL_PREFIX_ADDER": true,
		"USE_DIVISION_CORRECTION":   true,

		"RECYCLE_THREAD_MEMORY": false,

		"USE_FANTASTIC_FOUR": false,

		"random_generation": false,
	},
	Constants: map[string]int{
		"MPC_RANDOM_DATA_RANGE": 100,
	},
}

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

func NewComputingParty(partyID int, hostname string) ComputingParty {
	return ComputingParty{
		PartyID:  partyID,
		Hostname: hostname,
	}
}

type ComputingParty struct {
	PartyID  int
	Hostname string
}

func NewExecuteConfig(computingParties []ComputingParty, fileName string, threadsCount int, batchSize int) ExecuteConfig {
	return ExecuteConfig{
		ComputingParties: computingParties,
		FileName:         fileName,
		ThreadsCount:     threadsCount,
		BatchSize:        batchSize,
	}
}

type ExecuteConfig struct {
	ComputingParties []ComputingParty

	FileName string

	ThreadsCount int
	BatchSize    int
}

func (e *ExecuteConfig) GetComputingPartiesHostnames() string {
	var hostnames = ""
	for _, computingParty := range e.ComputingParties {
		hostnames += computingParty.Hostname + ","
	}
	return hostnames
}
