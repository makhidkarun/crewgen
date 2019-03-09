#!/bin/bash

# Run the same build commands

# Starting go vet
go vet main.go

# Check return
if [ $? == 0 ]
then
	# Build
	time go build -o bin/crewgen main.go 
	#run
	bin/crewgen
else
  echo "go vet failed."
fi
