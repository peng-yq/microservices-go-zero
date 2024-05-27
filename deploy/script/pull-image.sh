#!/bin/bash

# define all the docker images we need
images=(
    "nginx:1.21.5"
    "lyumikael/gomodd:v1.20.3"
    "jaegertracing/all-in-one:1.42.0"
    "prom/prometheus:v2.28.1"
    "grafana/grafana:8.0.6"
    "docker.elastic.co/elasticsearch/elasticsearch:7.13.4"
    "docker.elastic.co/kibana/kibana:7.13.4"
    "kevinwan/go-stash:1.0"
    "elastic/filebeat:7.13.4"
    "wurstmeister/zookeeper"
    "wurstmeister/kafka"
    "hibiken/asynqmon:latest"
    "mysql/mysql-server:8.0.28"
    "redis:6.2.5"
)

for image in "${images[@]}"; do
    echo "Pulling $image..."
    docker pull $image &
    echo "$image pulled successfully!"
done

wait

echo "All images have been pulled successfully!"
