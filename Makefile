
APP_NAME="corpos-christie"
APP_ID="lucasnoga.corpos-christie"
APP_VERSION="2.0.0"
APP_BUILD=2

.PHONY: package build-setup build-linux build-windows build-mac
.PHONY: run run-console

all: clean package

clean: 
	@echo "Cleaning build/ folder"
	@rm -rf build/*

# Package executables for all OS after built-in
package: build-setup build-linux build-windows
	@echo Building executable ${APP_NAME}...
	@./package.sh ${APP_VERSION}
	@echo "${APP_NAME} built"

# Test package app
package-test: package
	@echo Unzip build ${APP_NAME}...
	@unzip build/linux/linux-${APP_NAME}-${APP_VERSION}.zip -d build/test
	@pwd
	@./build/test/${APP_NAME}

# Setup 
build-setup:
	@echo "Creating build directory"
	@rm -rf fyne-cross
	@mkdir -p build/
	cp -r resources build/resources
	

# Build executable for Linux
build-linux: build-setup
	@echo "Build for Linux"
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

# Run app in console
run-console:
	go run . --console

# Run test all
test:
	go test ./...
	
# See doc
doc:
	go doc

# list all target in makefile
list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v '\.PHONY'