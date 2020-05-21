# go-cqrs-taxi

## Before run docker-compose up
Set this below value on host to fix elasticsearch error
```
sysctl -w vm.max_map_count=262144
```

### Login to get JWT token
```
curl -X POST  http://localhost:8084/auth/login  -H 'Cache-Control: no-cache'  -H 'Content-Type: application/json' -d '{"email": "user1@volho.com", "password":"user1"}'
```

### Insert new taxi via command service
```
curl -X POST  http://localhost:8081/command/InsertTaxi -H 'Authorization: Bearer <Insert JWT>' -H 'Cache-Control: no-cache'  -H 'Content-Type: application/json' -d '{"Taxi":{"id":"2","id_car":"myIdCar","id_center":"myIDCenter","id_driver":"myIDDriver","body":"","lat":48.137154,"lon":11.576124,"status":"AVAILABLE","created_at":"2020-02-16T03:38:26.434550066Z","updated_at":"2020-02-16T03:38:26.434550066Z","location":{"lat":48.138154,"lon":11.574124}}}'
```

### Update taxi location via command service
```
curl -X PUT  http://localhost:8081/command/UpdateTaxiLocation -H 'Authorization: Bearer <Insert JWT>' -H 'Cache-Control: no-cache'  -H 'Content-Type: application/json' -d '{"ID": "<Insert uuid>", "Lat": 48.13, "Lon": 11.57}'
```

### Get all taxies via query service
```
curl -X GET -H 'Authorization: Bearer <Insert JWT>' http://localhost:8082/query/SearchTaxies?query=48,11
```

### Search all from Elasticsearch directly
```
curl -X GET -H 'Authorization: Bearer <Insert JWT>' http://localhost:9200/taxies/_search?pretty=true&q=*:*
```

### For test websocket easily, https://github.com/websockets/wscat
```
npm install -g wscat

wscat -n -c wss://localhost/pusher -H Authorization:eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIyIiwiaWF0IjoxNTgyNDU2MzU2LCJleHAiOjE1ODMzMjAzNTZ9.9q7D4yZPjHpvwEixrCwjAHFBfpER_1QTZsWpevvBpIT7ks4KWGI9J8hF_em71s6dg2LhbDyGa-6UkqTzUmNaXw
```
