#!/bin/bash
set -euo pipefail

# The workspace is bind-mounted and owned by a different uid than the container user
git config --global --add safe.directory "${GITHUB_WORKSPACE:-/github/workspace}"

export LD_BRANCH="${LD_BRANCH:-${GITHUB_HEAD_REF:-${GITHUB_REF_NAME:-}}}"
export LD_REVISION="${LD_REVISION:-${GITHUB_SHA:-}}"

if [ -n "${LD_OUT_DIR:-}" ]; then
    mkdir -p "$LD_OUT_DIR"
fi

exec ld-find-code-refs
