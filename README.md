# Card-Service

## Build & Run

### Build Docker Container
```bash
make docker-build
```

### Start Docker Container
```bash
make docker-up
```

---

## Local Development

- Run the service locally:
```bash
make run
```

- The service will be available at:
  http://localhost:1500

---

## API Details

### GET /card – Get Card Data

**Request Header:**
```
Authorization: {authToken}
```

**Request Body:**
```json
{}
```

**Response:**
```json
{
  "cardId": "000018cfe1a211ef95a30242ac180002",
  "userId": "000018b0e1a211ef95a30242ac180002",
  "name": "My Debit Card",
  "issuer": "TestLab",
  "number": "6821 5668 7876 2379",
  "status": "Active",
  "color": "#00a1e2",
  "borderColor": "#ffffff"
}
```

---

### GET /health – Health Check

**Response:**
```json
{
  "status": "healthy"
}
```

---

## Unit Testing

### Run Unit Tests
```bash
make test-service
```

### Generate Mocks with Mockery

1. Install mockery:
```bash
make mock-install
```

2. Generate mock files:
```bash
mockery --all --dir=internal/core/port --output=internal/core/port/mocks --outpkg=mocks
```

---

See `Makefile` for other useful commands.
