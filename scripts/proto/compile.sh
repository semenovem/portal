#!/bin/bash

BIN=$(dirname "$(echo "$0" | grep -E "^/" -q && echo "$0" || echo "$PWD/${0#./}")")

# Образ с установленным компилятором protoc
IMAGE="portal/proto/compile:1.1"

info() {
  echo "\033[0;32m[INFO] $*\033[0m"
}
err() {
  echo "\033[0;31m[ERR] $*\033[0m"
}
help() {
  info "proto compile: compile.sh"
  info "show versions: compile.sh ver"
}

show_versions() {
  info "Версии библиотек: "

  pipe() {
    while read -r data; do info "$data"; done
  }

  grep -iEo 'protoc-\d+\.\d' "${BIN}/Dockerfile" | pipe
  grep -iEo 'protoc-gen-go@v\d+\.\d+\.\d' "${BIN}/Dockerfile" | pipe
  grep -iEo 'protoc-gen-go-grpc@v\d+\.\d+\.\d' "${BIN}/Dockerfile" | pipe
  grep -iEo 'protoc-gen-grpc-gateway@v\d+\.\d+\.\d' "${BIN}/Dockerfile" | pipe
  grep -iEo 'protoc-gen-openapiv2@v\d+\.\d+\.\d' "${BIN}/Dockerfile" | pipe
  return 0
}

which "docker" 1>/dev/null
[ $? -ne 0 ] && echo "[ERR] Нужно установить docker" && exit 1

HAS=$(docker image ls --filter=reference="$IMAGE" -q) || exit 1
if [ -z "$HAS" ]; then
  info "нет образа '${IMAGE}'. Сейчас соберем"
  docker build --platform linux/amd64 -f "${BIN}/Dockerfile" -t $IMAGE "$BIN" || exit 1
fi

for p in "$@"; do
  case $p in
  "ver")
    show_versions
    exit 0
    ;;

  "buf")
    docker run -it --rm -w /app --platform linux/amd64 \
      --user "$(id -u):$(id -g)" \
      -v "${BIN}:/app" \
      "$IMAGE" \
      sh -c "buf generate"

    exit 0
    ;;

  *)
    [ -d "$p" ] && continue
    err "unknown arg '${p}'"
    exit 1
    ;;
  esac
done

CMD="protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=."
CMD="${CMD} --go-grpc_opt=paths=source_relative -I. audit_grpc/*.proto"

echo "cmd = $CMD"

docker run -it --rm -w /app --platform linux/amd64 \
  --user "$(id -u):$(id -g)" \
  -v "${BIN}/../../proto:/app" \
  "$IMAGE" sh -c "${CMD}" && info "Compile - ок"
