#!/usr/bin/env bash

function do_test(){
  go test .
}

function do_commit(){
  git commit -a
}

function do_revert(){
  git checkout HEAD main.go
}

while ((1)); do
  inotifywait -r -e modify . >&2
  do_test && do_commit || do_revert
done
