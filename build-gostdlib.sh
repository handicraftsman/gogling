#!/usr/bin/env bash

set -x

source ./packages.sh

mkdir -p $PWD/lib

IFS=':' read -ra PKG <<< "$GO2LUAPKGS"
for p in "${PKG[@]}"; do
  go build -buildmode=plugin -o "${PWD}/lib/go2lua_${p}.so" "github.com/handicraftsman/gogling/gostdlib/${p}"
done;