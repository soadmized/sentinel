CREATE EXTENSION IF NOT EXISTS timescaledb;

CREATE TABLE sensor_values (
    id SERIAL,
    stamp TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    sensor_id INT NOT NULL DEFAULT 0,
    value DOUBLE PRECISION NOT NULL DEFAULT 0,
    value_type VARCHAR(50) NOT NULL DEFAULT '',

    CONSTRAINT valid_value_type CHECK (value_type IN ('temperature', 'light', 'motion')),
    PRIMARY KEY (id, stamp)
);

SELECT create_hypertable('sensor_values', 'stamp', chunk_time_interval => INTERVAL '1 month');

CREATE INDEX stamp_idx ON sensor_values (stamp DESC);
CREATE INDEX id_type_stamp_ids ON sensor_values (sensor_id, value_type, stamp);
