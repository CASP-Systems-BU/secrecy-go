# secrecy-go
A Go package that encapsulates the secrecy framework for one computing party and provides secure computation management across the computing parties.
Specifically, the package has the following functionalities:

- **Subpackage `core`**: Exports functions to control the execution of the secure computation for each computing party. It is used to configure compilation and runtime parameters and can be imported inside other golang projects for invoking the secrecy C++ library.
- **Subpackage `cli`**: Provides command line interface to control a local computing party, secrecy deployed as microservice in other party, or a computing parties cluster to start a new secure execution locally.
- **Subpackage `api`**: Built on Gin, it provides a RESTful API interface that can deploy secrecy as a microservice waiting for incoming computation requests.


## Getting started
Let's install the prerequisites and the package as follows:

- Install golang from here: https://go.dev/doc/install 
- Setup the secrecy framework inside the directory `~/.secrecy-go/secrecy`.
- Clone this package using git `git clone git@github.com:CASP-Systems-BU/secrecy-go.git`.

We can show help using the following command:
```
ubuntu@machine:~/secrecy-go$ go run main.go --help
NAME:
   secrecy-go - A go package to manage the secrecy MPC framework across the computing parties.

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
   configure  Configures the secrecy-go package defaults (~/.secrecy-go/config.env).
   local      Manages a local computing party and sends requests to it.
   serve      Starts the computing party service.
   client     Sends API requests to the computing party service.
   execute    Runs a secure multi-party computation across the specified computing parties.
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --host value   The host address for the computing party service. (default: "http://localhost:8080")
   --port value   The host of the computing party service. (default: "8080")
   --help, -h     show help
   --version, -v  print the version
```

## Running an example
After setting up secrecy and secrecy-go, we can do the following from inside the secrecy-go package.
- Configure computing parties
```
go run main.go configure party --id apollo --address localhost
go run main.go configure party --id starbuck --address localhost
go run main.go configure cluster --id cluster --parties appollo --parties starbuck
```
- Execute the secrecy MPC file.
```
go run main.go local --file ./examples/billionaires.cpp --cluster cluster 
```