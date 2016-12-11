#Clickhouse Exporter

Clickhouse exporter for prometheus.io, written in go.

## Using

    CLICKHOUSE_CONNECTION=http://localhost:8123 prometheus_clickhouse_exporter --listen 0.0.0.0:3000

## Installing

## Building

Requires [viper](github.com/spf13/viper) 
Requires [go-clickhouse](github.com/roistat/go-clickhouse) 

    go get github.com/implustech/prometheus_clickhouse_exporter
    ./prometheus_clickhouse_exporter -h

## Available groups of data

Exported from [system.events](https://clickhouse.yandex/reference_en.html#system.events) table 