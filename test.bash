#!/usr/bin/env bash

set -e

if [ ! -f test.bash ]; then
	echo 'test.bash must be run from $GOROOT/src' 1>&2
	exit 1
fi

eval $(go env)

export GOROOT

GOPATH=$(pwd)
export GOPATH

go test interswitch