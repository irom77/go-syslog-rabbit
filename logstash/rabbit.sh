#!/usr/bin/env bash
docker stop rabbit
docker rm rabbit 
docker run -d --network rabbit --hostname rabbit --name rabbit -p 5672:5672 -p 15672:15672 -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=n3w@yn rabbitmq:3-management