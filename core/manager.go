package core

type Manager struct {
	Compiler
	Executor
}

func NewManager(secrecyLibraryDirectory string) *Manager {
	return &Manager{
		Compiler: *NewCompiler(secrecyLibraryDirectory),
		Executor: *NewExecutor(secrecyLibraryDirectory),
	}
}

func (m *Manager) BuildAndExecute(fileName string, fileContent string, compileConfig CompileConfig, executeConfig ExecuteConfig) {
	m.BuildMPCFile(fileName, fileContent, compileConfig)
	m.Execute(executeConfig)
}

func (m *Manager) HelloWorld() {
	println("Hello World")
}
