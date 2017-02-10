docker@ubuntu-server:~/logstash$ pwd
/home/docker/logstash
docker@ubuntu-server:~/logstash$ ls
config  Dockerfile  exec  logstash.sh  opt
docker@ubuntu-server:~/logstash$ cat Dockerfile 
FROM logstash:5

# Add your logstash plugins setup here
# RUN logstash-plugin install logstash-filter-de_dot
# Example: RUN logstash-plugin install logstash-filter-json
RUN logstash-plugin install logstash-output-rabbitmq logstash-output-exec logstash-filter-json_encode

COPY ./config/* /etc/logstash/conf.d/
#COPY ./opt/* /opt/
#COPY ./exec/* /usr/local/bin/ 

CMD ["logstash", "-f", "/etc/logstash/conf.d/"]
docker@ubuntu-server:~/logstash$ cat logstash.sh 
#!/usr/bin/bash
docker stop logstash-prod
docker rm logstash-prod
docker build -t prod-logstash .
docker run -d --network rabbit -p 12514:12514/udp -v $PWD/opt:/opt -v /etc/localtime:/etc/localtime:ro --name logstash-prod prod-logstash 
sleep 10 
docker logs logstash-prod
docker@ubuntu-server:~/logstash$ cat config/logstash.conf 
input {
    udp {
            port => 12514
            #type => "threat"
            workers => 2 
            queue_size => 8000 # default 2000
    }
}
filter {
}
output {
    rabbitmq {
      exchange => "threat"
        exchange_type => "fanout"
        #key => "threat"
        host => "rabbit"
        #workers => 2         # if you have alot of messages raise this slowly
        #codec => "json"
        durable => false 
        persistent => false 
        port => 5672
        user => "guest"
        password => "guest"
        ssl => false         # over unsecure network do not use plain!
      }
    }