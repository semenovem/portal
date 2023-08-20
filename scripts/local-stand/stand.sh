#!/bin/bash

ROOT=$(dirname "$(echo "$0" | grep -E "^/" -q && echo "$0" || echo "$PWD/${0#./}")")
. "${ROOT}/lib.sh" || exit 1

PROPS_FILE="${ROOT}/props.env"
LOCAL_OVERRIDE_PROPS_FILE="${ROOT}/.override-props.env"

ARG_STAND_CLEAR=           # Отчистить БД / redis и тд
export __ARG_MODE_DEBUG__= # режим отладки go приложения
ARG_HELP=                  # флаг запроса справки

OPER=
STAND_OPER=
WANT_STAND_OPER=

export __REPEAT__=
LOGS=
ERR=

help() {
  info "use: sh stand.sh  - development stand management"
  info "commands: [up|down [-clear]] [command [-option]]"
  info "    up [-clear]        - start docker compose (DB/redis etc. Look in the docker-compose.yml)"
  info "    down               - stop docker compose"
  info ""
  info "    papi | api-portal  - run api-portal"
  info "    audit              - run audit"
  info ""
  info "    ps                 - docker compose ps"
  info "    ports              - show exposed ports"
  info "    redis-adm          - start redis admin gui, if not yet"
  info "    pg-adm             - start postgres admin gui, if not yet"
  info ""
  info "    elk-up    - start elk stack"
  info "    elk-down  - stop elk stack"
  info ""
  info "options:"
  info "    -debug        - golang application debug mode. Work for api_clients, api_admins etc"
  info "    -clear        - cleaning up the database (work with up)"
  info "    -logs         - "
}

func_check_and_create_override_props &&
  func_apply_whole_env_file "$PROPS_FILE" &&
  func_apply_env_file "$LOCAL_OVERRIDE_PROPS_FILE" || exit 1

# ------------------------------------------------
# --------------      arguments     --------------
# ------------------------------------------------
PREV=
for p in "$@"; do
  if [ -n "$PREV" ]; then
    case "$PREV" in
    "-stand")
      CONNECT_TO_DB_STAND="$p"
      CONNECT_TO_REDIS_STAND="$p"
      CONNECT_TO_S3_STAND="$p"
      ;;
    "-db") CONNECT_TO_DB_STAND="$p" ;;
    "-redis") CONNECT_TO_REDIS_STAND="$p" ;;
    "-s3") CONNECT_TO_S3_STAND="$p" ;;
    *)
      ERR=1
      err "Unknown argument '${PREV}' '${p}'"
      ;;
    esac

    PREV=
    continue
  fi

  case "$p" in
  # stand operations
  "up") STAND_OPER="up" ;;
  "down") STAND_OPER="down" ;;
  "ps") STAND_OPER="ps" ;;

  "elk-up") OPER="elk-up" ;;
  "elk-down") OPER="elk-down" ;;

    # service operations
  "papi" | "api-portal") OPER="api-portal" ;;
  "audit") OPER="audit" ;;
  "curl") OPER="curl" ;;

  "redis-adm")
    __REDIS_ADMIN_GUI_REPLICAS__=1
    WANT_STAND_OPER=up
    ;;

  "pg-adm")
    __PG_ADMIN_GUI_REPLICAS__=1
    WANT_STAND_OPER=up
    ;;

  "ports" | "port")
    func_show_expose_envs
    exit 0
    ;;

  "props-env-file")
    TMP=$(mktemp) || exit 1
    env | grep -iEo "^__.+" | sort >"$TMP" || exit 1
    echo "$TMP"
    exit 0
    ;;

    # params
  "-stand") PREV="-stand" ;;
  "-db") PREV="-db" ;;
  "-redis") PREV="-redis" ;;
  "-s3") PREV="-s3" ;;
  "-clear") ARG_STAND_CLEAR=1 ;;
  "-logs") LOGS=1 ;;
  "-r" | "-repeat") __REPEAT__=1 ;;
  "-debug") __ARG_MODE_DEBUG__=1 ;;
  "-h" | "-help") ARG_HELP=1 ;;
  *)
    ERR=1
    err "Unknown argument '${p}'"
    ;;
  esac
done

unset PREV p

[ -z "$STAND_OPER" ] && [ -n "$WANT_STAND_OPER" ] && STAND_OPER="$WANT_STAND_OPER"
[ -z "$OPER" ] && [ -z "$STAND_OPER" ] && [ -z "$WANT_STAND_OPER" ] && ARG_HELP=1
[ -n "$ARG_HELP" ] && help && exit 0

which "docker" >/dev/null
[ $? -ne 0 ] && err "docker needs to be installed" && ERR=1
[ -n "$ERR" ] && help && exit 1

# ------------------------------------------------
# -----------   docker-compose   -----------------
# ------------------------------------------------
case "$STAND_OPER" in
"down") func_stand_down || exit 1 ;;

"up")
  if [ -n "$ARG_STAND_CLEAR" ]; then
    func_stand_down && sleep 1 && func_clear || exit 1
  fi

  func_create_network && func_build_if_not_exist_s3_mc_image || exit 1

  OPTS="$([ -n "$ARG_STAND_CLEAR" ] && echo "--force-recreate" || echo "--no-recreate")"
  if [ -z "$LOGS" ] || [ -n "$OPER" ]; then
    OPTS="${OPTS} --detach"
  fi

  echo ">>>> \$LOGS = $LOGS"
  echo ">>>> \$OPER = $OPER"

  docker compose \
    -p "$__STAND_NAME__" \
    --project-directory "$ROOT" \
    -f "${ROOT}/docker-compose.yml" \
    up --quiet-pull $OPTS
  ;;

"ps")
  docker compose \
    -p "$__STAND_NAME__" \
    --project-directory "$ROOT" \
    -f "${ROOT}/docker-compose.yml" \
    ps
  ;;
esac

# ------------------------------------------------
# -----------      services      -----------------
# ------------------------------------------------
if [ -n "$OPER" ]; then
  func_build_if_not_exist_dlv_image || exit 1
  func_create_network || exit 1
fi

case "$OPER" in
"api-portal")
  CMD="dlv debug /debugging/cmd/api/main.go --headless --listen=:40000 --api-version=2 --accept-multiclient --output /tmp/__debug_bin"
  [ -z "$__ARG_MODE_DEBUG__" ] && CMD="$(func_run_cmd "go run /debugging/cmd/api/main.go")"

  docker run -it --rm \
    --name "api-portal" \
    --hostname "api-portal" \
    --network "$__NET__" \
    -p "${__API_PORTAL_REST_PORT_EXPOSE__}:8080" \
    -p "${__API_PORTAL_GRPC_PORT_EXPOSE__}:9090" \
    -p "${__API_PORTAL_DEBUGGING_PORT_EXPOSE__}:40000" \
    -w "/debugging" \
    -v "${ROOT}/../../:/debugging:ro" \
    --env-file "${ROOT}/../../deployments/local.env" \
    -e "GRPC_AUDIT_CLIENT_HOST=audit:9090" \
    "$(func_get_work_image)" bash -c "$CMD"
  ;;

"audit")
  CMD="dlv debug /debugging/cmd/api/main.go --headless --listen=:40000 --api-version=2 --accept-multiclient --output /tmp/__debug_bin"
  [ -z "$__ARG_MODE_DEBUG__" ] && CMD="$(func_run_cmd "go run /debugging/cmd/audit/main.go")"

  docker run -it --rm \
    --name "audit" \
    --hostname "audit" \
    --network "$__NET__" \
    -p "${__AUDIT_GRPC_PORT_EXPOSE__}:9090" \
    -p "${__AUDIT_DEBUGGING_PORT_EXPOSE__}:40000" \
    -w "/debugging" \
    -v "${ROOT}/../../:/debugging:ro" \
    --env-file "${ROOT}/../../deployments/audit/local.env" \
    "$(func_get_work_image)" bash -c "$CMD"
  ;;

"elk-up")
  docker compose \
    -p elk-st \
    --project-directory "$ROOT" \
    -f "${ROOT}/docker-elk.yml" \
    up
  ;;

"elk-down")
  #    has=$(docker compose ls "--filter=name=${__STAND_NAME__}" -q) || return 1
  #    [ -z "$has" ] && return 0

  docker compose \
    -p elk-st \
    --project-directory "$ROOT" \
    -f "${ROOT}/docker-elk.yml" \
    down
  ;;

\
  "curl")
  HAS=$(docker images --filter=reference="$__DOCKER_CURL_IMAGE__" -q) || exit 1
  if [ -z "$HAS" ]; then
    docker build -f "${ROOT}/docker-files/curl.dockerfile" -t "$__DOCKER_CURL_IMAGE__" "$ROOT" || exit 1
  fi

  docker run -it --rm --network "$__NET__" -w /app "$__DOCKER_CURL_IMAGE__" bash
  ;;
esac
