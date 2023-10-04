#!/usr/bin/env bash

go build main.go
if [ $? -ne 0 ]; then
    echo "Ops"
    exit 1
fi

docker build . -t vitinhocalvo/greeter-server

if [ $? -ne 0 ]; then
    echo "Ops"
    exit 1
fi

docker push vitinhocalvo/greeter-server

if [ $? -ne 0 ]; then
    echo "Ops"
    exit 1
fi

echo "Everything is ok!"
