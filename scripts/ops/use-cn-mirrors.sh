#!/usr/bin/env bash

set -euo pipefail

GO_PROXY_URL="https://goproxy.cn,direct"
GO_SUMDB_VALUE="sum.golang.google.cn"
NPM_REGISTRY_URL="https://registry.npmmirror.com"

if ! command -v go >/dev/null 2>&1; then
  printf 'go command not found\n' >&2
  exit 1
fi

if ! command -v npm >/dev/null 2>&1; then
  printf 'npm command not found\n' >&2
  exit 1
fi

go env -w GOPROXY="${GO_PROXY_URL}"
go env -w GOSUMDB="${GO_SUMDB_VALUE}"
npm config set registry "${NPM_REGISTRY_URL}"

printf 'Go proxy set to %s\n' "${GO_PROXY_URL}"
printf 'Go sumdb set to %s\n' "${GO_SUMDB_VALUE}"
printf 'npm registry set to %s\n' "${NPM_REGISTRY_URL}"
