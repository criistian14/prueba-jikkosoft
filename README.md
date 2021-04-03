# Prueba Jikkosoft

Prueba para la empresa Jikkosoft hecha en [GO](https://golang.org/) version 1.15 con arquitectura limpia

# End Points

## Sort Numbers

Ordenar los numeros mediante el algoritmo [QuickSort](https://www.youtube.com/watch?v=WaNLJf8xzC4), además se ubican de
último los numeros repetidos

### Request

```curl
curl -X POST -H "Content-Type: application/json" \
-d '{"numbers":[123,4,14,5,10,5,4]}' \
"http://localhost:3000/numbers/sort"
```

### Response

    HTTP/1.1 200 OK
    content-length: 61
    content-type: application/json

    {"sorted":[4,5,10,14,123,4,5],"unsorted":[123,4,14,5,10,5,4]}

# Tablas

## Inquiries (Solicitudes - PQR's)

Name   | Type     | extra
------ | -------- | ------
ID | bigint | pk
Title | varchar(250) |
Message | text |
State | enum('new', 'in_progress', 'closed') |
Category | enum('petition', 'complaint', 'claim') |

## Invoices (Facturas)

Name   | Type     | extra
------ | -------- | ------
ID | bigint | pk
PaymentDeadline | datetime |
TotalAmount | bigint |
Paid | tinyint |
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