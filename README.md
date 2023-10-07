# Consumer-and-producer-kafka-with-Golang



* install kafka in local
  * install kakfa : 
  * $ tar -xzf kafka_version.tgz
  -$ cd kafka_version

  /$ bin/zookeeper-server-start.sh config/zookeeper.properties
  /$ bin/kafka-server-start.sh config/server.properties
  
bin/kafka-server-start.sh --bootstrap-server=localhost:9092 --create --topic coordinates --replication-factor 1 --partitions 1


$ mkdir ~/projects/kafka-golang
$ cd ~/projects/kafka-golang
$ go mod init kafka-golang
$ go get github.com/confluentinc/confluent-kafka-go/kafka
for Create kafka topic : 


