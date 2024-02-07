package core

import (
	"os"
	"os/exec"
)

type Compiler struct {
	SecrecyLibraryDirectory string
	SecrecySourceDirectory  string
	SecrecyBuildDirectory   string
	SecrecyDataDirectory    string
}

func NewCompiler(secrecyLibraryDirectory string) *Compiler {
	return &Compiler{
		SecrecySourceDirectory: secrecyLibraryDirectory + "/external/secrecy/src/",
		SecrecyBuildDirectory:  secrecyLibraryDirectory + "/external/secrecy/build/",
		SecrecyDataDirectory:   secrecyLibraryDirectory + "/external/secrecy/data/",
	}
}

func (e *Compiler) CreateCPPFile(fileName string, fileContent string) {
	// create file with given name and content inside the secrecy source directory
	cppFile, err := os.Create(e.SecrecySourceDirectory + "/" + fileName + ".cpp")
	if err != nil {
		panic(err)
	}

	defer cppFile.Close()

	_, err = cppFile.WriteString(fileContent)
	if err != nil {
		panic(err)
	}

	cppFile.Sync()
}

func (e *Compiler) CreateCMake(compileConfig CompileConfig) {
	// TODO: can we get rid of deletion step?
	// make sure the build directory is empty
	err := os.RemoveAll(e.SecrecyBuildDirectory)
	if err != nil {
		panic(err)
	}

	err = os.Mkdir(e.SecrecyBuildDirectory, 0755)
	if err != nil {
		panic(err)
	}

	var cmakeCommand = "cmake .. " + DefaultCompileConfig.GenerateCMakeDefintionsString(compileConfig)

	cmd := exec.Command(cmakeCommand)
	cmd.Dir = e.SecrecyBuildDirectory
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	println(string(stdout))

}

func (e *Compiler) BuildMPCFile(fileName string, fileContent string, compileConfig CompileConfig) {
	// Create the file
	e.CreateCPPFile(fileName, fileContent)

	// Create the build directory
	e.CreateCMake(compileConfig)

	// Build the file
	cmd := exec.Command("make", "-j", "4", fileName)
	cmd.Dir = e.SecrecyBuildDirectory
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	println(string(stdout))
}
