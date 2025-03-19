DROP INDEX IF EXISTS stamp_idx;
DROP INDEX IF EXISTS id_type_stamp_ids;

SELECT drop_hypertable('sensor_values', if_exists => TRUE);

DROP TABLE IF EXISTS sensor_values;
