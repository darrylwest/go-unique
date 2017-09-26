#!/bin/sh
# dpw@alameda.local
# 2017.09.26
#

image=darrylwest/unique-tcp:latest

docker run -d --name unique-tcp-3001 -p 3001:80 $image

