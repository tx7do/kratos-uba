CREATE TABLE acceptance_status_queue
(
    id             Int64,
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
            kafka_group_name = 'logger_saver',
            kafka_format = 'JSONEachRow';

CREATE MATERIALIZED VIEW acceptance_status_view TO acceptance_status AS
SELECT id,
       part_date,
       data_name,
       error_reason,
       error_handling,
       report_type,
       report_data,
       part_event,
       status
FROM acceptance_status_queue;

CREATE TABLE realtime_warehousing_queue
(
    id          Int64,
    create_time DateTime DEFAULT now(),
    event_name  String,
    report_data String
)
    ENGINE = Kafka
    SETTINGS kafka_broker_list = 'localhost:9092',
            kafka_topic_list = 'bi.etl',
            kafka_group_name = 'etl_saver',
            kafka_format = 'JSONEachRow';

CREATE MATERIALIZED VIEW realtime_warehousing_view TO realtime_warehousing AS
SELECT id,
       create_time,
       event_name,
       report_data
FROM realtime_warehousing_queue;
