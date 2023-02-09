build:
	@echo "Building Michael in ./bin..."
	@go build -o bin/michael .
	@echo "Finished!"

install:
	@echo "Installing Michael with 'go install'"
	@go install .
	@echo "Finished!"
