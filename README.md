To Start the GoLang APP

## open the project Directory and run
go run main.go
## create customer
curl --location --request POST 'localhost:8080/customers' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Mohamed",
    "phone": "(212) 600798925",
    "balance":1000
}'
## getCustomer
curl --location --request GET 'localhost:8080/customers/customer_id_goes_here'

## Create Transfer
curl --location --request POST 'localhost:8080/transfers' \
--header 'Content-Type: application/json' \
--data-raw '{
    "origin_id": "customer_id_goes_here",
    "destination_id": "customer_id_goes_here",
    "amount": 10
}'

## get transfer
curl --location --request GET 'localhost:8080/transfers/transfer_id_goes_here'