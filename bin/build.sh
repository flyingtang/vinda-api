#!bin/sh

docker login --username=txg5214 registry.cn-hangzhou.aliyuncs.com

docker build -t vinda-article .

docker tag vinda-article:latest registry.cn-hangzhou.aliyuncs.com/vinda/vinda-article:1.0.0

docker push registry.cn-hangzhou.aliyuncs.com/vinda/vinda-article:1.0.0