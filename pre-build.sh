#!/bin/bash

# A simple shell script to pre-build and pack binaries for this project. 
# Used to make it easier to include pre-build binaries to release assets on github.

echo "Compiling windows-amd64 ..."
GOOS=windows GOARCH=amd64 go build -o ./bin/michael.windows-amd64 .

echo "Compiling darwin-amd64 ..."
GOOS=darwin GOARCH=amd64 go build -o ./bin/michael.darwin-amd64 .

echo "Compiling linux-amd64 ..."
GOOS=linux GOARCH=amd64 go build -o ./bin/michael.linux-amd64 .

echo "Running tar -czf on binaries"
for f in $(find ./bin/*); do tar -czf $f.tar.gz $f; done;
