CREATE OR REPLACE TABLE acceptance_status
(
    id             Int64,
    data_name      String,
    report_type    String,
    report_data    String,
    status         Int32,
    error_reason   String,
    error_handling String,
    part_event     String,
    part_date      DateTime DEFAULT now()
)
    ENGINE = MergeTree()
    PARTITION BY (toYYYYMMDD(part_date))
        ORDER BY (toYYYYMMDD(part_date), id, data_name, error_reason, error_handling, report_type, status)
        TTL part_date + toIntervalMonth(3)
        SETTINGS index_granularity = 8192;

CREATE OR REPLACE TABLE realtime_warehousing
(
    id          Int64,
    event_name  String,
    report_data String,
    create_time DateTime DEFAULT now()
)
    ENGINE = MergeTree()
    PARTITION BY (toYYYYMMDD(create_time))
        ORDER BY (toYYYYMMDD(create_time), id, event_name)
        TTL create_time + toIntervalMonth(3)
        SETTINGS index_granularity = 8192;
