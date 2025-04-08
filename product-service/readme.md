# Product Service Microservice

A Go-based microservice for managing product catalog, inventory, and search functionality for a phone accessories web application.

## Features

- RESTful API for product management
- Product catalog with filtering options
- Category management
- Inventory tracking
- Search functionality
- Swagger documentation
- Docker support

## Tech Stack

- Go (Golang)
- Gin web framework
- GORM ORM with PostgreSQL
- Swagger for API documentation
- Docker and Docker Compose

## Project Structure

```
product-service/
├── cmd/
│   └── main.go                 # Application entry point
├── docs/                       # Swagger documentation
├── internal/
│   ├── api/                    # API handlers
│   ├── config/                 # Configuration
│   ├── models/                 # Data models
│   ├── repository/             # Data access layer
│   └── service/                # Business logic
├── .env                        # Environment variables
├── .gitignore                  # Git ignore file
├── Dockerfile                  # Docker build instructions
├── docker-compose.yml          # Docker Compose setup
├── go.mod                      # Go modules
├── go.sum                      # Go modules checksums
└── README.md                   # Project documentation
```

## Prerequisites

- Go 1.19 or higher
- Docker and Docker Compose
- PostgreSQL (if running locally)

## Getting Started

### Running with Docker Compose (Recommended)

The easiest way to run the application is with Docker Compose:

```bash
# Clone the repository
git clone <repository-url>
cd phone-accessories

# Start the services
docker-compose up -d

# Check if the services are running
docker-compose ps
```

The API will be available at http://localhost:8080, and Swagger documentation at http://localhost:8080/swagger/index.html.

### Running Locally

To run the application locally:

```bash
# Clone the repository
git clone <repository-url>
cd phone-accessories

# Install dependencies
go mod download

# Generate Swagger documentation
go install github.com/swaggo/swag/cmd/swag@latest
swag init

# Make sure PostgreSQL is running locally
# Update .env file with your PostgreSQL credentials if needed

# Run the application
go run main.go
```

## API Endpoints

The following API endpoints are available:

### Products

- `GET /api/v1/products` - List all products (with filtering)
- `GET /api/v1/products/{id}` - Get a product by ID
- `POST /api/v1/products` - Create a new product
- `PUT /api/v1/products/{id}` - Update a product
- `DELETE /api/v1/products/{id}` - Delete a product
- `PATCH /api/v1/products/{id}/stock` - Update product stock

### Categories

- `GET /api/v1/categories` - List all categories
- `GET /api/v1/categories/{id}` - Get a category by ID
- `POST /api/v1/categories` - Create a new category
- `PUT /api/v1/categories/{id}` - Update a category
- `DELETE /api/v1/categories/{id}` - Delete a category

### Search

- `GET /api/v1/search?q={query}` - Search products

### Other

- `GET /api/v1/health` - Health check endpoint
- `GET /swagger/*any` - Swagger documentation

## Configuration

Configuration is managed through environment variables or a .env file:

- `SERVER_PORT` - Port for the HTTP server (default: 8080)
- `DB_HOST` - PostgreSQL host (default: localhost)
- `DB_PORT` - PostgreSQL port (default: 5432)
- `DB_USER` - PostgreSQL username (default: postgres)
- `DB_PASSWORD` - PostgreSQL password (default: postgres)
- `DB_NAME` - PostgreSQL database name (default: product_service)
- `DEFAULT_PAGE_SIZE` - Default page size for pagination (default: 20)
- `MAX_PAGE_SIZE` - Maximum page size for pagination (default: 100)

## Testing the API

You can test the API using curl, Postman, or any other HTTP client. Here are some examples:

### List Products

```bash
curl -X GET "http://localhost:8080/api/v1/products?categoryId=1&minPrice=10&maxPrice=100&page=1&pageSize=10" -H "accept: application/json"
```

### Create a Product

```bash
curl -X POST "http://localhost:8080/api/v1/products" -H "accept: application/json" -H "Content-Type: application/json" -d '{
  "name": "Phone Case",
  "description": "Protective case for smartphones",
  "price": 19.99,
  "sku": "CASE-001",
  "stockLevel": 100,
  "categoryId": 1,
  "attributes": {
    "color": "black",
    "material": "silicone"
  }
}'
```

### Search Products

```bash
curl -X GET "http://localhost:8080/api/v1/search?q=case&page=1&pageSize=10" -H "accept: application/json"
```

## License

[MIT License](LICENSE)
