#!/bin/bash
set -euxo pipefail

BEATS_VERSION=${1:?Missing version argument}

go version
go env
env

go get "github.com/elastic/beats/v7@$BEATS_VERSION"
go mod tidy

cat go.mod
git diff
