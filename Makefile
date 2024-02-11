ifneq ("$(wildcard .env)", "")
	include .env
endif
export

.PHONY: *

run_influx:
	docker run -d -p ${INFLUX_PORT}:8086 -v influxVolume:/var/lib/influxdb2 --name influx influxdb:latest

run_sentinel:
	docker build -t sentinel . && docker run -d -p ${APP_PORT}:${APP_PORT} --env-file .env --name sentinel sentinel
