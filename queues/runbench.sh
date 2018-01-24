#!/bin/bash

# the tests must use fixed benchIterations
# this makes it avoid infinite-loop
#  -benchtime=0.1s
go test -cpu=1,2,5,10,100 -benchtime=0.1s -timeout=1h -bench . > benchmark.txt