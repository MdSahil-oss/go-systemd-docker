#! /usr/bin/bash

set -e

OUTPUT_FILE_NAME=sysd

go build -o $OUTPUT_FILE_NAME ./cmd/sysd

export PATH=${PATH}:${PWD}/${OUTPUT_FILE_NAME}
