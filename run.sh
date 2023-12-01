#!/bin/bash

set -euo pipefail

sep() {
    echo "================="
}

go build -o bin/day01 cmd/day01/*.go
sep
echo "Running Day 01, Part 1"
sep
./bin/day01 < "cmd/day01/calibration-input.txt"

sep
