CREATE TABLE IF NOT EXISTS sensors (
    id SERIAL NOT NULL PRIMARY KEY,
    sensor_id VARCHAR(100) NOT NULL DEFAULT '',
    first_seen TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_seen TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX sensor_name_idx ON sensors(sensor_id);
