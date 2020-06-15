# boot
*Boot* is a kind of meta module, to include a typical set of dependencies all at once and provide code 
generators and conventions. It helps you to bootstrap your project setup and to use some reasonable default
conventions.

## tutorial
To complete this tutorial, you have to ensure the following prerequisites:
* the latest [Go](https://golang.org) version, at least Go 1.14
* a Linux or MacOS development host (Windows may work, but is not fun)
* *make* support and a commandline of your choice

```bash
# ensure, that your are not in GOPATH
cd ~/myprojects

# create your project folder
mkdir mymonolith && cd mymonolith

# create your go server module
go mod init mycompany.com/mymonolith

# create a cmd which is always called to perform code (re)-generation
mkdir -p cmd/gen

# add the boot dependency
go get github.com/golangee/boot

# create the boot-generate file
cat > cmd/gen/main.go << EOL
//go:generate go run main.go
package main

import (
    "github.com/golangee/boot"
)

func main() {
    err := boot.Generate()
    if err != nil{
        panic(err)
    }
}

EOL

gofmt -w cmd/gen/main.go

```

Now it is time to invoke the first *generate* run. This will detect an empty project and will create a default
database repository, and a demonstration forms module:

```bash
# run the first generate by hand
go generate ./...    
```

Afterwards also *makefiles* have been created. Your best friends are:

```bash
make help # which targets are available?
make run # the 'play' button
```

This setup is known to work nicely with the [Goland IDE](https://www.jetbrains.com/go/) but others like VSCode may also work.
Just open the server project. Keep in mind, that intentionally you cannot mix data structures between client
and server. Try to rely solely on your REST definition - a client is always regenerated.