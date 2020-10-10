#!/bin/bash
set -e

GOOS=darwin go build -o xlsx-fake-data
upx xlsx-fake-data
GOOS=windows go build -o xlsx-fake-data.exe
upx xlsx-fake-data.exe
