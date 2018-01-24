#!/usr/bin/env bash

# We're disabling cgo which gives us a static binary.
# We're also setting the OS to Linux (in case someone builds this on a Mac or Windows).
# The -a flag means to rebuild all the packages we are using, which means all imports will be rebuilt with cgo disabled.
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o stemcstudio-search .
eval $(minikube docker-env)
docker build -t stemcstudio-search:v1 -f Dockerfile .