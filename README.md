# Days Ago

A simple API to calculate the date N days ago, implemented in Go and containerized with Docker.

## Features

- Calculate the date N days ago via a REST API.
- Lightweight and efficient implementation using Go.
- Configurable server port via `.env` file.
- Dockerized for ease of deployment.

---

## Prerequisites

- Docker installed on your system.
- A `.env` file in the project root with the following content:
PORT=8089


## Getting Started

### 1. Clone the Repository
Clone the repository to your local machine.

```bash
git clone <repository-url>
cd days-calculator
docker build -t days-calculator .
docker run --rm --env-file .env -p 8089:8089 days-calculator
Server started at http://localhost:8089
```

## Using the API

- You can use the API to calculate the date N days ago by sending a GET request to the /api/calculate endpoint.


### Example Request
curl "http://localhost:8089/api/calculate?days=2"


### Example Response
{"date":"2024/12/21"}

## Using the CLI

In addition to the HTTP API, the tool can be run directly from the command line.
Build and run the binary with a `-days` flag or a positional argument:

```bash
go build -o days_calculator ./app/days_calculator.go
./days_calculator -days 5   # using flag
./days_calculator 3         # using positional argument
```

Both commands output the date N days prior to today.


