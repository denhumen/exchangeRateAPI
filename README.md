# Exchange Rate API

## Overview

This project implements a gRPC service and client that allows users to:
1. Get the current USD to UAH exchange rate.
2. Subscribe an email to receive daily updates about the exchange rate.
3. Unsubscribe an email from receiving updates.

## Features

- **Get Current Exchange Rate**: Fetches the current exchange rate from an external API.
- **Subscribe Email**: Allows users to subscribe to daily exchange rate updates via email.
- **Unsubscribe Email**: Allows users to unsubscribe from the daily email updates.
- **Daily Email Updates**: Sends daily emails to all subscribed users with the current exchange rate.

## API Specification

The API is described using the Swagger documentation and can be viewed using [Swagger Editor](https://editor.swagger.io/). The API follows the gRPC protocol.

## Project Structure

```
.
├── go.mod
├── go.sum
├── src
│ └── exchangeRateAPI
│ ├── client
│ │ ├── 1
│ │ │ └── client.go
│ │ └── 2
│ │ └── client.go
│ ├── db
│ │ └── db.go
│ ├── proto
│ │ ├── exchange_rate
│ │ │ ├── exchange_rate.pb.go
│ │ │ └── exchange_rate_grpc.pb.go
│ │ └── exchange_rate.proto
│ ├── server
│ │ └── server.go
│ └── services
│ ├── email_service.go
│ └── exchange_rate_service.go
└── subscribers.db
```

## Installation

### Prerequisites

- Go 1.22 or higher
- SQLite3
- gRPC and Protocol Buffers
- A valid API key from [Open Exchange Rates](https://openexchangerates.org/)
- SMTP server credentials for sending emails

### Clone the Repository

```sh
git clone https://github.com/your-username/exchangeRateAPI.git
cd exchangeRateAPI
```

### Enviroment variables

1. You need to create API key for openexchangerates server
2. 

```
OPENEXCHANGERATES_API_KEY=your_openexchangerates_api_key_here
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your_email@gmail.com
SMTP_PASSWORD=your_password_or_app_password
```

### Install Dependencies

```
go mod download
```

### Run the server

1. Navigate to the root directory
2. Run following command:
```
go run src/exchangeRateAPI/server/server.go
```

### Test the clients

1. Get exchange rate and subscribe email:

```
go run src/exchangeRateAPI/client/1/client.go
```

2. Get exchange rate and unsubscribe email:

```
go run src/exchangeRateAPI/client/2/client.go
```

### Implementation details

#### Database
- The service uses SQLite for storing subscribed emails
- The db.go file initializes the database and performs migrations

#### gRPC Services
- exchange_rate_service.go: Implements the gRPC service for getting the current exchange rate and managing email subscriptions.
- email_service.go: Contains the logic for sending daily emails to subscribed users.

#### Cron Job
- A cron job is set up using the cron package to send daily emails with the current exchange rate.

#### Running with Docker

1. Run in root directory following command:

```
docker-compose up --build
```

2. After this you can run clients localy without using Docker for them

### Conclusion
This project demonstrates how to create a gRPC service in Go that provides exchange rate information and email notifications. The service uses SQLite for data storage and includes features such as database migration, email sending, and daily notifications via a cron job. The provided Docker configuration allows easy containerization of the application.

