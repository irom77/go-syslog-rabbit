### Syslog to Raabbit server

go get -u github.com/irom77/go-syslog-rabbit/

$ $GOPATH/bin/go-syslog-rabbit -r="guest:guest@localhost:5672"
Connecting to amqp://guest:guest@localhost:5672
UDP server up and listening on port 12514


$GOPATH/bin/go-syslog-rabbit -s="127.0.0.1:12514" -r="guest:guest@localhost:5672" -p=true


syslog-generator -ip="127.0.0.1" -protocol="udp" -port="12514"