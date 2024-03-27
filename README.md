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
|-----|-----dbConfig
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
|-----|-----dbConfig
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
## How To Run
every step need to open different terminal,
step to run: 
1. run db
   > docker-compose up
2. run api
   > cd api && go build && ./api
3. run api
   > cd bank && go build && ./bank

after run, we can test it via postman, the [collection](postman.json)  already in the repo



