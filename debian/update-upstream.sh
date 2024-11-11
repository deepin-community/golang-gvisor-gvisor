#!/bin/bash

set -ex

MISSING_SOURCE="\
  pkg/sentry/platform/systrap/sysmsg/*.S \
  pkg/sentry/platform/systrap/sysmsg/*.c \
  pkg/sentry/platform/systrap/sysmsg/*.h \
  pkg/sentry/platform/systrap/sysmsg/*.sh \
  runsc/sandbox/bpf/ \
  tools/constraintutil \
  tools/go_fieldenum/ \
  tools/go_generics/ \
  tools/go_marshal/ \
  tools/go_stateify/ \
  tools/xdp/cmd/bpf/*.c \
  vdso/ \
"

GEN_FILES="
  ./pkg/sentry/loader/vdsodata/vdso_* \
  ./pkg/sentry/platform/systrap/sysmsg/*.bin \
  ./pkg/sentry/platform/systrap/sysmsg/sighandler_*.go \
  ./runsc/sandbox/bpf/*.o \
  ./tools/xdp/cmd/bpf/*.o \
"

upstream_tag="$1"

if [[ -z "$upstream_tag" ]]; then
  git remote add gvisor https://github.com/google/gvisor || true
  git fetch gvisor --tags

  # Find latest upstream tag
  upstream_tag=$(git tag -l|sort -V|grep '^release-'|tail -1)
fi

# Find the corresponding commit on `go` branch
go_commit=$(git rev-list --ancestry-path "$upstream_tag"..gvisor/go|tail -1)

# Work on upstream branch
git checkout upstream
git rm -r .
git checkout "$go_commit"  -- .
# Restore missing source
git checkout "$upstream_tag" -- $MISSING_SOURCE
# Remove broken tests
git rm -r --cached ./tools/*/tests ./tools/*/test
# Remove generated files
git rm --cache $GEN_FILES
git clean -x -d -f
tree=$(git write-tree)
commit=$(git commit-tree -p upstream^{commit} -p $go_commit -m "$upstream_tag" "$tree")
git update-ref refs/heads/upstream "$commit"

# Create upstream/x tag
git tag -a upstream/0.0_"${upstream_tag//release-/}" -m "$upstream_tag"

# Go back to debian/sid branch
git checkout debian/sid
gbp import-ref -u 0.0_"${upstream_tag//release-/}"
