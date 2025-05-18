# Kart Order Food API

A RESTful API for food ordering system built with Go and Fiber.

## Features

- Product management (CRUD operations)
- Order creation and management
- PostgreSQL database
- Swagger API documentation

## Prerequisites

- Go 1.22 or higher
- PostgreSQL
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/hieu2304/order-food-be.git
cd order-food-be
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your database credentials
```

4. Run the application:
```bash
go run main.go
```

## API Documentation

Access the Swagger documentation at: http://localhost:3000/swagger/