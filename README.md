# Monitoring

## Install influxdb client

```
go mod init <package_uri>

go get github.com/influxdata/influxdb-client-go/v2
```

## run influxdb in local using docker

```
docker run --rm influxdb influxd config > influxdb.conf
```


```
docker run -d -p 8086:8086 -v /Users/baekido/Documents/05.SCALE_COE/sources/monitoring/data:/var/lib/influxdb --name local_influxdb influxdb
```

## create database in influxdb

```
curl -X POST http://localhost:8086/query --data-urlencode "q=CREATE DATABASE goweb"

{"results":[{"statement_id":0}]}
```

```
curl -X POST http://localhost:8086/query --data-urlencode "q=show databases"

{"results":[{"statement_id":0,"series":[{"name":"databases","columns":["name"],"values":[["_internal"],["mydb"]]}]}]}


curl -X POST http://localhost:8086/query --data-urlencode "q=SHOW MEASUREMENTS ON goweb"
```

## Build Docker 

```
docker build -t monitoring:v1.0 .
```

## 실행하기. 

```
docker run -e DB_HOST="http://192.168.0.14:8086" -e INTERVAL=10s  monitor:scale.1 --name scalemonitor
```

### with volume

```
docker run --name scalemonitor -v $(pwd)/target2.json:/bin/data/targets.json -e DB_HOST="http://192.168.0.14:8086" -e INTERVAL=2s monitor:scale.1
```