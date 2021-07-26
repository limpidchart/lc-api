#!/usr/bin/env bash
#
# Script to generate swagger spec from source.

set -eu
set -o pipefail

package="./cmd/lc-api/main.go"

go generate "${package}"