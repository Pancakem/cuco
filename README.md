# Cuco: A Currency Converter


## Introduction
`Cuco` is a web API that offers currency conversion between three
currencies. The

`Cuco` is implemented using the `Golang` standard library alone. I believe
it shows a simple glimpse of the power that the Go ecosystem has in store.


## Install

This project uses [Go](https://golang.org/). Check it out if you don't have it locally installed.

Clone the project

```sh
$ make build
```

## Usage

Run `Cuco`
```sh
$ ./cuco
```

Spawn another terminal in order to use it using curl.

### Example
Convert from one currency (`NGN`) to another (`GHS`)
```sh
$ curl -X GET "http://localhost:5000/convert?from=ngn&to=ghs&amount=50"
```
The response will look something like this:
```
{"error":"None","data":{"currency_code":"GHS","currency_value":0.72}}
```

To get the conversion rates from a currency to the rest of the
supported currencies
```sh
$ curl -X GET curl -X GET "http://localhost:5000/conversion-table?currency=ghs"
```
The response will look something like this:
```
{"error":"None","data":{"base_currency":"ghs","rates":{"ksh":18.16860465116279,"ngn":69.11337209302326}}}
```
