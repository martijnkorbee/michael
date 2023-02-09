build:
	@echo "Building Michael in ./bin..."
	@go build -o bin/michael .
	@echo "Finished!"

install:
	@echo "Installing Michael with 'go install'"
	@go install .
	@echo "Finished!"

pre-compile:
	@echo "Compiling windows-amd64 ..."
	@GOOS=windows GOARCH=amd64 go build -o ./bin/michael.windows-amd64 .
	@echo "Compiling darwin-amd64 ..."
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/michael.darwin-amd64 .
	@echo "Compiling linux-amd64 ..."
	@GOOS=linux GOARCH=amd64 go build -o ./bin/michael.linux-amd64 .
	@echo "Running tar -czf on binaries"
	@cd bin && for f in $(ls bin); do tar -czf $f.tar.gz $f; done;
