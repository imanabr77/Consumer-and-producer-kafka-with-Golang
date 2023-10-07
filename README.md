# Consumer-and-producer-kafka-with-Golang



* install kafka in local
  * install kakfa : 
  * $ tar -xzf kafka_version.tgz
  * $ cd kafka_version
  * $ bin/zookeeper-server-start.sh config/zookeeper.properties
  * $ bin/kafka-server-start.sh config/server.properties
 
  * ![image](https://github.com/imanabr77/Consumer-and-producer-kafka-with-Golang/assets/92488673/b4ff04b8-c1f7-4eb5-b792-54b90acfa17a)
  * ![image](https://github.com/imanabr77/Consumer-and-producer-kafka-with-Golang/assets/92488673/ec6d752a-2eea-419f-95e3-e57e820b65ed)


 
### for Create kafka topic :   
* bin/kafka-server-start.sh --bootstrap-server=localhost:9092 --create --topic coordinates --replication-factor 1 --partitions 1


* $ mkdir ~/projects/kafka-golang
* $ cd ~/projects/kafka-golang
* $ go mod init kafka-golang
* $ go get github.com/confluentinc/confluent-kafka-go/kafka

* $ go run producer.go 
![image](https://github.com/imanabr77/Consumer-and-producer-kafka-with-Golang/assets/92488673/d4fa6cf6-30a5-465f-b0a6-07786e8580a5)

* $ go run consumer.go
![image](https://github.com/imanabr77/Consumer-and-producer-kafka-with-Golang/assets/92488673/58ad7533-8c49-499a-9f8e-b4e7b056d692)
