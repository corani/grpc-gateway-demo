#!/bin/bash
# © 2021 SAP SE or an SAP affiliate company. All rights reserved.

prog=$0

function print_usage {
  echo "Usage: $prog [options...]"
  echo
  echo "Available options:"
  echo "  -h               prints this help"
  echo "  -b               build project"
  echo "  -t               test project"
  echo "  -r               run server"
  echo "  -i               install tools"
  echo "  -g               run generators"
}

function go_install_tools {
    echo "[CMD] tools/install.sh"
    tools/install.sh
}

function go_generate {
  root=$(pwd)

  export PATH="${root}/tools/bin":$PATH

  echo "[CMD] buf generate"
  buf generate

  cd "${root}" || return
}

function go_build {
    mkdir -p bin 
    echo "[CMD] go build -o bin/server"
    go build -trimpath \
        -o bin/server cmd/server/*.go
    echo "[CMD] go build -o bin/server.debug"
    go build -gcflags='all=-N -l' \
        -o bin/server.debug cmd/server/*.go
    echo "[CMD] go build -o bin/client"
    go build -trimpath \
        -o bin/client cmd/client/*.go
    echo "[CMD] go build -o bin/client.debug"
    go build -gcflags='all=-N -l' \
        -o bin/client.debug cmd/client/*.go
}

function go_run {
    echo "[CMD] go run"
    go run cmd/server/*.go
}

function go_test {
    mkdir -p gen
    root=$(pwd)

    # Collect a map containing all packages that need to be tested
    declare -A packages
    packages["internal"]="${root}/internal"

    # Test packages one-by-one, collecting the results in individual files
    for name in "${!packages[@]}"; do
        path="${packages[$name]}"
        echo "[CMD] go test ${name}"

        cd "$path" || return
        go test -mod=vendor -v -timeout 1m \
            -cover -coverpkg=./... -coverprofile="${root}/gen/cover_${name}.out" \
            ./... | tee "${root}/gen/test_${name}.out"
    done
    cd "$root" || return
}

if [ "$#" -eq "0" ]; then
    print_usage
    exit 0
fi

while [ "$#" -gt "0" ]; do
  arg=$1
  shift

  case $arg in
    -h)
        print_usage
        exit 0
        ;;
    -b)
        go_build
        ;;
    -i)
        go_install_tools
        ;;
    -g)
        go_generate
        ;;
    -r)
        go_run
        ;;
    -t)
        go_build
        go_test
        ;;
    *)
        echo "error: unrecognized argument '$arg'"
        print_usage
        exit 1
        ;;
  esac

done

# © 2021 SAP SE or an SAP affiliate company. All rights reserved.
