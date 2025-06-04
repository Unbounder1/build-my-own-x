#!/bin/bash

docker build -t html-multi ./environment

docker run -d -p 8081:80 -e FLAVOR=1 --name box1 html-multi
docker run -d -p 8082:80 -e FLAVOR=2 --name box2 html-multi
docker run -d -p 8083:80 -e FLAVOR=3 --name box3 html-multi
docker run -d -p 8084:80 -e FLAVOR=4 --name box4 html-multi