#!/usr/bin/env bash

go build main.go
if [ $? -ne 0 ]; then
    echo "Ops"
    exit 1
fi

docker build . -t localhost:32000/greeter-server:demo

if [ $? -ne 0 ]; then
    echo "Ops"
    exit 1
fi

docker push localhost:32000/greeter-server:demo

if [ $? -ne 0 ]; then
    echo "Ops"
    exit 1
fi

echo "Everything is ok!"
