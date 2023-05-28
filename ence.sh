#!/usr/bin/env bash

function wire() {
  cd internal/wire || exit 1
  $GOPATH/bin/wire
}
