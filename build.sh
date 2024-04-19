#!/usr/bin/env bash

function usage() {
    cat <<USAGE
Usage: $0 [--patch patch_path] [--skip-update]

Options:
  -s, --skip-update:  skips Go update
  -p, --patch:        applies the given patch(es)
  -h, --help:         show this help
USAGE
    exit 1
}

SKIP_UPDATE=false
PATCHES=()

while [ "$1" != "" ]; do
    case $1 in
		-s | --skip-update)
			SKIP_UPDATE=true
			;;
		-p | --patch)
			shift
			PATCHES+=($1)
			;;
		-h | --help)
			usage
			;;
		*)
			usage
			exit 1
			;;
    esac
    shift
done

# CD into the Go source
pushd "$(dirname -- "$0")"/go || exit

RELEASE_BRANCH="release-branch.go1.22"

if [[ $SKIP_UPDATE == false ]]; then
	# Update Go
	git fetch origin "$RELEASE_BRANCH"
fi

git checkout -f "$RELEASE_BRANCH"
git reset --hard origin/"$RELEASE_BRANCH"

# Patch Go
if [ ${#PATCHES[@]} -eq 0 ]; then
	git apply --ignore-space-change --ignore-whitespace --3way ../patches/*.diff
else
	git apply --ignore-space-change --ignore-whitespace --3way ${PATCHES[@]}
fi

cp -p -P -v -R ../overlays/* ./

# Build Go
pushd "src" || exit
GODEBUG=installgoroot=all ./make.bash
