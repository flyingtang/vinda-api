#!bin/sh

docker login --username=txg5214 registry-internal.cn-shenzhen.aliyuncs.com

docker build -t vinda-article .

docker tag vinda-article:latest registry-internal.cn-shenzhen.aliyuncs.com/vinda/vinda-article:1.0.0

docker push registry-internal.cn-shenzhen.aliyuncs.com/vinda/vinda-article:1.0.0