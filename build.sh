#!/usr/bin/env bash

# CD into the Go source
pushd "$(dirname -- "$0")"/go || exit

# Update Go
RELEASE_BRANCH="release-branch.go1.21"
git fetch origin "$RELEASE_BRANCH"
git checkout -f "$RELEASE_BRANCH"
git reset --hard origin/"$RELEASE_BRANCH"

# Patch Go
git apply --ignore-space-change --ignore-whitespace --3way ../patches/*.diff
git checkout --theirs .
cp -p -P -v -R ../overlays/* ./

# Build Go
pushd "src" || exit
./make.bash