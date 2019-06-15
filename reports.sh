#!/usr/bin/env bash

mkdir -p reports/{golang,js}
touch reports/golang/coverage.out reports/golang/test-report.out reports/golang/vet.out
touch reports/js/lcov.dat
touch c.out