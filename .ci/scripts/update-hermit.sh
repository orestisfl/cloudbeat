#!/bin/bash
set -euxo pipefail

echo test1
bin/hermit list
bin/hermit upgrade aws-iam-authenticator awscli elastic-package gcloud golangci-lint jq just mage opa pre-commit rain regal shellcheck shfmt yq
ls -lah bin
bin/hermit list
echo test2
