#!/usr/bin/env bash

set -euo pipefail

log() {
  printf '[deploy] %s\n' "$1"
}

fail() {
  printf '[deploy] ERROR: %s\n' "$1" >&2
  exit 1
}

run_remote() {
  local message="$1"
  local command="$2"

  log "$message"
  ssh "${SSH_OPTS[@]}" "${REMOTE}" "cd '${DIR}' && printf '[remote] %s\\n' '$message' && ${command}" \
    || fail "$message failed"
}

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

log "ensuring remote directory exists: ${REMOTE}:${DIR}"
ssh "${SSH_OPTS[@]}" "${REMOTE}" "mkdir -p '${DIR}'" || fail "failed to create remote directory ${DIR}"

log "syncing project files to remote host"
rsync -az "${EXCLUDES[@]}" -e "ssh -p ${PORT} -o BatchMode=yes -o StrictHostKeyChecking=accept-new" \
  "${ROOT_DIR}/" "${REMOTE}:${DIR}/" || fail "rsync upload failed"

log "updating remote script permissions"
ssh "${SSH_OPTS[@]}" "${REMOTE}" "chmod +x '${DIR}/scripts/ops/'*.sh" || fail "failed to update remote script permissions"

run_remote "rebuilding and starting remote containers" "docker compose up -d --build"
run_remote "showing remote container status" "docker compose ps"

log "waiting for remote services to become healthy"
ssh "${SSH_OPTS[@]}" "${REMOTE}" 'bash -s' -- "${DIR}" <<'EOF' || fail "remote health check failed"
set -euo pipefail

DIR="$1"
cd "$DIR"

printf '[remote] waiting for service health checks\n'

services=$(docker compose config --services)
deadline=$((SECONDS + 90))

for service in $services; do
  printf '[remote] checking service: %s\n' "$service"

  while :; do
    container_id=$(docker compose ps -q "$service")

    if [ -z "$container_id" ]; then
      printf '[remote] ERROR: service %s has no container id\n' "$service" >&2
      exit 1
    fi

    state=$(docker inspect --format '{{.State.Status}}' "$container_id")
    health=$(docker inspect --format '{{if .State.Health}}{{.State.Health.Status}}{{else}}none{{end}}' "$container_id")

    printf '[remote] service=%s state=%s health=%s\n' "$service" "$state" "$health"

    if [ "$state" != "running" ]; then
      printf '[remote] ERROR: service %s is not running\n' "$service" >&2
      docker compose ps "$service" >&2 || true
      exit 1
    fi

    if [ "$health" = "healthy" ] || [ "$health" = "none" ]; then
      break
    fi

    if [ "$health" = "unhealthy" ] || [ $SECONDS -ge $deadline ]; then
      printf '[remote] ERROR: service %s failed health check\n' "$service" >&2
      docker compose ps "$service" >&2 || true
      docker compose logs --tail=100 "$service" >&2 || true
      exit 1
    fi

    sleep 5
  done
done
EOF

log "deployment completed successfully on ${REMOTE}:${DIR}"
