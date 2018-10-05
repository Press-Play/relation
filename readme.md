# Quickstart

1. Install Golang: https://golang.org/dl/
2. Install Mongodb: https://docs.mongodb.com/manual/installation/
3. Clone repository: `git clone git@github.com:Press-Play/relation.git`
4. Start Mongodb instance:
    * OSX: `mongod`
    * Ubuntu: `sudo service mongod start && tail -f /var/log/mongodb/mongod.log`
5. Open Mongodb console: `mongo --host 127.0.0.1:27017`
6. Start server: `go run main.go`
