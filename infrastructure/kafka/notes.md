docker exec -it kafka bash
./opt/kafka_2.11-0.10.1.0/bin/kafka-topics.sh --create --topic "timestamp.command" --zookeeper localhost:2181 --partitions 3 --replication-factor 1
./opt/kafka_2.11-0.10.1.0/bin/kafka-topics.sh --list --zookeeper localhost:2181
