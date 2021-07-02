
EXE="corpos-christie"
VERSION="0.0.7"

all: build

# Building executable
.PHONY: build
build:
	@echo Building executable ${EXE}...
	@./build.sh
	@echo "${EXE} built"

# Test building
testbuild:
	@echo Unzip build ${EXE}...
	@unzip build/Corpos-christie-${VERSION}.zip -d build
	@pwd
	@./build/${EXE}

# Run program
.PHONY: run
run:
	go run main.go

# Run program in console mode
.PHONY: runconsole
runconsole:
	go run main.go --console

.PHONY: rungui
rungui:
	go run main.go --gui

# Run test all
test:
	go test ./...
	
# See doc
doc:
	go doc

# list all target in makefile
list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v '\.PHONY'