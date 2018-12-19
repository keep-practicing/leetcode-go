#!/bin/bash

function cleanTestCache() {
    go clean -testcache;
}

function test() {
    go test ./...;
}

case $1 in
    clean) cleanTestCache;
    ;;
esac

test;
