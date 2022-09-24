#!/usr/bin/env sh

dir=$1
api=$2

goctl api go -api ${dir}/${api} -dir ${dir} -home ./dev/goctl