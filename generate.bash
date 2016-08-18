#!/bin/bash

set -euo pipefail

:main() {
    local filename="food.go"
    :truncate "$filename"

    local packages=($(
        find api -mindepth 1 -maxdepth 1 -type d -printf '%f\n' \
            | grep -Pv '^\.' \
            | sort
    ))

    :import "$filename" "${packages[@]}"

    for package in "${packages[@]}"; do
        :compile "$filename" "$package"
    done

    :format "$filename"

    exit 0
}

:truncate() {
    local filename="$1"
    cat > "$filename" <<PACKAGE
package faces

PACKAGE
}

:compile() {
    local filename="$1"
    local package="$2"

    local object="$(sed 's/[^ ]\+/\L\u&/g' <<< "$package")"

    cat >> "$filename" <<NEW
func (context context) New$object() (*$package.$object, error) {
    face := new($package.$object)

    err := fabricate(face, context.Execution, "$package")
    if err != nil {
        return nil, err
    }

    return face, nil
}

func New$object() (*$package.$object, error) {
    return new(context).New$object()
}

NEW
}

:import() {
    local filename="$1"
    shift

    echo "import (" >> "$filename"
    while [ $# -ne 0 ]; do
        echo "	\"github.com/reconquest/faces/commands/$1\"" >> "$filename"
        shift
    done
    echo ")"$'\n' >> "$filename"
}

:format() {
    local filename="$1"

    gofmt -w -s "$filename"
}

:log() {
    echo "[$(date -Iseconds | cut -dT -f1)] $*" >&2
}

:main "$@"
