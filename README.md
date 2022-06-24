# gocar-api

## Start server

Requires `docker` and `docker-compose`.

```sh
docker compose up
```

## Documentation

### Create Car

```sh
curl --request POST \
  --url http://localhost:8080/cars \
  --header 'Content-Type: application/json' \
  --data '{
	"first_name": "Dulcia",
	"last_name": "Gaskamp",
	"email": "dgaskamp7@w3.org",
	"gender": "Female",
	"address": "02060 Gerald Junction, San Jose",
	"car_manufactur": "Mercedes-Benz",
	"car_model": "SL-Class",
	"car_model_year": 2007
}'
```

### Search Cars

```sh
curl --request GET \
  --url 'http://localhost:8080/cars?first_name=Dulcia&last_name=Gaskamp&city=San%20Jose&manufacturer=Mercedes-Benz'
```

### Export Cars

```sh
curl --request GET \
  --url 'http://localhost:8080/cars/export?=&first_name=Dulcia&last_name=Gaskamp&city=San%20Jose&manufacturer=Mercedes-Benz'
```
