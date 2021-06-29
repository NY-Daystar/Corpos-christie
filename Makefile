
EXE="corpos-christie"

all: build

# Building executable
.PHONY: build
build:
	@echo Building executable ${EXE}...
	@go build
	@echo "${EXE} built"

# Run program
.PHONY: run
run:
	go run main.go

# Run test all
test:
	go test ./...
	
# See doc
doc:
	go doc