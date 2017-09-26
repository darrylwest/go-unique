#!/bin/sh
# dpw@alameda.local
# 2017.09.26
#

image=darrylwest/unique-tcp:latest

docker build -t $image .
echo "docker push $image"
