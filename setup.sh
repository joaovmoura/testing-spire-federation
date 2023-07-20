#!/bin/bash

sudo docker run -d -p 32000:5000 --restart=always --name registry registry:2

cd app/server

sudo ./hack/build.sh

cd ../client

sudo ./hack/build.sh