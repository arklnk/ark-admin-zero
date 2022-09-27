#!/usr/bin/env sh

dir=$1
api=$2

# shellcheck disable=SC2086
goctl api go -api "${dir}"/${api} -dir ${dir} -home /server/dev/goctl