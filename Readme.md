<br>

<h1 align="center">Logektor ⚡️</h1>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/lalabuy948/logektor"><img src="https://goreportcard.com/badge/github.com/lalabuy948/logektor"/></a>
  <a href="/go.mod"><img src="https://img.shields.io/github/go-mod/go-version/lalabuy948/logektor"/></a>
  <a href="https://docs.python.org/3/index.html"><img src="https://img.shields.io/badge/python-3.7.6-blue.svg"/></a>
  <a href="/LICENCE"><img src="https://img.shields.io/badge/licence-ccpl-green"/></a>
</p>

<p align="center">
  System with high availability to store events/logs.
</p>

<br><br>

## Stack

- Client: `go` + `fasthttp`
- Event bus: `kafka`
- Workers: `python`
- DB: `postgres` <- To simplify my life

## Architecture

Single app & SOA

![SA](github/EventTrackingSA.svg)

## Setup

```sh
docker-compose up
```

For development

Client
```
go mod install

go run client/*.go
```

Worker
```
python3 -m venv venv

source venv/bin/activate

pip3 install -r requirements.txt

python3 -m worker
```