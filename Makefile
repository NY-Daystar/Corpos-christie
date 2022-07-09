
EXE="corpos-christie"
VERSION="1.0.0"

.PHONY: build buildLinux buildWindows buildMac
.PHONY: run runconsole rungui

all: clean build

clean: 
	@echo "Cleaning build/ folder"
	@rm -f build/*

# Building executable for all OS
build: build-linux build-windows build-mac
	@echo Building executable ${EXE}...
	@./build.sh ${VERSION}
	@echo "${EXE} built"

# build executable for linux 
build-linux:
	@echo "Build for Linux"
	GOOS=linux GOARCH=amd64 go build -o build/linux-corpos-christie

# build executable for Windows 
build-windows:
	@echo "Build for Windows"
	GOOS=windows GOARCH=amd64 go build -o build/windows-corpos-christie.exe

# build executable for MacOS
build-mac:
	@echo "Build for Mac"
	GOOS=darwin GOARCH=amd64 go build -o build/mac-corpos-christie

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
	go build -o ${EXE}
	docker build -t ${EXE}:${VERSION} . && rm ${EXE}

# Docker run
docker-run:
	docker run -it --rm --name ${EXE} ${EXE}:${VERSION}

# list all target in makefile
list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v '\.PHONY'