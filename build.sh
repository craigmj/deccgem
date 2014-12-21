#!/bin/bash
set -e
export GOPATH=`pwd`
for d in "github.com/craigmj/commander";
	do 
	if [ ! -d src/$d ]; then
		go get $d
	fi
done
if [ ! -d bin ]; then mkdir bin; fi

go build -o bin/deccgem src/cmd/deccgem.go

d=`which deccgem` || true
if [ ! -z "$d" ]; then
	sudo rm $d
fi
sudo ln -s `pwd`/bin/deccgem /usr/bin/
