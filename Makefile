
EXE="corpos-christie"
VERSION="1.1.0"

.PHONY: build buildLinux buildWindows buildMac
.PHONY: run runconsole rungui

all: clean build

clean: 
	@echo "Cleaning build/ folder"
	@rm -rf build/*

# Building executable for all OS
build: build-linux build-windows build-mac
	@echo Building executable ${EXE}...
	@./build.sh ${VERSION}
	@echo "${EXE} built"

# build executable for Linux
build-linux:
	@echo "Build for Linux"
	@GOOS=linux GOARCH=amd64 go build -o build/linux-${EXE}

# build executable for Windows
build-windows:
	@echo "Build for Windows"
	@echo "Creating icon file"
	@rsrc -ico assets/logo.ico -o ${EXE}.syso
	@GOOS=windows GOARCH=amd64 go build -o windows-${EXE}.exe
	@mv ${EXE}.syso build/${EXE}.syso
	@mv windows-${EXE}.exe build/windows-${EXE}.exe
	

# build executable for MacOS
build-mac:
	@echo "Build for MacOS"
	@GOOS=darwin GOARCH=amd64 go build -o build/mac-${EXE}

# Test building
testbuild:
	@echo Unzip build ${EXE}...
	@unzip build/linux/${EXE}-${VERSION}.zip -d build/test
	@pwd
	@./build/test/${EXE}

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