#!/usr/bin/env bash

mkdir -p /data/timescaledb
mkdir -p /data/redis
mkdir -p /data/minio
mkdir -p /data/kafka

docker-compose up -d
