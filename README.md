# Prueba Jikkosoft

Prueba para la empresa Jikkosoft hecha en [GO](https://golang.org/) version 1.15 con arquitectura limpia y subida al
hosting gratuito [Heroku](https://www.heroku.com/).

# End Points

Los end points sé puedes probar en la siguiente página [ReqBin](https://reqbin.com/curl) copiando y pegando los curls.

## Sort Numbers

Ordenar los numeros mediante el algoritmo [QuickSort](https://www.youtube.com/watch?v=WaNLJf8xzC4), además se ubican de
último los numeros repetidos

### Request

```curl
curl -X POST -H "Content-Type: application/json" \
-d '{"numbers":[123,4,14,5,10,5,4]}' \
"https://pruebajikkosoft.herokuapp.com/numbers/sort"
```

### Response

    HTTP/1.1 200 OK
    content-length: 61
    content-type: application/json

    {"sorted":[4,5,10,14,123,4,5],"unsorted":[123,4,14,5,10,5,4]}

## Get User

Obtiene toda la información acerca del usuario, mediante él id de usuario pasado. Los resultados seran solicitados de
las tablas:

- Users
- Cities
- Invoices
- Inquiries

### Request

```curl
curl -X GET -H "Content-Type: application/json" \
"https://pruebajikkosoft.herokuapp.com/users/1"
```

### Response

    HTTP/1.1 200 OK
    content-length: 1874
    content-type: application/json

# Tablas

## Users (Usuarios)

Name   | Type     | extra
------ | -------- | ------
ID | bigint | pk
FirstName | varchar(250) |
LastName | varchar(250) |
Email | varchar(300) |
Address | varchar(250) |
Phone | varchar(11) |
CityID | bigint | fk

## Cities (Ciudades)

Name   | Type     | extra
------ | -------- | ------
ID | bigint | pk
Name | varchar(250) |
CountryID | bigint | fk

## Countries (Paises)

Name   | Type     | extra
------ | -------- | ------
ID | bigint | pk
Name | varchar(250) |

## Inquiries (Solicitudes - PQR's)

Name   | Type     | extra
------ | -------- | ------
ID | bigint | pk
Title | varchar(250) |
Message | text |
State | enum('new', 'in_progress', 'closed') |
Category | enum('petition', 'complaint', 'claim') |
UserID | bigint | fk

## Invoices (Facturas)

Name   | Type     | extra
------ | -------- | ------
ID | bigint | pk
PaymentDeadline | datetime |
TotalAmount | bigint |
Paid | tinyint |
UserID | bigint | fk
PublicServiceID | bigint | fk

## Public Services (Servicios Públicos)

Name   | Type     | extra
------ | -------- | ------
ID | bigint | pk
Company | varchar(250) |
Type | enum('water', 'electric') |
Email | varchar(250) |

# Librerias usadas

- [Fiber](https://docs.gofiber.io/)
- [GoDotEnv](https://github.com/joho/godotenv)
- [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
- [mapstructure](https://github.com/mitchellh/mapstructure)
- [Gorm](https://gorm.io/)
- [Testify](https://github.com/stretchr/testify)