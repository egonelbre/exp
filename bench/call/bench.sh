#!/bin/bash

declare -a versions=(
	"go1.7.6"
	"go1.8.7"
	"go1.9.7"
	"go1.10.8"
	"go1.11.13"
	"go1.12.17"
	"go1.13.15"
	"go1.14.15"
	"go1.15.15"
	"go1.16.15"
	"go1.17.13"
	"go1.18.10"
	"go1.19.12"
	"go1.20.7"
	"go1.21.0"
)

echo "# Installing Go versions"

for ver in "${versions[@]}"
do
	echo "go install golang.org/dl/${ver}@latest"
	echo "${ver} download"
done

echo "# Benchmarking"

for ver in "${versions[@]}"
do
	echo "${ver} test -bench . -count 10 | tee ${ver}.log"
done