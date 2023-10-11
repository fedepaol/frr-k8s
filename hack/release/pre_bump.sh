#!/bin/bash
set -e

if [ -z "$FRRK8S_VERSION" ]; then
    echo "must set the FRRK8S_VERSION environment variable"
    exit -1
fi

gitstatus=$(git status --porcelain)
if [ -n "$gitstatus" ]; then
	echo "uncommitted changes"
	echo $gitstatus
	exit 1
fi

pushd hack/semver
semver="$(go run . $FRRK8S_VERSION)"
popd
semver_split=($semver)
major=${semver_split[0]}
version=${semver_split[1]}
minor=${semver_split[2]}

git checkout main

if ! grep -q "## Release v$FRRK8S_VERSION" RELEASE_NOTES.md; then
  echo "Version $FRRK8S_VERSION missing from release notes"
  exit 1
fi

if [ $minor = "0" ]; then # patch release
	git checkout -b v$major.$version
else
	git checkout v$major.$version
fi

git checkout main -- RELEASE_NOTES.md
