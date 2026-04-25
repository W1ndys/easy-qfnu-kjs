#!/usr/bin/env bash

set -euo pipefail

if [ "$#" -ne 4 ]; then
  printf 'Usage: %s HOST PORT USER DIR\n' "$0" >&2
  exit 1
fi

HOST="$1"
PORT="$2"
USER="$3"
DIR="$4"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "${SCRIPT_DIR}/../.." && pwd)"
REMOTE="${USER}@${HOST}"

SSH_OPTS=(
  -p "${PORT}"
  -o BatchMode=yes
  -o StrictHostKeyChecking=accept-new
)

EXCLUDES=(
  --exclude .git
  --exclude .monkeycode
  --exclude node_modules
  --exclude frontend/node_modules
  --exclude web
  --exclude tmp
)

ssh "${SSH_OPTS[@]}" "${REMOTE}" "mkdir -p '${DIR}'"

rsync -az "${EXCLUDES[@]}" -e "ssh -p ${PORT} -o BatchMode=yes -o StrictHostKeyChecking=accept-new" \
  "${ROOT_DIR}/" "${REMOTE}:${DIR}/"

ssh "${SSH_OPTS[@]}" "${REMOTE}" "chmod +x '${DIR}/scripts/ops/'*.sh"

printf 'Deployed project files to %s:%s\n' "${REMOTE}" "${DIR}"
