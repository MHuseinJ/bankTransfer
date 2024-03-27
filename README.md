# Bank Transfer API
## requirement
- go
- Docker and docker compose (for postgres)
  
## structure code
consist of 2 apps
- api: main app
```
api
|---module
|-----|-----transactions
|-----|-----------|------entity
|-----|-----------|------handler
|-----|-----------|------repository
|-----|-----------|------service
|-----|-----accounts
|-----|-----------|------entity
|-----|-----------|------handler
|-----|-----------|------repository
|-----|-----------|------service
|---config
|-----|-----DBConfig
go.mod
main.go
.env
```
- bank: mock bank app
```
bank
|---module
|-----|-----transactions
|-----|-----------|------entity
|-----|-----------|------handler
|-----|-----------|------service
|-----|-----accounts
|-----|-----------|------entity
|-----|-----------|------handler
|-----|-----------|------repository
|-----|-----------|------service
|---config
|-----|-----DBConfig
go.mod
main.go
.env
```
## Flow of the apps
whe have 3 API on each apps (mock and api)

1. api
   1. validate
      > handle on user module, this api will call the validate on mock service, and mock service will check whether the account exist and name is simmilar with the request
    2. disbursment
       > handle on transaction module, is to create transfer request (disbursement) and the the transaction will put on hold until the bank setled or canceled the transaction via callback
      3. disbursement callback
         > handle on transaction module, to receive status from bank service (mock) and store the status into db


2. bank (mock)
   1. validate
      > handle on user module, the api will receive name and account number, will check account number on db and validate with name
    2. transfer
       > handle on transaction module, the api will receive transaction object and give HOLD status to the transaction
      3. settlement
         > handle on transaction module, the api will receive transaction object with the updated status and call api service callback to update the status of the transaction
## Postman collection
```
{
	"info": {
		"_postman_id": "de990c16-eaff-4a9d-9898-62c6fb1910da",
		"name": "Brick",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "457095"
	},
	"item": [
		{
			"name": "Bank API",
			"item": [
				{
					"name": "Settlement",
					"request": {
						"method": "POST",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\n    \"originAccount\" : \"1234567890\",\n    \"destinationAccount\" : \"1234567891\",\n    \"referenceID\": \"220074\",\n    \"amount\" : 10000,\n    \"status\" : true\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8000/disbursement-callback",
							"host": [
								"localhost"
							],
							"port": "8000",
							"path": [
								"disbursement-callback"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "API Transfer",
			"item": [
				{
					"name": "Validate Account",
					"request": {
						"method": "GET",
						"header": []
					},
					"response": []
				},
				{
					"name": "Make Transfer",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n    \"originAccount\" : \"1234567890\",\n    \"destinationAccount\" : \"1234567891\",\n    \"referenceID\": \"220075\",\n    \"amount\" : 10000\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8001/disbursement",
							"host": [
								"localhost"
							],
							"port": "8001",
							"path": [
								"disbursement"
							]
						}
					},
					"response": []
				}
			]
		}
	]
}
```

