#!/bin/sh

cd ../kafka_2.13-3.7.1/

bin/kafka-topics.sh --bootstrap-server localhost:9092 --create --topic OrdersReceived --config retention.ms=259200000
bin/kafka-topics.sh --bootstrap-server localhost:9092 --create --topic OrdersConfirmed --config retention.ms=259200000
bin/kafka-topics.sh --bootstrap-server localhost:9092 --create --topic OrdersPickedAndPacked --config retention.ms=259200000
bin/kafka-topics.sh --bootstrap-server localhost:9092 --create --topic Notifications --config retention.ms=259200000
bin/kafka-topics.sh --bootstrap-server localhost:9092 --create --topic Errors --config retention.ms=259200000
