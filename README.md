# exmaple: Golang RESTful API 
[![Build Status](https://travis-ci.org/codehand/echo-restful-crud-api-example.svg?branch=master)](https://travis-ci.org/codehand/echo-restful-crud-api-example)

apig is an RESTful API server generator.

## Contents

* [Contents](#contents)
* [How to build and install](#how-to-build-and-install)
* [How to use](#how-to-use)
  + [1. Generate boilerplate](#1-generate-boilerplate)
  + [2. Write model code](#2-write-model-code)
  + [3. Generate controllers, tests, documents etc. based on models.](#3-generate-controllers-tests-documents-etc-based-on-models)
  + [4. Build and run server](#4-build-and-run-server)
* [Usage](#usage)
  + [`new` command](#new-command)
  + [`gen` command](#gen-command)
  + [API Document](#api-document)
* [API server specification](#api-server-specification)
  + [Endpoints](#endpoints)
  + [Available URL parameters](#available-url-parameters)
  + [Data Type](#data-type)
  + [Pagination](#pagination)
  + [Versioning](#versioning)
* [License](#license)

## How to build and install

Go 1.11 or higher is required.

After installing required version of Go, you can build and install `apig` by

```bash
$ go get -d -u github.com/codehand/echo-restful-crud-api-example
$ cd $GOPATH/src/github.com/codehand/echo-restful-crud-api-example
$ make
$ make install
```

`make` generates binary into `bin/src`.
`make install` put it to `$GOPATH/bin`.

## How to use

First, creating by  command.

```
$ go run main.go
```
### Endpoints

Each resource has 5 RESTful API endpoints.
Resource name is written in the plural form.

|Endpoint|Description|Example (User resource)|
|--------|-----------|-------|
|`GET /<resources>`|List items|`GET /products` List products|
|`POST /<resources>`|Create new item|`POST /products` Create new product|
|`GET /<resources>/{id}`|Retrieve the item|`GET /products/1` Get the product which ID is 1|
|`PUT /<resources>/{id}`|Update the item|`PUT /products/1` Update the product which ID is 1|
|`DELETE /<resources>/{id}`|Delete the item|`DELETE /products/1` Delete the product which ID is 1|

### Data Type

#### Request

API server accepts the form of `JSON` or `Form`.

`application/json`

```
curl -X POST http://localhost:9090/resources \
     -H "Content-type: application/json" \
     -d '{"field":"value"}'
```


#### Response

Response data type is always `application/json`.

### Pagination (not support)

API server supports 2 pagination types.

#### Offset-based pagination

Retrieve items by specifying page number and the number of items per page.

For example:

```
http://103.20.150.235/api/v1/products?limit=5&page=2
```

```
+---------+---------+---------+---------+---------+---------+---------+
| ID: 5   | ID: 6   | ID: 7   | ID: 8   | ID: 9   | ID: 10  | ID: 11  |
+---------+---------+---------+---------+---------+---------+---------+
          |                                                 |
 Page 1 ->|<-------------------- Page 2 ------------------->|<- Page 3
```

Response header includes `Link` header.

```
Link:   <http://103.20.150.235/api/v1/products?limit=5&page=3>; rel="next",
        <http://103.20.150.235/api/v1/products?limit=5&page=1>; rel="prev"
```

#### ID/Time-based pagination

Retrieve items by specifying range from a certain point.

For example:

```
http://103.20.150.235/api/v1/products?limit=5&last_id=100&order=desc
```

```
+---------+---------+---------+---------+---------+---------+---------+
| ID: 94  | ID: 95  | ID: 96  | ID: 97  | ID: 98  | ID: 99  | ID: 100 |
+---------+---------+---------+---------+---------+---------+---------+
          |               5 items (ID < 100)                |
          |<------------------------------------------------|
```

Response header includes `Link` header.

```
Link:   <http://103.20.150.235/api/v1/products?limit=5&last_id=95&order=desc>; rel="next"
```

## License
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)