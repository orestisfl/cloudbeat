#!/bin/bash
set -euxo pipefail

echo test1
bin/hermit upgrade aws-iam-authenticator awscli elastic-package gcloud golangci-lint jq just mage opa pre-commit rain regal shellcheck shfmt yq
echo test2
