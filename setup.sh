#!/bin/bash

if ! docker ps -a | grep -q "registry"; then
   sudo docker run -d -p 32000:5000 --restart=always --name registry registry:2
else
   echo "O container 'registry' já existe. Não será criado novamente."
fi

cd app/server

sudo ./hack/build.sh

cd ../client

sudo ./hack/build.sh
