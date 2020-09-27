# Mail Service

## Usage

You need to set environment variables.

```
export EMAIL_HOST="YOUR_SMTP_HOST"
export EMAIL_USERNAME="YOUR_MAIL_HERE"
export EMAIL_PASSWORD="YOUR_PASSWORD_HERE"
export DB_CONNECTION="YOUR_CONNECTION_STRING_HERE"
```
Then run application.

```
go run main.go
```

## Example Request
```
curl --location --request POST 'localhost:8080/v1/email' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"YOUR_NAME_HERE",
    "surname":"YOUR_SURNAME_HERE",
    "type":2,
    "email":"YOUR_EMAIL_HERE"
}'
```