#!/bin/bash

BUILD_DIR=maze-gen
if [ -d $BUILD_DIR ];then
	rm -rf $BUILD_DIR
fi

# Initialize the maze-gen module
if [ ! -f "go.mod" ];then
	go mod init maze-gen
fi
go build
