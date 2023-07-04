-- acceptance_status kafka表
CREATE OR REPLACE TABLE acceptance_status_kafka
(
    part_date      DateTime('Asia/Shanghai'),
    data_name      String,
    error_reason   String,
    error_handling String,
    report_type    String,
    report_data    String,
    part_event     String,
    status         Int32
)
    ENGINE = Kafka
    SETTINGS kafka_broker_list = 'host.docker.internal:9092',
            kafka_topic_list = 'logger.report.event',
            kafka_group_name = 'ck-saver',
            kafka_format = 'JSONEachRow',
            kafka_skip_broken_messages = 1;

-- acceptance_status kafka表 物化视图
CREATE MATERIALIZED VIEW acceptance_status_mv TO acceptance_status AS
SELECT part_date,
       data_name,
       error_reason,
       error_handling,
       report_type,
       report_data,
       part_event,
       status
FROM acceptance_status_kafka;

-- realtime_warehousing kafka表
CREATE OR REPLACE TABLE realtime_warehousing_kafka
(
    createTime DateTime('Asia/Shanghai'),
    eventName  String,
    reportData String
)
    ENGINE = Kafka
    SETTINGS kafka_broker_list = 'host.docker.internal:9092',
            kafka_topic_list = 'logger.report.event',
            kafka_group_name = 'ck-saver',
            kafka_format = 'JSONEachRow',
            kafka_skip_broken_messages = 1;

-- realtime_warehousing kafka表 物化视图
CREATE MATERIALIZED VIEW realtime_warehousing_mv TO realtime_warehousing AS
SELECT createTime as create_time,
       eventName as event_name,
       reportData as report_data
FROM realtime_warehousing_kafka;
