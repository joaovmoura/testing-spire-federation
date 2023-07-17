#!/usr/bin/env bash

go build main.go
if [ $? -ne 0 ]; then
    echo "Ops"
    exit 1
fi

echo "Everything is ok!"