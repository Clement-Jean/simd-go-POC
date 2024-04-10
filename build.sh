#!/usr/bin/env bash

function usage() {
    cat <<USAGE

    Usage: $0 [--skip-update]

    Options:
        -s, --skip-update:  skips Go update
USAGE
    exit 1
}

SKIP_UPDATE=false

for arg in "$@"; do
    case $arg in
    -s | --skip)
        SKIP_UPDATE=true
		shift
        ;;
    -h | --help)
        usage # run usage function on help
        ;;
    *)
        usage # run usage function if wrong argument provided
        ;;
    esac
done

# CD into the Go source
pushd "$(dirname -- "$0")"/go || exit

if [[ $SKIP_UPDATE == false ]]; then
	# Update Go
	RELEASE_BRANCH="release-branch.go1.22"
	git fetch origin "$RELEASE_BRANCH"
	git checkout -f "$RELEASE_BRANCH"
	git reset --hard origin/"$RELEASE_BRANCH"
else
	/bin/rm -rf src/simd
	/bin/rm -rf src/runtime/internal/simd
fi

# Patch Go
git apply --ignore-space-change --ignore-whitespace --3way ../patches/*.diff
cp -p -P -v -R ../overlays/* ./

# Build Go
pushd "src" || exit
GODEBUG=installgoroot=all ./make.bash
