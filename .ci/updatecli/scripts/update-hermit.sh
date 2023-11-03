#!/bin/bash
set -euxo pipefail

bin/hermit upgrade \
    aws-iam-authenticator \
    awscli \
    elastic-package \
    gcloud \
    golangci-lint \
    jq \
    just \
    mage \
    opa \
    pre-commit \
    rain \
    regal \
    shellcheck \
    shfmt \
    yq
git status
