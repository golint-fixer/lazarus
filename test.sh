#!/usr/bin/env bash

# Runs go tests in all subpackages and pushes details out to codecov via travis-ci

set -e
echo "" > coverage.txt

for d in $(find ./* -maxdepth 10 -type d | grep -v vendor); do
    if ls $d/*.go &> /dev/null; then
        go test -coverprofile=profile.out -covermode=atomic $d
        if [ -f profile.out ]; then
            cat profile.out >> coverage.txt
            rm profile.out
        fi
    fi
done
