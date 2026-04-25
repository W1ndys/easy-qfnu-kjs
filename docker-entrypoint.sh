#!/bin/sh
set -eu

mkdir -p /app/data /app/logs
chown -R app:app /app/data /app/logs

exec su-exec app:app /app/easy-qfnu-kjs
