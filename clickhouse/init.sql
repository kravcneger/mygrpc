DROP DATABASE IF EXISTS mygrpc;
CREATE DATABASE mygrpc;
\c mygrpc

CREATE TABLE UserLog
(
    login String,
    email String,
    created_at DateTime
) ENGINE = Kafka SETTINGS kafka_broker_list = 'kafka:9092',
    kafka_topic_list = 'create_user',
    kafka_group_name = 'mygrpc',
    kafka_format = 'JSONEachRow'