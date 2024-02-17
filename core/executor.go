package core

import (
	"fmt"
	"os/exec"
)

// Executor is a struct that manages the execution of an already compiled MPC file.
type Executor struct {
	SecrecyLibraryDirectory string
	SecrecySourceDirectory  string
	SecrecyBuildDirectory   string
	SecrecyDataDirectory    string
}

// NewExecutor is a function that creates a new executor with the given secrecy library directory.
func NewExecutor(secrecyLibraryDirectory string) *Executor {
	return &Executor{
		SecrecySourceDirectory: secrecyLibraryDirectory + "/external/secrecy/src/",
		SecrecyBuildDirectory:  secrecyLibraryDirectory + "/external/secrecy/build/",
		SecrecyDataDirectory:   secrecyLibraryDirectory + "/external/secrecy/data/",
	}
}

// Execute is a function that runs the compiled MPC file using the specified cluster and runtime configuration.
func (e *Executor) Execute(executeConfig ExecuteConfig) {

	// execute the compiled program
	cmd := exec.Command("mpirun",
		"-np", fmt.Sprint(executeConfig.ComputingParties),
		"--host "+executeConfig.GetComputingPartiesHostnames(),
		e.SecrecyBuildDirectory+executeConfig.FileName,
		"-t "+fmt.Sprint(executeConfig.ThreadsCount),
		"-b "+fmt.Sprint(executeConfig.BatchSize),
		"-n "+fmt.Sprint(len(executeConfig.ComputingParties)),
	)

	cmd.Dir = e.SecrecyLibraryDirectory
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(stdout))
}
