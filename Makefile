
APP_NAME="corpos-christie"
APP_ID="lucasnoga.corpos-christie"
APP_VERSION="3.2.0"
APP_BUILD=2

.PHONY: package build-setup build-linux build-windows build-mac
.PHONY: run

all: clean package

clean: 
	@echo "Cleaning build/ folder"
	@rm -rf build/*

# Package executables for all OS after built-in
package: clean build-setup build-linux build-windows
	@echo Building executable ${APP_NAME}...
	@./package.sh ${APP_VERSION}
	@echo "${APP_NAME} built"

# Setup 
build-setup:
	@echo "Creating build directory"
	@rm -rf fyne-cross
	@mkdir -p build/
	cp -r resources build/resources
	

# Build executable for Linux
build-linux: build-setup
	@echo "Build for Linux & Mac"
	@fyne-cross linux -arch=amd64 --app-id=${APP_ID} --app-build=${APP_BUILD} --app-version=${APP_VERSION}
	@echo "Move executable into build folder"
	cp fyne-cross/bin/linux-amd64/Corpos-Christie build/linux-${APP_NAME}
	@chmod +x build/linux-${APP_NAME}

# Build executable for Windows
build-windows: build-setup
	@echo "Build for Windows"
	@fyne-cross windows -arch=amd64 --app-id=${APP_ID} --app-build=${APP_BUILD} --app-version=${APP_VERSION}
	@echo "Move executable into build folder"
	cp fyne-cross/bin/windows-amd64/Corpos-Christie.exe build/windows-${APP_NAME}.exe
	
	
# Build executable for MacOS
build-mac:
	@echo "Build for MacOS"
	@GOOS=darwin GOARCH=amd64 go build -o build/mac-${APP_NAME}

# Run app
run:
	go run .

# get test coverage
coverage:
	go mod download golang.org/x/tools
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out

# run sca analysis
check_cve:
	govulncheck ./...

# Run test all
test:
	go test ./...
	
# See doc
doc:
	go doc

# list all target in makefile
list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v '\.PHONY'