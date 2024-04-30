# http-load-tester

go-http-knocker is a simple load testing tool written in Go.
It allows you to send concurrent HTTP requests to a specified URL to test the performance and stability of your web application or API.

## Features

- Configurable number of concurrent requests
- Continuous request sending
- Real-time stats display

## Usage

1. Clone the repository

```
git clone https://github.com/zzucch/http-load-tester
```

2. Build the project:

```
cd http-load-tester
go build -o http-load-tester cmd/main/main.go
```

2. Create a `config.json` file with the structure similar to the `config-example.json` file
3. Run the binary:

```
./http-load-tester
```
