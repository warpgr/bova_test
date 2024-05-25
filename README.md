# Bova Test

This repository contains a service that provides pair ticker prices from Kraken.

## Getting Started

### Prerequisites

- Go (1.18 or higher)
- Docker (optional)
- Make (optional)

### Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/bova_test.git
cd bova_test
```

### Build and Run

Build the service
```bash
make build
```
or
```bash
go build -o price-service -ldflags="-s -w" cmd/main.go
```

### Testing and Linting

```bash
make lint
```
or
```bash
golangci-lint run -v
```

### Docker
```bash
make dc
```
or
```bash
docker-compose up
```

### Usage

The service provides two enpoints.
- GET /api/v1/ltp/list
    - HEADER: symbol="BTC/USDT,ETH/USDT"

- GET /api/v1/ltp



