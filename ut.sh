#!/bin/bash

function cleanTestCache() {
    go clean -testcache;
}

function test_() {
    go test ./...;
}

case $1 in
    clean) cleanTestCache;
    ;;
esac

test_;
