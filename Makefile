.PHONY: build clean

all: clean run

run:
	@echo "running go run main.go..."
	@go run main.go

clean:
	@echo "cleaning..."
	@rm -rf ./output
