#!/bin/bash

# Exit immediately when a command fails
set -o errexit
# Error on unset variables
set -o nounset
# Only exit with zero if all commands of the pipeline exit successfully
set -o pipefail

# Source configuration
CUR_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
source "${CUR_DIR}/go_build_config.sh"

GOSEC_DIR_INSIDE_MODULES="${SRC_ROOT}/vendor/github.com/securego/gosec/v2"
GOSEC_DIR_INSIDE_GOPATH="${GOPATH}/src/github.com/securego/gosec/v2"

GOSEC_DIR=$( \
    realpath "${GOSEC_DIR:-$( \
        cd "${SRC_ROOT}"; \
        ls -d -1 "${GOSEC_DIR_INSIDE_MODULES}" 2>/dev/null || echo "${GOSEC_DIR_INSIDE_GOPATH}" \
    )}" \
)

echo "Gosec code with the following options:"
echo "      SRC_ROOT=${SRC_ROOT}"
echo "      GOSEC_DIR=${GOSEC_DIR}"

if [[ "${GOSEC_DIR}" == "${GOSEC_DIR_INSIDE_MODULES}" ]]; then
    echo "MODULES dir ${GOSEC_DIR} is used to run gosec from"
elif [[ "${GOSEC_DIR}" == "${GOSEC_DIR_INSIDE_GOPATH}" ]]; then
    echo "GOPATH dir ${GOSEC_DIR} is used to run gosec from"
else
    echo "CUSTOM dir ${GOSEC_DIR} is used to run gosec from"
fi

bash -c "cd ${GOSEC_DIR}; make build"
"${GOSEC_DIR}"/gosec -quiet "${CMD_ROOT}"/... "${PKG_ROOT}"/...
#  gosec -exclude-dir=rules -exclude-dir=cmd ./...
# gosec -tests ./...
# gosec can ignore generated go files with default generated code comment.
# // Code generated by some generator DO NOT EDIT.
# gosec -exclude-generated ./...
#it is possible to annotate the code with a #nosec comment.
#import "md5" // #nosec
#    /* #nosec */
#    if x > y {
#        h := md5.New() // this will also be ignored
#    }
