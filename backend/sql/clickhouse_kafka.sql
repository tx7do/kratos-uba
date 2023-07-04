CREATE OR REPLACE TABLE acceptance_status_queue
(
    part_date      DateTime DEFAULT now(),
    data_name      String,
    error_reason   String,
    error_handling String,
    report_type    String,
    report_data    String,
    part_event     String,
    status         Int32
)
    ENGINE = Kafka
    SETTINGS kafka_broker_list = 'localhost:9092',
            kafka_topic_list = 'logger.report.event',
            kafka_group_name = 'ck_saver',
            kafka_format = 'JSONEachRow';

CREATE MATERIALIZED VIEW acceptance_status_mv TO acceptance_status AS
SELECT part_date,
       data_name,
       error_reason,
       error_handling,
       report_type,
       report_data,
       part_event,
       status
FROM acceptance_status_queue;

CREATE OR REPLACE TABLE realtime_warehousing_kafka
(
    createTime DateTime DEFAULT now(),
    eventName  String,
    reportData String
)
    ENGINE = Kafka
    SETTINGS kafka_broker_list = 'localhost:9092',
            kafka_topic_list = 'logger.report.event',
            kafka_group_name = 'ck-saver',
            kafka_format = 'JSONEachRow',
            kafka_skip_broken_messages = 10;

CREATE MATERIALIZED VIEW realtime_warehousing_mv TO realtime_warehousing AS
SELECT toDateTime(createTime),
       eventName,
       reportData
FROM realtime_warehousing_kafka;