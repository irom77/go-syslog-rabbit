### Syslog to Raabbit client

irek@ubuntu-server:~$ cd $GOPATH
irek@ubuntu-server:~/go$ go install ./src/github.com/irom77/go-syslog-rabbit/client
irek@ubuntu-server:~/go$ ls -l $GOPATH/bin
total 13300
-rwxrwxr-x 1 irek irek 6809680 Mar  5 17:19 client
-rwxrwxr-x 1 irek irek 6805432 Mar  5 16:32 go-syslog-rabbit

irek@ubuntu-server:~/go$ cd $GOPATH/bin
irek@ubuntu-server:~/go/bin$ ls
client  go-syslog-rabbit
irek@ubuntu-server:~/go/bin$ mv client go-rabbit-client 
irek@ubuntu-server:~/go/bin$ cd
irek@ubuntu-server:~$ go-rabbit-client -h
Copyright 2017 @IrekRomaniuk. All rights reserved.
Usage of go-rabbit-client:
  -q string
        Name of the queue (default "threat")
  -r string
        Rabbit server (default "guest:guest@192.168.3.51:5672")
  -v    Prints current version