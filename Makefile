
APP_NAME="corpos-christie"
APP_ID="lucasnoga.corpos-christie"
APP_VERSION="4.0.0"

.PHONY: run

all: clean

clean: 
	@echo "Cleaning build/ folder"
	@rm -rf build/*

setup: icon
	@echo "Build ${APP_NAME} - ${APP_VERSION}"
	@cp ./icon.png ./build/appicon.png
	@cp ./icon.ico ./build/windows/icon.ico
	@echo "[X] Icon generated"
	@wails build
	
run:
	wails dev

doctor:
	wails doctor

icon:
	winicon g -sizes 16,32,48,64,128,256 ./icon.png

icon-info:
	winicon i ./build/windows/icon.ico

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