#!/usr/bin/env bash

GOOS=linux go build -o graph-server main.go
docker build -t graph-server .
