#!/usr/bin/env bash

set -o errexit
set -o pipefail

log() (
    echo >&2 "$*"
)

panic() (
    log "$*"
    exit 1
)

print_usage() (
    set +o xtrace
    script_name="$(basename "${BASH_SOURCE[0]}")"
    log "Usage: ${script_name} [-v] [-h] <command> [<arg>...]
Global Options:
  -v, --verbose    Enable verbose output
  -h, --help       Print this help and exit
Commands:
  list
  create   <name>
  delete   <name>
  start    <name>
  stop     <name>
  restart  <name>
  logs     <name>
  status   <name>
  validate <name>
  runas    <name> <arg>..."
)

assert_root() (
    if [[ "$(id -u)" -ne 0 ]]; then
        panic "Run me as root!"
    fi
)

sanitize() (
    echo "$*" | inline-detox
)

ping_docker_socket() (
    sock_path=$(realpath "$1")
    if [[ ! -S ${sock_path} ]]; then
        log "Warning: No socket found at ${sock_path}"
        return 1
    fi
    response=$(curl --silent --max-time 1 --unix-socket "${sock_path}" "http://localhost/_ping")
    if echo "${response}" | grep --silent --fixed-strings "OK"; then
        return 0
    else
        log "Warning: Ping not ok for socket ${sock_path}"
        return 1
    fi
)

basedir_init() (
    mkdir -p "${basedir}"
    cd "${basedir}"
    if ! git status &>/dev/null; then
        git init --quiet
    fi
)

run_in_env() (
    cd "${service_dir}"
    # some env vars need to be provided as command context
    local env
    env=()
    env+=("PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin")
    env+=("XDG_RUNTIME_DIR=/run/user/$(id -u "${service}")")
    env+=("DOCKER_HOST=unix:///run/user/$(id -u "${service}")/docker.sock")
    env+=("DBUS_SESSION_BUS_ADDRESS=unix:path=/run/user/$(id -u "${service}")/bus")
    # run as env user, in user home
    sudo --set-home --user "${service}" sh -c "cd; ${env[*]} $*"
)

# COMMANDS

list() (
    cd "${basedir}"
    ls
)

create() (
    if id "${service}" &>/dev/null; then panic "User exists: ${service}"; fi
    if [[ -e "${service_dir}" ]]; then panic "Path exists: ${service_dir}"; fi

    log "Creating service ${service} in ${service_dir}"

    addgroup \
        --system \
        "${service}"

    adduser \
        --system \
        --uid "$(getent group "${service}" | cut -d: -f3)" \
        --ingroup "${service}" \
        --disabled-password \
        --home "${service_dir}" \
        "${service}"

    # add sub-ids
    # https://rootlesscontaine.rs/getting-started/common/subuid/
    grep -qxF "${service}:100000:65536" /etc/subuid || echo "${service}:100000:65536" >>/etc/subuid
    grep -qxF "${service}:100000:65536" /etc/subgid || echo "${service}:100000:65536" >>/etc/subgid

    # linger allows the user to start long-running services without being logged in
    # https://rootlesscontaine.rs/getting-started/common/login/#optional-start-the-systemd-user-session-on-boot
    loginctl enable-linger "${service}"
    run_in_env "systemctl --user start dbus"
    log "Sleeping until bus socket is ready"
    while [[ ! -S "/run/user/$(id -u "${service}")/bus" ]]; do
        sleep 1
        sleep_count=$((sleep_count + 1))
        if [[ sleep_count -gt 10 ]]; then
            panic "Timeout waiting for bus socket"
        fi
    done

    # install and enable a rootless docker daemon
    # https://docs.docker.com/engine/security/rootless/#install
    run_in_env "dockerd-rootless-setuptool.sh install"
    run_in_env "systemctl --user start docker"
    run_in_env "systemctl --user enable docker"

    # create service skeleton
    mkdir -p "${service_dir}/.config/systemd/user"
    cp --verbose "/usr/local/share/eigenstack/example-service/"* "${service_dir}"
    cp "/usr/local/share/eigenstack/compose.service.template" "${service_dir}/.config/systemd/user/compose.service"
    sed -i "s/@SERVICE_NAME@/${service}/g" "${service_dir}/.config/systemd/user/compose.service"
    # store init state in git
    # TODO: standalone commit wrapper command
    #cd "${service_dir}"
    #git add "${service_dir}"
    #git commit --author "Bot <eigenstack@$(hostname --fqdn)>" --message "Initialized service directory ${service}️️" --quiet

    # ensure correct file ownership
    chown -R "${service}:${service}" "${service_dir}"

    # reload daemon
    cd "${service_dir}"
    run_in_env "systemctl --user daemon-reload"

    validate

    log "Finished initializing docker env ${service} in ${service_dir}"
)

delete() (
    log "NOT YET IMPLEMENTED"
)

start() (
    run_in_env "systemctl --user daemon-reload"
    run_in_env "systemctl --user start compose.service"
    run_in_env "systemctl --user enable compose.service"
)

stop() (
    run_in_env "systemctl --user daemon-reload"
    run_in_env "systemctl --user stop compose.service"
    run_in_env "systemctl --user disable compose.service"
)

restart() (
    run_in_env "systemctl --user daemon-reload"
    run_in_env "systemctl --user stop compose.service"
    run_in_env "systemctl --user start compose.service"
)

logs() (
    run_in_env "docker compose logs"
)

status() (
    run_in_env "systemctl --user daemon-reload"
    run_in_env "systemctl --user status compose.service"
    run_in_env "docker compose ls --all"
    run_in_env "docker compose ps --all"
    run_in_env "docker compose top"
)

validate() (
    log "Validating service ${service} in ${service_dir}"
    valid="yes"

    if id "${service}" &>/dev/null; then
        log "✓️ service user"
    else
        valid="no"
        log "✗ service user ( ${service} )"
    fi

    if [[ -d "${service_dir}/" ]]; then
        log "✓️ service directory"
    else
        valid="no"
        log "✗ service directory ( ${service_dir}/ )"
    fi

    if [[ -f "${service_dir}/docker-compose.yaml" ]]; then
        log "✓️ docker-compose.yaml"
    else
        valid="no"
        log "✗ docker-compose.yaml ( ${service_dir}/docker-compose.yaml )"
    fi

    service_unit=$(realpath "${service_dir}/.config/systemd/user/compose.service")
    if [[ -f ${service_unit} ]]; then
        log "✓️ service unit"
    else
        valid="no"
        log "✗ service unit ( ${service_unit} )"
    fi

    bus_socket="/run/user/$(id -u "${service}")/bus"
    if [[ -S "${bus_socket}" ]]; then
        log "✓️ bus socket"
    else
        valid="no"
        log "✗ bus socket ( ${bus_socket} )"
    fi

    docker_socket="/run/user/$(id -u "${service}")/docker.sock"
    if [[ -S "${docker_socket}" ]] && ping_docker_socket "${docker_socket}"; then
        log "✓️ docker socket"
    else
        valid="no"
        log "✗ docker socket ( ${docker_socket} )"
    fi

    if [[ ${valid} != "yes" ]]; then
        panic "Invalid environment ${service} in ${service_dir}"
    fi
)

runas() (
    log "Running as ${service}: $*"
    run_in_env "$*"
)

# ENTRYPOINT

basedir="/opt/docker"

# read global options
while [[ $# -gt 0 ]]; do
    case $1 in
    "-v" | "--verbose")
        shift
        set -o xtrace
        ;;
    "-h" | "--help")
        print_usage
        exit 0
        ;;
    *) break ;;
    esac
done

# root only
assert_root

# assert basedir initialized
basedir_init

# read command arg
command=$1
shift
if [[ -z ${command} ]]; then
    print_usage
    exit 64
fi

# commands without service arg
case ${command} in
"list")
    list "$@"
    exit
    ;;
esac

# read service arg
service=$(sanitize "$1")
shift
if [[ -z ${service} ]]; then
    print_usage
    exit 64
fi
service_dir="${basedir}/${service}"

# commands with service arg
case ${command} in
"create")
    create "$@"
    exit
    ;;
"delete")
    delete "$@"
    exit
    ;;
"start")
    start "$@"
    exit
    ;;
"stop")
    stop "$@"
    exit
    ;;
"restart")
    restart "$@"
    exit
    ;;
"logs")
    logs "$@"
    exit
    ;;
"status")
    status "$@"
    exit
    ;;
"validate")
    validate "$@"
    exit
    ;;
"runas")
    runas "$@"
    exit
    ;;
esac

# no command executed
print_usage
exit 64