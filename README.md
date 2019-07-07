# exmaple: Golang RESTful API 
[![Build Status](https://travis-ci.org/codehand/echo-restful-crud-api-example.svg?branch=master)](https://travis-ci.org/codehand/echo-restful-crud-api-example)

## Contents

* [Contents](#contents)
* [How to build and install](#how-to-build-and-install)
* [How to use](#how-to-use)
* [API server specification](#api-server-specification)
  + [Endpoints](#endpoints)
  + [Data Type](#data-type)
  + [Pagination](#pagination)
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
### Endpoints v1

Each resource has 5 RESTful API endpoints.
Resource name is written in the plural form.

|Endpoint|Description|Example (User resource)|
|--------|-----------|-------|
|`GET /api/v1/<resources>`|List items|`GET /api/v1/products` List products|
|`POST /api/v1/<resources>`|Create new item|`POST /api/v1/products` Create new product|
|`GET /api/v1/<resources>/{id}`|Retrieve the item|`GET /api/v1/products/1` Get the product which ID is 1|
|`PUT /api/v1/<resources>/{id}`|Update the item|`PUT /api/v1/products/1` Update the product which ID is 1|
|`DELETE /api/v1/<resources>/{id}`|Delete the item|`DELETE /api/v1/products/1` Delete the product which ID is 1|

### Endpoints v2

Version 2 is difference version 1, it must to have token in header such as: 
Authorization: Bearer example-test

Unless, API return :

    
        {
            "message": "missing key in request header"
        }
    

Each resource has 5 RESTful API endpoints.
Resource name is written in the plural form.

|Endpoint|Description|Example (User resource)|
|--------|-----------|-------|
|`GET /api/v2/<resources>`|List items|`GET /api/v2/products` List products|
|`POST /api/v2/<resources>`|Create new item|`POST /api/v2/products` Create new product|
|`GET /api/v2/<resources>/{id}`|Retrieve the item|`GET /api/v2/products/1` Get the product which ID is 1|
|`PUT /api/v2/<resources>/{id}`|Update the item|`PUT /api/v2/products/1` Update the product which ID is 1|
|`DELETE /api/v2/<resources>/{id}`|Delete the item|`DELETE /api/v2/products/1` Delete the product which ID is 1|


### Data Type

#### Request

API server accepts the form of `JSON` or `Form`.

`application/json`

```
curl -X POST http://localhost:9090/api/v1/resources \
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