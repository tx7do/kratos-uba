CREATE OR REPLACE TABLE acceptance_status
(
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
        ORDER BY (toYYYYMMDD(part_date), data_name, error_reason, error_handling, report_type, status)
        TTL part_date + toIntervalMonth(3)
        SETTINGS index_granularity = 8192;

CREATE OR REPLACE TABLE realtime_warehousing
(
    eventName  String,
    reportData String,
    createTime DateTime DEFAULT now()
)
    ENGINE = MergeTree()
    PARTITION BY (toYYYYMMDD(createTime))
        ORDER BY (toYYYYMMDD(createTime), eventName)
        TTL createTime + toIntervalMonth(3)
        SETTINGS index_granularity = 8192;
