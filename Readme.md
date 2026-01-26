# mkidp(mk identity provider)

A fun learning project to explore the Authentication and Authorization world. 

## Overview

This is a minimal OAuth2 Identity Provider (IdP) implementation built in Go. The goal is to understand how authentication and authorization flows work, starting with the basics and gradually adding more complex features.


## Roadmap

- [ ] Implement Branca tokens (encrypted, authenticated tokens) on CC type


## Getting Started

### Prerequisites

- Go 1.21+
- Air (for hot reload) - `go install github.com/cosmtrek/air@latest`

### Run

```bash
air
```

Or compile and run:

```bash
go run .
```

Server starts on `http://localhost:8080`

## Endpoints

### Get Clients

```bash
GET /api/v1/clients
```

### OAuth Token Endpoint

```bash
POST /oauth/token
```

Request (form-encoded):
```
client_id=1
client_secret=secret1
grant_type=client_credentials
```

Response:
```json
{
  "access_token": "random-token-here",
  "token_type": "Bearer",
  "expiry": "3600"
}
```

## Learning Goals

- Understand OAuth2 specification and security considerations
- Learn token generation and validation
- Explore different grant types
- Implement industry-standard token formats
