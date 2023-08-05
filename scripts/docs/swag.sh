#!/bin/sh

ROOT=$(dirname "$(echo "$0" | grep -E "^/" -q && echo "$0" || echo "$PWD/${0#./}")")

IMAGE="portal/backend/docs/swag/compile:v1.18.12"

[ -z "$1" ] && echo "[ERR] Не переданы аргументы" &&
  echo "Пример: 'swag fmt -d ./ && swag init --parseDependency --parseInternal \\" &&
  echo " -g ./cmd/api/main.go --outputTypes yaml,go --output ./docs'" && exit 1

HAS=$(docker image ls --filter=reference="$IMAGE" -q) || exit 1
if [ -z "$HAS" ]; then
  docker build -f "${ROOT}/Dockerfile" -t $IMAGE "$ROOT" || exit 1
fi

docker run -it --rm \
  --user "$(id -u):$(id -g)" \
  -w /app \
  -e GOCACHE="/tmp/.cache/go-build" \
  -v "${ROOT}/../..:/app" \
  "${IMAGE}" \
  sh -c "$*"
