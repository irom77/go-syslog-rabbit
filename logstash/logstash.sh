#!/usr/bin/bash
docker stop logstash-prod
docker rm logstash-prod
docker build -t prod-logstash .
docker run -d --network rabbit -p 12514:12514/udp -v $PWD/opt:/opt -v /etc/localtime:/etc/localtime:ro --name logstash-prod prod-logstash
sleep 10
docker logs logstash-prod