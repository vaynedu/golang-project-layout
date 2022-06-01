#!/usr/bin/env bash

RUN_NAME="golang.project-layout"

mkdir -p output/bin output/conf
cp script/* output/
cp conf/* output/conf
chmod +x output/run.sh

go build -o output/bin/${RUN_NAME}