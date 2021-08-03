PROJECTNAME := $(shell basename "$(PWD)")

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent


## crosscompile: cross platform compilation.
crosscompile:
	@echo "Cross Platform Compilation:"
	@echo "GOOS=freebsd GOARCH=386 go build -o bin/"$(PROJECTNAME)"-freebsd main.go"
	GOOS=freebsd GOARCH=386 go build -o bin/"$(PROJECTNAME)"-freebsd main.go
	@echo "GOOS=linux GOARCH=386 go build -o bin/"$(PROJECTNAME)"-linux main.go"
	GOOS=linux GOARCH=386 go build -o bin/"$(PROJECTNAME)"-linux main.go
	@echo "GOOS=windows GOARCH=386 go build -o bin/"$(PROJECTNAME)"-windows main.go"
	GOOS=windows GOARCH=386 go build -o bin/"$(PROJECTNAME)"-windows main.go

## compile: execute compilation for the current platform and architecture.
compile:
	@echo "go build -o bin/"$(PROJECTNAME)" main.go"
	go build -o bin/"$(PROJECTNAME)" main.go

.PHONY: help
all: help
help: Makefile
	@echo
	@echo "Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo