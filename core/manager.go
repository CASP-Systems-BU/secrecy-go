package core

// Manager is a struct that manages the compilation and execution
// of a give MPC file written in CPP using the secrecy framework.
type Manager struct {
	Compiler
	Executor
}

// NewManager is a function that creates a new manager with the given
// secrecy library directory.
func NewManager(secrecyLibraryDirectory string) *Manager {
	return &Manager{
		Compiler: *NewCompiler(secrecyLibraryDirectory),
		Executor: *NewExecutor(secrecyLibraryDirectory),
	}
}

// BuildAndExecute is a function that builds and executes the given MPC file.
func (m *Manager) BuildAndExecute(fileName string, fileContent string, compileConfig CompileConfig, executeConfig ExecuteConfig) {
	m.BuildMPCFile(fileName, fileContent, compileConfig)
	m.Execute(executeConfig)
}

// HelloWorld is a function that prints "Hello World".
// used for testing purposes.
func (m *Manager) HelloWorld() {
	println("Hello World")
}
