ifneq ("$(wildcard .env)", "")
	include .env
endif
export

.PHONY: *

run_influx:
	docker run -p ${INFLUX_PORT}:8086 -v influxVolume:/var/lib/influxdb2 influxdb:latest

run_sentinel:
	docker build -t sentinel . && docker run -d -p ${APP_PORT}:${APP_PORT} --env-file .env --name sentinel sentinel
