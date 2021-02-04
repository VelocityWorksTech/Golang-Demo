<a href="https://www.velocityworks.io/home">Velocity Works Coding Demo</a>
# Golang-Demo

This Golang application consumes a JSON payload from https://www.data.gov/, populates a database and displays the database contents on a web page.

It is a simple http server is used to get donwload the JSON from https://labs.data.gov/dashboard/offices/qa . It features Rest endpoint to get the JSON response & save in a database. It is written using Echo Web Framework to make server high performant


## Contents

- [Golang-Demo](#golang-demo)
  - [Usage](#usage)
  - [Frameworks](#frameworks)
  - [Performance Metrics](#performance-metrics)
  - [Limitations](#limitations)


## Usage

To install Golang-Demo, you need to install [Go](https://golang.org/)(**version 1.12+ is required**) and set your Go workspace.

1. This project uses go modules and provides a make file. You should be able to simply install and start:

```sh
$ git clone https://github.com/anil-appface/golang-demo.git
$ cd golang-demo
$ make
$ ./go-datagov
```


## Frameworks

This project uses the below frameworks:

1. <a href="https://github.com/labstack/echo"><strong>Echo Framework: Simple & high performant server</strong></a>
2. <a href="https://github.com/go-resty/resty"><strong>Resty: Simple HTTP helper to get information</strong></a>
3. <a href="https://github.com/buger/jsonparser"><strong>JSON parser: To process the JSON data</strong></a>
4. <a href="https://github.com/jinzhu/gorm"><strong>Gorm: To populate the database</strong></a>

## Performance Metrics

Benchmarking for this application is not done.

<p align="justify"><i>"As this application uses Echo web framework, the default logs of echo server shows the Method type, uri, and Status code. Also it shows the logging of method name, line number and file."</i></p>


## Limitations

1. There is no authentication wrapper around the API's.
2. There is no field validation in api.
4. There is no much test cases written in the interest of time.

https://www.consumerfinance.gov/data.json

https://resources.data.gov/resources/dcat-us/