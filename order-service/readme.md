# Order Service Microservice

This microservice manages the order and cart functionality for a phone accessories e-commerce application. It handles cart management, order processing, and payment simulation.

## Features

- **Cart Management**: Add, update, remove items from cart
- **Order Processing**: Create orders, track status, view order history
- **Payment Simulation**: Process simulated payments, track payment status
- **API Documentation**: Swagger UI for easy API exploration
- **Docker Integration**: Containerized deployment for consistency across environments

## Tech Stack

- **Backend**: Node.js, Express.js
- **Database**: MongoDB
- **Documentation**: Swagger UI
- **Containerization**: Docker, Docker Compose

## Prerequisites

Before you begin, ensure you have the following installed:
- [Docker](https://www.docker.com/products/docker-desktop)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Node.js](https://nodejs.org/) (v14 or higher, for local development)

## Getting Started

### Running with Docker

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd order-service
   ```

2. Configure the environment variables (optional):
   - Review and modify the `.env` file if needed

3. Build and start the containers:
   ```bash
   docker-compose up --build
   ```

4. The service is now running at:
   - API: [http://localhost:3000](http://localhost:3000)
   - Swagger Documentation: [http://localhost:3000/api-docs](http://localhost:3000/api-docs)

### Running Locally (Without Docker)

1. Install dependencies:
   ```bash
   npm install
   ```

2. Start MongoDB locally or update the `.env` file to point to your MongoDB instance

3. Start the application:
   ```bash
   npm start
   ```

   Or for development with hot-reload:
   ```bash
   npm run dev
   ```

## API Endpoints

### Cart API
- `GET /api/cart/{userId}` - Get user's cart
- `POST /api/cart/{userId}/items` - Add item to cart
- `PUT /api/cart/{userId}/items/{productId}` - Update cart item quantity
- `DELETE /api/cart/{userId}/items/{productId}` - Remove item from cart
- `DELETE /api/cart/{userId}/clear` - Clear user's cart

### Order API
- `POST /api/orders` - Create a new order from cart
- `GET /api/orders/{orderId}` - Get order by ID
- `GET /api/orders/user/{userId}` - Get user's orders
- `PUT /api/orders/{orderId}/status` - Update order status

### Payment API
- `POST /api/payments/process` - Process payment for an order
- `GET /api/payments/{paymentId}` - Get payment by ID
- `GET /api/payments/order/{orderId}` - Get payments by order ID

## Example API Usage

### Add Item to Cart

```bash
curl -X POST http://localhost:3000/api/cart/user123/items \
  -H "Content-Type: application/json" \
  -d '{
    "productId": "prod001",
    "name": "Phone Case",
    "price": 19.99,
    "quantity": 1,
    "imageUrl": "/images/phone-case.jpg"
  }'
```

### Create an Order

```bash
curl -X POST http://localhost:3000/api/orders \
  -H "Content-Type: application/json" \
  -d '{
    "userId": "user123",
    "shippingAddress": {
      "fullName": "John Doe",
      "street": "123 Main St",
      "city": "Anytown",
      "state": "CA",
      "zipCode": "12345",
      "country": "USA"
    },
    "paymentMethod": "credit_card"
  }'
```

### Process Payment

```bash
curl -X POST http://localhost:3000/api/payments/process \
  -H "Content-Type: application/json" \
  -d '{
    "orderId": "your-order-id",
    "paymentMethod": "credit_card",
    "paymentDetails": {
      "cardLast4": "1234",
      "cardBrand": "Visa"
    }
  }'
```

## Project Structure

```
.
├── src/                    # Source code
│   ├── models/             # MongoDB models
│   │   ├── Cart.js         # Cart model
│   │   ├── Order.js        # Order model
│   │   └── Payment.js      # Payment model
│   ├── routes/             # API routes
│   │   ├── cartRoutes.js   # Cart endpoints
│   │   ├── orderRoutes.js  # Order endpoints
│   │   └── paymentRoutes.js # Payment endpoints
│   └── app.js              # Main application file
├── Dockerfile              # Docker image definition
├── docker-compose.yml      # Docker-compose configuration
├── .env                    # Environment variables
├── package.json            # Node.js dependencies
└── README.md               # This file
```

## Integration with Other Microservices

This Order Service is designed to work as part of a microservices ecosystem. It can integrate with:

- **Product Service**: For retrieving product details and inventory information
- **User Service**: For user authentication and profile information
- **Notification Service**: For sending order and payment confirmations

## Development

### Testing

Run the test suite:

```bash
npm test
```

### Linting

Check code quality:

```bash
npm run lint
```

## Deployment

### Production Environment

For production deployment:

1. Update the `.env` file with production values
2. Build optimized Docker images:
   ```bash
   docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
   ```

### Scaling

For horizontal scaling:

1. Deploy behind a load balancer
2. Scale MongoDB with replica sets
3. Use a container orchestration platform like Kubernetes

## Troubleshooting

### Common Issues

1. **MongoDB Connection Error**:
   - Ensure MongoDB is running
   - Check the MongoDB connection string in `.env`

2. **Container Not Starting**:
   - Check Docker logs: `docker-compose logs order-service`
   - Verify port availability: `lsof -i :3000`

## Contributing

1. Fork the repository
2. Create your feature branch: `git checkout -b feature/your-feature`
3. Commit your changes: `git commit -m 'Add some feature'`
4. Push to the branch: `git push origin feature/your-feature`
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
