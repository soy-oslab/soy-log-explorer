# soy-log-explorer

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/f159dbff0a6b4d92880708e7a5eb9166)](https://app.codacy.com/gh/soyoslab/soy_log_explorer?utm_source=github.com&utm_medium=referral&utm_content=soyoslab/soy_log_explorer&utm_campaign=Badge_Grade_Settings)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/718c099fb64f43d4818c3749dacbffef)](https://www.codacy.com/gh/soyoslab/soy_log_explorer/dashboard?utm_source=github.com&utm_medium=referral&utm_content=soyoslab/soy_log_explorer&utm_campaign=Badge_Coverage)
![build](https://github.com/soyoslab/soy_log_explorer/actions/workflows/linux-build-test.yml/badge.svg)
![dockerize](https://github.com/soyoslab/soy_log_explorer/actions/workflows/dockerize.yml/badge.svg)
![coverage](https://github.com/soyoslab/soy_log_explorer/actions/workflows/coverage.yml/badge.svg)

## Introduction

This project sends the messages got from soy_log_collector to Elasticsearch.
The internal process is below.
1. Classifies messages received from soy_log_collector according to hot/cold and pushes them to the corresponding queue.
2. Background daemon pops the message from the queue and unzips it.
3. Send the unzipped message to Elasticsearch using REST API.

## Installation

```bash
$ git clone https://github.com/soyoslab/soy_log_explorer.git
$ cd soy_log_explorer
```

## Usage

Set enviroment variables:
```bash
$ export EXPLORER_HOST = 0.0.0.0   # Server IP Address
$ export EXPLORER_PORT = 8972      # Server Port
$ export ES_HOST = x.x.x.x         # Elasticsearch IP Address
$ export ES_PORT = 9200            # Elasticsearch Port
```

If you did not set the environment variables, soy_log_explorer uses default values.

Default environment variables:
```
EXPLORER_HOST = 0.0.0.0
EXPLORER_PORT = 8972
ES_HOST = localhost
ES_PORT = 9200
```

Run soy_log_explorer:
```bash
$ go run cmd/explorer/main.go
```
