#!/bin/bash
set -x

# NOTE: Tested only on macos
NEW_SUB_COMMAND=${1}
echo "Generating ${NEW_SUB_COMMAND}'s scaffolding and moving it to a project specific directory..."
cobra add ${NEW_SUB_COMMAND}
mv cmd/${NEW_SUB_COMMAND}.go ./cmd/gbsctl/${NEW_SUB_COMMAND}.go
sed -ie 's/package cmd/package main/' ./cmd/gbsctl/${NEW_SUB_COMMAND}.go
rm  ./cmd/gbsctl/${NEW_SUB_COMMAND}.goe
go install ./cmd/gbsctl
gbsctl
