
EXE="corpos-christie"
VERSION="0.0.9"

.PHONY: build run runconsole rungui

all: build

# Building executable
build:
	@echo Building executable ${EXE}...
	@./build.sh ${VERSION}
	@echo "${EXE} built"

# Test building
testbuild:
	@echo Unzip build ${EXE}...
	@unzip build/Corpos-christie-${VERSION}.zip -d build
	@pwd
	@./build/${EXE}

# Run program
run:
	go run main.go

# Run program in console mode
runconsole:
	go run main.go --console

rungui:
	go run main.go --gui

# Run test all
test:
	go test ./...
	
# See doc
doc:
	go doc

# Docker build
docker-build: 
	go build
	docker build -t ${EXE}:${VERSION} .

# Docker run
docker-run:
	docker run -it --rm --name ${EXE} ${EXE}:${VERSION}

# list all target in makefile
list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v '\.PHONY'